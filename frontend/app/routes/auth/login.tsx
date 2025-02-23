import { Link, useNavigate } from 'react-router';
import { Button } from '~/components/ui/button';
import { Input } from '~/components/ui/input';
import { Label } from '~/components/ui/label';
import { ChevronLeft } from 'lucide-react';
import { client } from '~/clients/client';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import * as zod from 'zod';
import toast from 'react-hot-toast';
import type { Route } from './+types/login';
import { redirectIfLoggedIn } from '~/common/auth';
import { useEffect, useRef } from 'react';
import { GOOGLE_CLIENT_ID } from '~/constants/env.client';
import type { ErrorResponse } from '~/clients/types';

const LoginSchema = zod.object({
  email: zod.string().nonempty({ message: 'Please enter your email' }).email({ message: 'Please enter a valid email' }),
  password: zod.string().nonempty({ message: 'Please enter your password' }),
});
type LoginSchema = zod.infer<typeof LoginSchema>;

export async function loader({ request }: Route.LoaderArgs) {
  await redirectIfLoggedIn(request, '/section');
}

export default function Login() {
  const navigate = useNavigate();
  const googleSignInButtonRef = useRef<HTMLDivElement>(null);
  const {
    register,
    formState: { errors },
    handleSubmit,
  } = useForm({
    resolver: zodResolver(LoginSchema),
  });
  const { mutateAsync: login, isPending: isLoginPending } = client.useMutation('post', '/auth/v1/login');
  const { mutateAsync: loginWithGoogle, isPending: isLoginWithGooglePending } = client.useMutation(
    'post',
    '/auth/v1/login/google',
  );

  const handleLoginWithEmail = async ({ email, password }: LoginSchema) => {
    try {
      await login({ body: { email, password } });
      toast.dismiss();
      toast.success('Logged in successfully');
      navigate('/section');
    } catch (error) {
      const isPendingVerification = (error as ErrorResponse)?.error === 'account not verified';
      if (isPendingVerification) {
        toast('Verify your email', { icon: '⚠️' });
        navigate('/pending-verification');
        return;
      }
      const isInvalidCredentials = (error as ErrorResponse)?.error === 'invalid email or password';
      if (isInvalidCredentials) {
        toast.error('Invalid email or password');
        return;
      }
      toast.error('Something went wrong');
    }
  };

  const handleLoginWithGoogle = async ({ credential }: google.accounts.id.CredentialResponse) => {
    try {
      await loginWithGoogle({ body: { idToken: credential } });
      toast.dismiss();
      toast.success('Logged in successfully');
      navigate('/section');
    } catch {
      toast.error('Something went wrong');
    }
  };

  useEffect(() => {
    window.google.accounts.id.initialize({
      client_id: GOOGLE_CLIENT_ID,
      callback: handleLoginWithGoogle,
    });
    if (googleSignInButtonRef.current) {
      window.google.accounts.id.renderButton(googleSignInButtonRef.current, {
        type: 'standard',
        theme: 'outline',
        size: 'large',
      });
    }
  }, []);

  return (
    <div className="flex flex-col gap-12 p-8">
      <Link to="/" className="flex items-center gap-1.5 text-secondary">
        <ChevronLeft /> Back
      </Link>
      <div className="flex flex-col gap-14 w-7/8 md:w-5/6 lg:w-2/3 mx-auto">
        <div className="flex flex-col gap-6">
          <h1 className="text-4xl font-bold">Login</h1>
          <h2 className="text-lg text-secondary">Welcome back to Noah English!</h2>
        </div>
        <div className="flex flex-col gap-10">
          <form className="flex flex-col gap-5" onSubmit={handleSubmit(handleLoginWithEmail)}>
            <div className="flex flex-col gap-2.5">
              <Label htmlFor="email" className="text-secondary">
                Email
              </Label>
              <Input id="email" placeholder="Enter your email" className="h-14" {...register('email')} />
              <p className="text-destructive">{errors.email?.message}</p>
            </div>
            <div className="flex flex-col gap-2.5">
              <div className="flex justify-between items-center">
                <Label htmlFor="password" className="text-secondary">
                  Password
                </Label>
                <Link to="/forget-password" className="text-sm text-blue-500 hover:underline">
                  Forget password
                </Link>
              </div>
              <Input
                id="password"
                type="password"
                placeholder="Enter your password"
                className="h-14"
                {...register('password')}
              />
              <p className="text-destructive">{errors.password?.message}</p>
            </div>
            <Button className="h-14 text-md" disabled={isLoginPending || isLoginWithGooglePending}>
              Sign in
            </Button>
          </form>
          <div className="text-center text-secondary">
            Don&apos;t have an account?{' '}
            <Link to="/register" className="text-blue-500 hover:underline">
              Create an account
            </Link>
          </div>
          <div className="flex items-center gap-2">
            <div className="border-t border-secondary w-1/2"></div>
            <div className="text-gray-500">or</div>
            <div className="border-t border-secondary w-1/2"></div>
          </div>
          <div ref={googleSignInButtonRef} />
          {/* <Button className="h-14 text-md" variant="outline" disabled={isPending}>
            Sign in with Google <img src={google} />
          </Button> */}
        </div>
      </div>
    </div>
  );
}
