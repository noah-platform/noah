import { type RouteConfig, index, layout, route } from '@react-router/dev/routes';

export default [
  index('routes/index.tsx'),
  route('debug', 'routes/debug.tsx'),
  layout('routes/question-bank/layout.tsx', [route('section', 'routes/question-bank/section.tsx')]),
  layout('routes/auth/layout.tsx', [
    route('login', 'routes/auth/login.tsx'),
    route('register', 'routes/auth/register.tsx'),
    route('logout', 'routes/auth/logout.tsx'),
    route('forget-password', 'routes/auth/forget-password.tsx'),
    route('forget-password/:token', 'routes/auth/reset-password.tsx'),
    route('pending-verification', 'routes/auth/pending-verification.tsx'),
    route('verify-email/:token', 'routes/auth/verify-email.tsx'),
  ]),
] satisfies RouteConfig;
