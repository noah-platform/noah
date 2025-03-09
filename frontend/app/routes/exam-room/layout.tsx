import { ChevronLeft, ChevronRight, Volume2 } from 'lucide-react';
import type { ReactNode } from 'react';
import { Link } from 'react-router';
import { cn } from '~/lib/utils';
import { useExam } from './context';
import { TimeRemaining } from './components/time-remaining';
import { Slider } from '~/components/ui/slider';
import { Submitting } from './components/submitting';

interface LayoutProps {
  children: ReactNode;
}
export function Layout({ children }: LayoutProps) {
  const {
    exam,
    hasAudio,
    volume,
    setVolume,
    questionCount,
    startedAt,
    elapsedTime,
    currentSection,
    currentQuestion,
    questionSectionMap,
    reviews,
    toggleReview,
    nextQuestion,
    previousQuestion,
    jumpToQuestion,
    isCompleted,
    isSubmitting,
    handleSubmit,
  } = useExam();

  return (
    <div className="flex flex-col">
      <div className="flex flex-1 justify-center bg-primary min-h-[90px]">
        <div className={cn('grid grid-cols-3 w-5/6 items-center', isCompleted && 'grid-cols-2')}>
          <div className="flex justify-start items-center gap-4">
            <Link to="/section">
              <button className="text-md font-semibold text-black bg-white px-4 py-2 rounded-full">Exit</button>
            </Link>
          </div>
          {!isCompleted && <TimeRemaining startedAt={startedAt} elapsedTime={elapsedTime} duration={exam.duration} />}
          <div className="flex justify-end items-center gap-4">
            <button className="text-md font-semibold text-white bg-red-600 hover:bg-red-700 px-4 py-2 rounded-full">
              Help
            </button>
            <button className="text-md font-semibold text-black bg-white px-4 py-2 rounded-full" onClick={handleSubmit}>
              Submit
            </button>
            {hasAudio && (
              <div className="flex items-center gap-2 ml-4">
                <Volume2 className="text-white h-8 w-8" strokeWidth={2} />
                <Slider
                  className="w-[120px]"
                  value={[volume]}
                  onValueChange={([value]) => setVolume(value)}
                  max={100}
                  step={1}
                />
              </div>
            )}
          </div>
        </div>
      </div>
      <div className="flex flex-col w-10/11 lg:w-5/6 mx-auto my-4">
        {!isCompleted ? children : <Submitting exam={exam} isSubmitting={isSubmitting} />}
      </div>
      <div className="bg-primary fixed w-full bottom-0 left-0">
        <div className="flex justify-between items-center gap-2 w-5/6 min-h-[60px] mx-auto">
          <div className="flex items-center gap-2">
            <p className="text-white text-center text-sm mr-4">
              Double click to <br /> mark for review
            </p>
            {Array.from({ length: questionCount }).map((_, index) => (
              <button
                key={index}
                className={cn(
                  'w-8 h-8 bg-white font-bold text-black border-2 border-primary',
                  questionSectionMap[index] === currentSection && 'border-yellow-500',
                  index === currentQuestion && 'bg-yellow-500',
                  reviews.includes(index) && 'text-white bg-red-600',
                  reviews.includes(index) && index === currentQuestion && 'text-yellow-500'
                )}
                onClick={() => jumpToQuestion(index)}
                onDoubleClick={() => toggleReview(index)}
              >
                {index + 1}
              </button>
            ))}
          </div>
          <div className="flex items-center gap-3">
            <button
              className="flex justify-center items-center w-10 h-10 bg-white text-yellow-500 hover:border-2 hover:border-yellow-500 rounded-full hover:shadow-lg"
              onClick={previousQuestion}
            >
              <ChevronLeft size={32} strokeWidth={3} />
            </button>
            <button
              className="flex justify-center items-center w-10 h-10 bg-white text-yellow-500 hover:border-2 hover:border-yellow-500 rounded-full"
              onClick={nextQuestion}
            >
              <ChevronRight size={32} strokeWidth={3} />
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
