import { Link, useNavigate } from 'react-router';
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

export async function loader({ request }: Route.LoaderArgs) {
  await redirectIfLoggedIn(request, '/section');
}

const ForgetPasswordSchema = zod.object({
  email: zod.string().nonempty({ message: 'Please enter your email' }).email({ message: 'Please enter a valid email' }),
});
type ForgetPasswordSchema = zod.infer<typeof ForgetPasswordSchema>;

export default function ForgetPassword() {
  const navigate = useNavigate();
  const {
    register,
    formState: { errors },
    handleSubmit,
  } = useForm({
    resolver: zodResolver(ForgetPasswordSchema),
  });
  const { mutateAsync, isPending } = client.useMutation('post', '/account/v1/reset-password');

  const handleRegister = async ({ email }: ForgetPasswordSchema) => {
    try {
      await mutateAsync({ body: { email } });
      toast.dismiss();
      toast.success('Email sent successfully');
      navigate('/pending-verification');
    } catch {
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
          <h1 className="text-4xl font-bold">Forget password</h1>
          <h2 className="text-lg text-secondary">Don&apos;t worry! Just a few steps away.</h2>
        </div>
        <div className="flex flex-col gap-10">
          <form className="flex flex-col gap-5" onSubmit={handleSubmit(handleRegister)}>
            <div className="flex flex-col gap-2.5">
              <Label htmlFor="email" className="text-secondary">
                Email
              </Label>
              <Input id="email" placeholder="Enter your email" className="h-14" {...register('email')} />
              {errors.email && <p className="text-destructive">{errors.email.message}</p>}
            </div>
            <Button className="h-14 text-md" disabled={isPending}>
              Continue
            </Button>
          </form>
        </div>
      </div>
    </div>
  );
}
