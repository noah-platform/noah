import { redirect } from 'react-router';
import axios from 'axios';
import { AUTH_SESSION_SERVER_BASE_URL } from '~/constants/env.server';
import cookie from 'cookie';

export const COOKIE_NAME = 'noahses';
export const USER_ID_HEADER_NAME = 'x-noah-user-id';

export async function redirectIfLoggedIn(request: Request, path: string) {
  const cookies = request.headers.get('cookie') ?? '';
  const sessionId = cookie.parse(cookies)[COOKIE_NAME];
  if (sessionId) {
    throw redirect(path);
  }
}

export async function requireAuth(request: Request): Promise<{ userId: string; sessionId: string }> {
  const cookies = request.headers.get('cookie') ?? '';
  const sessionId = cookie.parse(cookies)[COOKIE_NAME];
  if (!sessionId) {
    throw redirect('/login');
  }

  const response = await axios.get(AUTH_SESSION_SERVER_BASE_URL + '/extauth/frontend', {
    headers: { Cookie: cookie.serialize(COOKIE_NAME, sessionId) },
  });
  if (response.status !== 200) {
    throw redirect('/login');
  }

  const userId = response.headers[USER_ID_HEADER_NAME];
  if (!userId) {
    throw redirect('/login');
  }

  return { userId, sessionId };
}
