import { requireAuth } from '~/common/auth';
import type { Route } from './+types/home';
import { Link } from 'react-router';
import { Button } from '~/components/ui/button';

export async function loader({ request }: Route.LoaderArgs) {
  await requireAuth(request);
}

export default function Home() {
  return (
    <div className="flex justify-between p-8">
      <h1 className="text-xl">Home</h1>
      <Button asChild variant="outline" className="w-fit">
        <Link to="/logout">Logout</Link>
      </Button>
    </div>
  );
}
