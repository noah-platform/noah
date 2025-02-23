import createFetchClient from 'openapi-fetch';
import createClient from 'openapi-react-query';
import type { paths } from './openapi';
import { API_BASE_URL } from '~/constants/env.client';

// TODO: Fix hard-coded baseUrl
// API_BASE_URL cuurently does not resolved server-side
export const fetchClient = createFetchClient<paths>({
  baseUrl: API_BASE_URL || 'http://api.noah-platform.thanayut.in.th',
  credentials: 'include',
});
export const client = createClient(fetchClient);
