import React, { createContext, useCallback, useContext, useMemo, useState, type ReactNode } from 'react';
import { LocalStorageKey, type Exam } from '../../common/types';
import { FormProvider, useForm } from 'react-hook-form';

interface ExamContext {
  testId: string;
  exam: Exam;

  startedAt: Date;

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
  exam: Exam;
  children: ReactNode;
}
export default function ExamContextProvider({ testId, exam, children }: ExamContextProviderProps) {
  const startedAt = useMemo(() => new Date(), []);

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
    () => exam.sections.flatMap((sections, index) => Array({ length: sections.questionCount }).map(() => index)),
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

  const form = useForm();

  return (
    <ExamRoomContext.Provider
      value={{
        testId,
        exam,
        startedAt,

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
      <FormProvider {...form}>{children}</FormProvider>
    </ExamRoomContext.Provider>
  );
}
