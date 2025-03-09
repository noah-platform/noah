import { type RouteConfig, index, layout, prefix, route } from '@react-router/dev/routes';

export default [
  route('debug', 'routes/debug.tsx'),
  layout('routes/layout.tsx', [index('routes/index.tsx')]),
  layout('routes/question-bank/layout.tsx', [
    route('section', 'routes/question-bank/section.tsx'),
    route('review', 'routes/question-bank/review.tsx'),
  ]),
  ...prefix('test', [
    route(':testId', 'routes/exam-cover/cover.tsx'),
    route(':testId/start', 'routes/exam-room/exam.tsx'),
  ]),
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
