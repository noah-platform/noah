import { useNavigate } from 'react-router';
import { Button } from '~/components/ui/button';
import { Input } from '~/components/ui/input';
import { Label } from '~/components/ui/label';
import { ChevronLeft } from 'lucide-react';
import { redirectIfLoggedIn } from '~/common/auth';
import type { Route } from './+types/register';
import * as zod from 'zod';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { client } from '~/clients/client';
import toast from 'react-hot-toast';
import type { ErrorResponse } from '~/clients/types';

export async function loader({ request }: Route.LoaderArgs) {
  await redirectIfLoggedIn(request, '/section');
}

const RegisterSchema = zod
  .object({
    fullname: zod.string().nonempty({ message: 'Please enter your name' }),
    email: zod
      .string()
      .nonempty({ message: 'Please enter your email' })
      .email({ message: 'Please enter a valid email' }),
    password: zod
      .string()
      .nonempty({ message: 'Please enter your password' })
      .min(8, { message: 'Password is too short' })
      .max(64, { message: 'Password is too long' }),
    confirmPassword: zod.string().nonempty({ message: 'Please enter your password again' }),
  })
  .superRefine((data, ctx) => {
    if (data.password !== data.confirmPassword) {
      ctx.addIssue({ code: 'custom', message: 'Passwords do not match', path: ['confirmPassword'] });
    }
  });
type RegisterSchema = zod.infer<typeof RegisterSchema>;

export default function Register() {
  const navigate = useNavigate();
  const {
    register,
    formState: { errors },
    handleSubmit,
    setError,
  } = useForm({
    resolver: zodResolver(RegisterSchema),
  });
  const { mutateAsync, isPending } = client.useMutation('post', '/account/v1/register');

  const handleRegister = async ({ fullname, email, password }: RegisterSchema) => {
    try {
      await mutateAsync({ body: { name: fullname, email, password } });
      toast.dismiss();
      toast.success('Welcome to NOAH ENGLISH!');
      navigate('/pending-verification');
    } catch (error) {
      const isEmailAlreadyExists = (error as ErrorResponse)?.error === 'account already exists';
      if (isEmailAlreadyExists) {
        setError('email', { message: 'This email is already registered' });
      } else {
        toast.error('Something went wrong');
      }
    }
  };

  return (
    <div className="flex flex-col gap-12 p-8">
      <button onClick={() => navigate(-1)} className="flex items-center gap-1.5 text-secondary">
        <ChevronLeft /> Back
      </button>
      <div className="flex flex-col gap-14 w-7/8 md:w-5/6 lg:w-2/3 mx-auto">
        <div className="flex flex-col gap-6">
          <h1 className="text-4xl font-bold">Sign up</h1>
          <h2 className="text-lg text-secondary">Welcome to NOAH ENGLISH!</h2>
        </div>
        <div className="flex flex-col gap-10">
          <form className="flex flex-col gap-5" onSubmit={handleSubmit(handleRegister)}>
            <div className="flex flex-col gap-2.5">
              <Label htmlFor="fullname" className="text-secondary">
                Full Name
              </Label>
              <Input id="fullname" placeholder="Enter your name" className="h-14" {...register('fullname')} />
              {errors.fullname && <p className="text-destructive">{errors.fullname.message}</p>}
            </div>
            <div className="flex flex-col gap-2.5">
              <Label htmlFor="email" className="text-secondary">
                Email
              </Label>
              <Input id="email" placeholder="Enter your email" className="h-14" {...register('email')} />
              {errors.email && <p className="text-destructive">{errors.email.message}</p>}
            </div>
            <div className="flex flex-col gap-2.5">
              <Label htmlFor="password" className="text-secondary">
                Password
              </Label>
              <Input
                id="password"
                type="password"
                placeholder="Enter your password"
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
                placeholder="Enter your password again"
                className="h-14"
                {...register('confirmPassword')}
              />
              {errors.confirmPassword && <p className="text-destructive">{errors.confirmPassword.message}</p>}
            </div>
            <Button className="h-14 text-md" disabled={isPending}>
              Create Account
            </Button>
          </form>
        </div>
      </div>
    </div>
  );
}
