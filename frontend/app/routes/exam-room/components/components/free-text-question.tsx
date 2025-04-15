import { useFormContext, Controller } from 'react-hook-form';
import type { Question } from '~/common/types';
import { RichTextPreview } from '~/components/richtext-preview';
import { useExam } from '../../context';
import { cn } from '~/lib/utils';

interface FreeTextQuestionProps {
  questions: Question[];
}
export function FreeTextQuestion({ questions }: FreeTextQuestionProps) {
  return (
    <div>
      {questions.map((question) => (
        <div className="grid grid-cols-2 gap-4" key={question.questionId}>
          <div className="flex flex-col gap-2 relative top-[-8px]">
            <RichTextPreview value={question.body} />
            {(question.imageUrls ?? []).map((url, index) => (
              <img key={index} src={url} />
            ))}
          </div>
          <TextArea questionId={question.questionId} />
        </div>
      ))}
    </div>
  );
}

interface TextAreaProps {
  questionId: string;
}
function TextArea({ questionId }: TextAreaProps) {
  const { isReviewing } = useExam();
  const { control } = useFormContext();

  return (
    <Controller
      control={control}
      name={questionId}
      render={({ field: { value, ...rest } }) => (
        <div className="flex flex-col gap-2">
          <textarea
            disabled={isReviewing}
            className={cn(
              'min-h-[600px] p-2 bg-white border rounded-sm disabled:bg-gray-100',
              isReviewing && 'hover:cursor-not-allowed'
            )}
            value={value}
            {...rest}
          />
          <p>{!value ? '0 word' : `${(value ?? '').trim().split(/\s+/).length} words`}</p>
        </div>
      )}
    />
  );
}
