import { requireAuth } from '~/common/auth';
import type { Route } from './+types/home';

export async function loader({ request }: Route.LoaderArgs) {
  await requireAuth(request);
}

export default function Home() {
  return <div>Home</div>;
}
