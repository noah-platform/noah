import { useState, useEffect } from 'react';
import { calculateRemainingTime } from '../utils';
import type { Exam } from '../../../common/types';

interface TimeRemainingProps {
  startedAt: Date;
  exam: Exam;
}
export function TimeRemaining({ startedAt, exam }: TimeRemainingProps) {
  const [timeRemaining, setTimeRemaining] = useState<{ minutes: number; seconds: number; duration: number }>(() =>
    calculateRemainingTime(startedAt, exam.duration - 1)
  );

  useEffect(() => {
    const interval = setInterval(() => {
      setTimeRemaining(calculateRemainingTime(startedAt, exam.duration));
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
