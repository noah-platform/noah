import { useRef, useState } from 'react';
import type { Question } from '~/common/types';
import { Button } from '~/components/ui/button';
import { useExam } from '../../context';

interface ShortVoiceQuestionProps {
  questions: Question[];
}
export function ShortVoiceQuestion({ questions }: ShortVoiceQuestionProps) {
  return (
    <div className="flex flex-col gap-8 items-center my-3">
      {questions.map((question, index) => (
        <VoiceQuestion key={question.questionId} index={index} question={question} />
      ))}
    </div>
  );
}

interface VoiceQuestionProps {
  index: number;
  question: Question;
}
function VoiceQuestion({ index, question }: VoiceQuestionProps) {
  const { isReviewing } = useExam();
  const [questionRepeatCount, setQuestionRepeatCount] = useState(0);
  const [isAudioPlaying, setIsAudioPlaying] = useState(false);
  const [isRecording, setIsRecording] = useState(false);
  const [timeRemaining, setTimeRemaining] = useState(question.maximumAnswerDuration);
  const intervalRef = useRef<NodeJS.Timeout | null>(null);

  const handlePlayAudio = () => {
    const audio = new Audio(question.audioUrl);
    audio.addEventListener('play', () => {
      setIsAudioPlaying(true);
    });
    audio.addEventListener('ended', () => {
      setIsAudioPlaying(false);
      setQuestionRepeatCount((prev) => prev + 1);
    });
    audio.play();
  };

  const handleStartRecording = async () => {
    setIsRecording(true);
    intervalRef.current = setInterval(() => {
      if (timeRemaining <= 0) {
        if (intervalRef.current) clearInterval(intervalRef.current);
        setIsRecording(false);
        return;
      }
      setTimeRemaining((prev) => prev - 1);
    }, 1000);
  };

  const handleStopRecording = () => {
    setIsRecording(false);
    if (intervalRef.current) clearInterval(intervalRef.current);
    setTimeRemaining(0);
  };

  if (isReviewing) {
    return (
      <div className="grid grid-cols-[1fr_2fr_1fr] items-center w-full" key={question.questionId}>
        <h3 className="font-medium">Question {index + 1}:</h3>
        <div className="flex items-center gap-2">
          <Button variant="outline" className="w-full" disabled>
            Recording completed!
          </Button>
        </div>
        <p className="justify-self-end text-sm">Maximum {question.maximumAnswerDuration} seconds</p>
      </div>
    );
  }
  return (
    <div className="grid grid-cols-[1fr_2fr_1fr] items-center w-full" key={question.questionId}>
      <h3 className="font-medium">Question {index + 1}:</h3>
      <div className="flex items-center gap-2">
        {questionRepeatCount < question.maximumQuestionRepeat && !isRecording && timeRemaining > 0 && (
          <Button variant="outline" onClick={handlePlayAudio} className="w-full" disabled={isAudioPlaying}>
            {!isAudioPlaying
              ? questionRepeatCount === 0
                ? 'Listen and Start Recording'
                : 'Listen Again'
              : 'Playing...'}
          </Button>
        )}
        {questionRepeatCount > 0 && !isRecording && timeRemaining > 0 && (
          <Button onClick={handleStartRecording} className="w-full">
            Start Recording
          </Button>
        )}
        {isRecording && (
          <Button variant="outline" className="w-full" onClick={handleStopRecording}>
            Recording... {timeRemaining}s
          </Button>
        )}
        {timeRemaining <= 0 && (
          <Button variant="outline" className="w-full" disabled>
            Recording completed!
          </Button>
        )}
      </div>
      <p className="justify-self-end text-sm">Maximum {question.maximumAnswerDuration} seconds</p>
    </div>
  );
}
