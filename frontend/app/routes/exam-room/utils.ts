import dayjs from 'dayjs';

export function calculateRemainingTime(startedAt: Date, duration: number) {
  const now = dayjs();
  const endsAt = dayjs(startedAt).add(duration, 'seconds');
  const minutesRemaining = endsAt.diff(now, 'minutes');
  const secondsRemaining = endsAt.diff(now, 'seconds') % 60;
  const remainingDuration = endsAt.diff(now, 'seconds');
  return { minutes: minutesRemaining, seconds: secondsRemaining, duration: remainingDuration };
}
