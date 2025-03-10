import { useState, useEffect } from 'react';
import { calculateRemainingTime } from '../utils';

interface TimeRemainingProps {
  startedAt: Date;
  elapsedTime: number;
  duration: number;
}
export function TimeRemaining({ startedAt, elapsedTime, duration }: TimeRemainingProps) {
  const [timeRemaining, setTimeRemaining] = useState<{ minutes: number; seconds: number; duration: number }>(() =>
    calculateRemainingTime(startedAt, duration - elapsedTime)
  );

  useEffect(() => {
    const interval = setInterval(() => {
      setTimeRemaining(calculateRemainingTime(startedAt, duration - elapsedTime));
    }, 1000);
    return () => clearInterval(interval);
  }, [startedAt]);

  return (
    <div className="flex flex-col items-center text-white">
      <p className="text-3xl text-red-600">
        {timeRemaining.minutes}:{String(timeRemaining.seconds).padStart(2, '0')}
      </p>
      <p>remaining</p>
    </div>
  );
}
