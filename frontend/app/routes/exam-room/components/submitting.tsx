import { Link } from 'react-router';
import * as Spinner from 'react-spinners';
import type { Exam } from '~/common/types';
import { Button } from '~/components/ui/button';

interface SubmittingProps {
  exam: Exam;
  isSubmitting: boolean;
}
export function Submitting({ exam, isSubmitting }: SubmittingProps) {
  return (
    <div className="flex flex-col my-8">
      <h1 className="relative z-50 text-2xl font-bold mx-8">{exam.title}</h1>
      <div className="relative top-[-14px] bg-gray-100 p-8 rounded-2xl">
        <div className="flex flex-col gap-12 items-center text-center my-20">
          <p className="text-3xl font-semibold">Test completed!</p>
          {isSubmitting ? (
            <div className="flex flex-col items-center gap-4">
              <p>Grading your test...</p>
              <Spinner.BeatLoader color="#070559" />
            </div>
          ) : (
            <Link to="/review">
              <Button className="min-w-[200px] h-12 text-lg font-medium">View your score</Button>
            </Link>
          )}
        </div>
      </div>
    </div>
  );
}
