import { ResponseType, type QuestionSet } from '~/common/types';
import { FreeTextQuestion } from './components/free-text-question';
import { NoteCompletionQuestion } from './components/note-completion-question';
import { match } from 'ts-pattern';
import { MultipleChoiceQuestion } from './components/multiple-choice-question';
import { MultipleSelectQuestion } from './components/multiple-select-question';

interface QuesitonSetProps {
  questionSet: QuestionSet;
  showInstruction?: boolean;
}
export function QuestionSet({ questionSet, showInstruction }: QuesitonSetProps) {
  return (
    <div className="flex flex-col gap-3 bg-gray-100 p-6 rounded-xl">
      {showInstruction && questionSet.instruction && <p className="font-bold">{questionSet.instruction}</p>}
      {match(questionSet.responseType)
        .with(ResponseType.FREE_TEXT, () => <FreeTextQuestion questions={questionSet.questions} />)
        .with(ResponseType.NOTE_COMPLETION, () => <NoteCompletionQuestion questions={questionSet.questions} />)
        .with(ResponseType.MULTIPLE_CHOICE, () => <MultipleChoiceQuestion questions={questionSet.questions} />)
        .with(ResponseType.MULTIPLE_SELECT, () => <MultipleSelectQuestion questions={questionSet.questions} />)
        .otherwise(() => (
          <p>Unsupported Response Type</p>
        ))}
    </div>
  );
}
