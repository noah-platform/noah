import { client } from '~/clients/client';
import { requireAuth } from '~/common/auth';
import type { Route } from './+types/debug';
import { Button } from '~/components/ui/button';

export async function loader({ request }: Route.LoaderArgs) {
  const { userId } = await requireAuth(request);

  const shouldThrow = new URL(request.url).searchParams.get('error') === 'true';
  if (shouldThrow) {
    throw new Error('Debug error from loader');
  }

  return userId;
}

export default function Debug({ loaderData }: Route.ComponentProps) {
  const userId = loaderData;
  const { isLoading, isError, data, error } = client.useQuery('get', '/auth/v1/me');

  const handleClick = () => {
    throw new Error('Debug error from button click');
  };

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
      <Button variant="outline" onClick={handleClick}>
        Throw error
      </Button>
    </div>
  );
}
