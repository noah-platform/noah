import { client } from '~/clients/client';
import { requireAuth } from '~/common/auth';
import type { Route } from './+types/debug';

export async function loader({ request }: Route.LoaderArgs) {
  const { userId } = await requireAuth(request);
  return userId;
}

export default function Debug({ loaderData }: Route.ComponentProps) {
  const userId = loaderData;

  const { isLoading, isError, data, error } = client.useQuery('get', '/auth/v1/me');
  if (isLoading) {
    return <p>Loading...</p>;
  }
  if (isError) {
    return <p>Error: {error instanceof Error ? error.message : 'Something went wrong'}</p>;
  }
  return (
    <div className="p-6">
      <p>UserID: {userId}</p>
      <pre>{JSON.stringify(data?.data, null, 2)}</pre>
    </div>
  );
}
