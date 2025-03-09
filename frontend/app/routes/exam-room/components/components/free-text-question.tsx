import { useFormContext, Controller } from 'react-hook-form';
import type { Question } from '~/common/types';
import { RichTextPreview } from '~/components/richtext-preview';

interface FreeTextQuestionProps {
  questions: Question[];
}
export function FreeTextQuestion({ questions }: FreeTextQuestionProps) {
  return (
    <div>
      {questions.map((question) => (
        <div className="grid grid-cols-2 gap-4" key={JSON.stringify(question)}>
          <div className="flex flex-col gap-2 relative top-[-8px]">
            <RichTextPreview value={question.body} />
            {(question.imageUrls ?? []).map((url, index) => (
              <img key={index} src={url} />
            ))}
          </div>
          {/* TODO: Add questionId */}
          <TextArea questionId={question.body.slice(0, 15)} />
        </div>
      ))}
    </div>
  );
}

interface TextAreaProps {
  questionId: string;
}
function TextArea({ questionId }: TextAreaProps) {
  const { control } = useFormContext();

  return (
    <Controller
      control={control}
      name={questionId}
      render={({ field: { value, ...rest } }) => (
        <textarea className="min-h-[600px] p-2 bg-white border rounded-sm" value={value} {...rest} />
      )}
    />
  );
}
