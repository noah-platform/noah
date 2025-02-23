import { Link, useNavigate } from 'react-router';
import { Button } from '~/components/ui/button';
import { Input } from '~/components/ui/input';
import { Label } from '~/components/ui/label';
import { ChevronLeft } from 'lucide-react';
import { redirectIfLoggedIn } from '~/common/auth';
import type { Route } from './+types/reset-password';
import * as zod from 'zod';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { client } from '~/clients/client';
import toast from 'react-hot-toast';
import type { ErrorResponse } from '~/clients/types';

export async function loader({ request }: Route.LoaderArgs) {
  await redirectIfLoggedIn(request, '/section');
}

const ResetPasswordSchema = zod
  .object({
    password: zod
      .string()
      .nonempty({ message: 'Please enter a password' })
      .min(8, { message: 'Password is too short' })
      .max(64, { message: 'Password is too long' }),
    confirmPassword: zod.string().nonempty({ message: 'Please enter the password again' }),
  })
  .superRefine((data, ctx) => {
    if (data.password !== data.confirmPassword) {
      ctx.addIssue({ code: 'custom', message: 'Passwords do not match', path: ['confirmPassword'] });
    }
  });
type ResetPasswordSchema = zod.infer<typeof ResetPasswordSchema>;

export default function ResetPassword({ params }: Route.ComponentProps) {
  const { token } = params;
  const navigate = useNavigate();
  const {
    register,
    formState: { errors },
    handleSubmit,
  } = useForm({
    resolver: zodResolver(ResetPasswordSchema),
  });
  const { mutateAsync, isPending } = client.useMutation('post', '/account/v1/reset-password/{token}');

  const handleResetPassword = async ({ password }: ResetPasswordSchema) => {
    try {
      await mutateAsync({ params: { path: { token } }, body: { password } });
      toast.dismiss();
      toast.success('Password reset successfully');
      navigate('/login');
    } catch (error) {
      const isTokenInvalid = (error as ErrorResponse)?.error === 'invalid password reset token';
      if (isTokenInvalid) {
        toast.error('Invalid password reset token');
        navigate('/forget-password');
        return;
      }
      toast.error('Something went wrong');
    }
  };

  return (
    <div className="flex flex-col gap-12 p-8">
      <Link to="/login" className="flex items-center gap-1.5 text-secondary">
        <ChevronLeft /> Back
      </Link>
      <div className="flex flex-col gap-14 w-7/8 md:w-5/6 lg:w-2/3 mx-auto">
        <div className="flex flex-col gap-6">
          <h1 className="text-4xl font-bold">Reset password</h1>
          <h2 className="text-lg text-secondary">Please enter a new password.</h2>
        </div>
        <div className="flex flex-col gap-10">
          <form className="flex flex-col gap-5" onSubmit={handleSubmit(handleResetPassword)}>
            <div className="flex flex-col gap-2.5">
              <Label htmlFor="password" className="text-secondary">
                Password
              </Label>
              <Input
                id="password"
                type="password"
                placeholder="Enter a password"
                className="h-14"
                {...register('password')}
              />
              {errors.password && <p className="text-destructive">{errors.password.message}</p>}
            </div>
            <div className="flex flex-col gap-2.5">
              <Label htmlFor="confirm-password" className="text-secondary">
                Confirm Password
              </Label>
              <Input
                id="confirm-password"
                type="password"
                placeholder="Enter the password again"
                className="h-14"
                {...register('confirmPassword')}
              />
              {errors.confirmPassword && <p className="text-destructive">{errors.confirmPassword.message}</p>}
            </div>
            <Button className="h-14 text-md" disabled={isPending}>
              Reset password
            </Button>
          </form>
        </div>
      </div>
    </div>
  );
}
