import * as Sentry from '@sentry/node';
import { nodeProfilingIntegration } from '@sentry/profiling-node';

Sentry.init({
  enabled: false,

  dsn: 'https://cd257a261d3aa8b073ce3c2f7fb35462@o1418462.ingest.us.sentry.io/4508830487543808',
  integrations: [nodeProfilingIntegration()],
  tracesSampleRate: 1.0,
  profilesSampleRate: 1.0,
});
