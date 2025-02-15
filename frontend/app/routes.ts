import { type RouteConfig, index, layout, route } from '@react-router/dev/routes';

export default [
  index('routes/index.tsx'),
  route('home', 'routes/home.tsx'),
  route('debug', 'routes/debug.tsx'),
  layout('routes/auth/layout.tsx', [
    route('login', 'routes/auth/login.tsx'),
    route('register', 'routes/auth/register.tsx'),
    route('pending-verification', 'routes/auth/pending-verification.tsx'),
  ]),
] satisfies RouteConfig;
