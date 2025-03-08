import { createContext, useCallback, useContext, useMemo, useState, type ReactNode } from 'react';
import { LocalStorageKey, type Exam, type Session } from '../../common/types';
import { FormProvider, useForm } from 'react-hook-form';
import { AutoSave } from './components/auto-save';

interface ExamContext {
  testId: string;
  exam: Exam;

  startedAt: Date;
  elapsedTime: number;

  hasAudio: boolean;
  volume: number;
  setVolume: (volume: number) => void;

  questionCount: number;
  questionSectionMap: number[];

  currentQuestion: number;
  currentSection: number;

  reviews: number[];
  toggleReview: (index: number) => void;

  nextQuestion: () => void;
  previousQuestion: () => void;
  jumpToQuestion: (index: number) => void;
}
export const ExamRoomContext = createContext<ExamContext>({} as ExamContext);

export const useExam = () => {
  const context = useContext(ExamRoomContext);
  if (context === undefined) {
    throw new Error('useExam must be used within a ExamRoomProvider');
  }
  return context;
};

interface ExamContextProviderProps {
  testId: string;
  session: Session;
  children: ReactNode;
}
export default function ExamContextProvider({ testId, session, children }: ExamContextProviderProps) {
  const startedAt = useMemo(() => new Date(), []);
  const exam = session.test;
  const hasAudio = useMemo(() => exam.sections.some((section) => !!section.audioUrl), [exam]);
  const [volume, _setVolume] = useState(() =>
    Number(typeof window !== 'undefined' ? localStorage.getItem(LocalStorageKey.VOLUME) ?? 100 : 100)
  );
  const setVolume = useCallback((volume: number) => {
    localStorage.setItem(LocalStorageKey.VOLUME, volume.toString());
    _setVolume(volume);
  }, []);

  const questionCount = useMemo(() => exam.sections.reduce((acc, section) => acc + section.questionCount, 0), [exam]);
  const questionSectionMap = useMemo(
    () => exam.sections.flatMap((section, index) => Array.from({ length: section.questionCount }).map(() => index)),
    [exam]
  );

  const [reviews, setReviews] = useState<number[]>([]);
  const toggleReview = useCallback((index: number) => {
    setReviews((prev) => {
      if (prev.includes(index)) {
        return prev.filter((i) => i !== index);
      }
      return [...prev, index];
    });
  }, []);

  const [currentQuestion, setCurrentQuestion] = useState(0);
  const currentSection = useMemo(() => questionSectionMap[currentQuestion], [currentQuestion, questionSectionMap]);

  const nextQuestion = useCallback(() => {
    setCurrentQuestion((prev) => Math.min(prev + 1, questionCount - 1));
  }, [questionCount]);
  const previousQuestion = useCallback(() => {
    setCurrentQuestion((prev) => Math.max(prev - 1, 0));
  }, [questionCount]);
  const jumpToQuestion = useCallback((index: number) => {
    setCurrentQuestion(index);
  }, []);

  const form = useForm({
    defaultValues: session.answers,
  });

  console.log(session);

  return (
    <ExamRoomContext.Provider
      value={{
        testId,
        exam,
        startedAt,
        elapsedTime: session.elapsedTime,

        hasAudio,
        volume,
        setVolume,

        questionCount,
        questionSectionMap,

        currentQuestion,
        currentSection,

        reviews,
        toggleReview,

        nextQuestion,
        previousQuestion,
        jumpToQuestion,
      }}
    >
      <FormProvider {...form}>
        {children}
        <AutoSave testId={testId} />
      </FormProvider>
    </ExamRoomContext.Provider>
  );
}
