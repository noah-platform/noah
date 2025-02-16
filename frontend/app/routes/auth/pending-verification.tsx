import { Link } from 'react-router';
import { ChevronLeft } from 'lucide-react';
import { redirectIfLoggedIn } from '~/common/auth';
import type { Route } from './+types/register';

export async function loader({ request }: Route.LoaderArgs) {
  await redirectIfLoggedIn(request, '/home');
}

export default function PendingVerification() {
  return (
    <div className="flex flex-col gap-12 p-8">
      <Link to="/login" className="flex items-center gap-1.5 text-secondary">
        <ChevronLeft /> Back
      </Link>
      <div className="flex flex-col gap-14 w-7/8 md:w-5/6 lg:w-2/3 mx-auto">
        <div className="flex flex-col gap-6">
          <h1 className="text-4xl font-bold">Verify your email</h1>
          <h2 className="text-lg text-secondary">Just one last step!</h2>
        </div>
        <div className="flex flex-col gap-10">
          <p>We have sent an email to your inbox.</p>
        </div>
      </div>
    </div>
  );
}
