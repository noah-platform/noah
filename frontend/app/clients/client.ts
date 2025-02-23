import createFetchClient from 'openapi-fetch';
import createClient from 'openapi-react-query';
import type { paths } from './openapi';
import { API_BASE_URL } from '~/constants/env.client';

const fetchClient = createFetchClient<paths>({
  baseUrl: API_BASE_URL,
  credentials: 'include',
});
export const client = createClient(fetchClient);
