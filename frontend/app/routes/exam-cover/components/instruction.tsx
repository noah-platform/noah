import dayjs from 'dayjs';
import { Link } from 'react-router';
import type { ExamCover } from '~/common/types';
import { RichTextPreview } from '~/components/richtext-preview';
import { Button } from '~/components/ui/button';

interface ExamInstructionProps {
  cover: ExamCover;
}
export function ExamInstruction({ cover }: ExamInstructionProps) {
  const duration = dayjs.duration(cover.duration, 'seconds');
  return (
    <div className="flex flex-col my-8">
      <h1 className="relative z-50 text-2xl font-bold mx-8">{cover.title}</h1>
      <div className="relative top-[-14px] bg-gray-100 p-8 rounded-2xl">
        <p className="my-3">
          Time: {duration.hours() > 0 && `${duration.hours()} hour`}{' '}
          {duration.minutes() > 0 && `${duration.minutes()} minutes`}
        </p>
        <RichTextPreview value={cover.instruction} />
        <div className="flex justify-center mt-5 mb-1">
          <Link to="start">
            <Button className="min-w-[200px] h-12 text-lg font-medium">Start</Button>
          </Link>
        </div>
      </div>
    </div>
  );
}
