import { Link, useNavigate } from 'react-router';
import { ChevronLeft } from 'lucide-react';
import { redirectIfLoggedIn } from '~/common/auth';

import toast from 'react-hot-toast';
import { useEffect } from 'react';
import { client } from '~/clients/client';
import type { Route } from './+types/verify-email';

export async function loader({ request }: Route.LoaderArgs) {
  await redirectIfLoggedIn(request, '/home');
}

export default function VerifyEmail({ params }: Route.ComponentProps) {
  const { token } = params;
  const navigate = useNavigate();
  const { mutateAsync } = client.useMutation('post', '/account/v1/verify-email');

  useEffect(() => {
    const verifyEmail = async () => {
      try {
        await mutateAsync({ body: { token } });
        toast.success('Email verified!');
        navigate('/login');
      } catch {
        toast.error('Something went wrong');
        return;
      }
    };
    verifyEmail();
  }, []);

  return (
    <div className="flex flex-col gap-12 p-8">
      <Link to="/login" className="flex items-center gap-1.5 text-secondary">
        <ChevronLeft /> Back
      </Link>
      <div className="flex flex-col gap-14 w-7/8 md:w-5/6 lg:w-2/3 mx-auto">
        <div className="flex flex-col gap-6">
          <h1 className="text-4xl font-bold">Verifying your email</h1>
          <h2 className="text-lg text-secondary">Please wait a moment...</h2>
        </div>
      </div>
    </div>
  );
}
