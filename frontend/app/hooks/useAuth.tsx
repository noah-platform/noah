import { useMatches } from 'react-router';

export function useAuth() {
  const matches = useMatches();
  const root = matches.find((match) => match.id === 'root');
  if (!root) {
    return { isLoggedIn: false };
  }
  return { isLoggedIn: (root.data as { isLoggedIn?: boolean }).isLoggedIn ?? false };
}
