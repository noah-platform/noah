import { requireAuth } from '~/common/auth';
import type { Route } from './+types/review';

export async function loader({ request }: Route.LoaderArgs) {
  await requireAuth(request);
}

export default function Review() {
  return <h1>Review</h1>;
}
