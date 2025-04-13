import { useRef, useState } from 'react';
import type { Question } from '~/common/types';
import { RichTextPreview } from '~/components/richtext-preview';
import { Button } from '~/components/ui/button';
import { useExam } from '../../context';

interface LongVoiceQuestionProps {
  questions: Question[];
}
export function LongVoiceQuestion({ questions }: LongVoiceQuestionProps) {
  return (
    <div className="flex flex-col gap-8 items-center">
      {questions.map((question) => (
        <VoiceQuestion key={question.questionId} question={question} />
      ))}
    </div>
  );
}

interface VoiceQuestionProps {
  question: Question;
}
function VoiceQuestion({ question }: VoiceQuestionProps) {
  const { isReviewing } = useExam();
  const [isTaskCardRevealed, setIsTaskCardRevealed] = useState(false);
  const [isRecording, setIsRecording] = useState(false);
  const [preparationTimeLeft, setPreparationTimeLeft] = useState(question.preparationDuration);
  const [recordingTimeLeft, setRecordingTimeLeft] = useState(question.maximumAnswerDuration);
  const intervalRef = useRef<NodeJS.Timeout | null>(null);

  const handleRevealTaskCard = () => {
    setIsTaskCardRevealed(true);
    const intervalId = setInterval(() => {
      if (preparationTimeLeft <= 0) {
        clearInterval(intervalId);
        return;
      }
      setPreparationTimeLeft((prev) => prev - 1);
    }, 1000);
  };

  const handleStartRecording = () => {
    setIsRecording(true);
    intervalRef.current = setInterval(() => {
      if (recordingTimeLeft <= 0) {
        if (intervalRef.current) clearInterval(intervalRef.current);
        setIsRecording(false);
        return;
      }
      setRecordingTimeLeft((prev) => prev - 1);
    }, 1000);
  };

  const handleStopRecording = () => {
    if (intervalRef.current) clearInterval(intervalRef.current);
    setRecordingTimeLeft(0);
  };

  if (isReviewing) {
    return (
      <div className="flex flex-col items-center gap-6">
        <div className="bg-white rounded-lg px-9 py-3">
          <RichTextPreview value={question.taskCard} />
        </div>
        <p>Recording completed!</p>
      </div>
    );
  }
  if (!isTaskCardRevealed) {
    return (
      <div className="flex justify-center my-4">
        <Button size="lg" onClick={handleRevealTaskCard}>
          Reveal the card
        </Button>
      </div>
    );
  }
  if (!isRecording) {
    return (
      <div className="flex flex-col items-center gap-6">
        <div className="bg-white rounded-lg px-9 py-3">
          <RichTextPreview value={question.taskCard} />
        </div>
        <p>{preparationTimeLeft} seconds left to prepare</p>
        <div className="flex flex-col gap-2">
          <Button size="lg" onClick={handleStartRecording}>
            Start Now
          </Button>
          <p className="text-sm text-gray-500">
            You will have {question.maximumAnswerDuration} seconds to record your answer.
          </p>
        </div>
      </div>
    );
  }
  if (recordingTimeLeft <= 0) {
    return (
      <div className="flex flex-col items-center gap-6">
        <div className="bg-white rounded-lg px-9 py-3">
          <RichTextPreview value={question.taskCard} />
        </div>
        <p>Recording completed!</p>
      </div>
    );
  }
  return (
    <div className="flex flex-col items-center gap-6">
      <div className="bg-white rounded-lg px-9 py-3">
        <RichTextPreview value={question.taskCard} />
      </div>
      <div className="flex flex-col gap-2">
        <Button size="lg" onClick={handleStopRecording}>
          Recording... {recordingTimeLeft}s
        </Button>
      </div>
    </div>
  );
}
