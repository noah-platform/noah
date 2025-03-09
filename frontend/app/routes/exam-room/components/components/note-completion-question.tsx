import { useFormContext } from 'react-hook-form';
import reactStringReplace from 'react-string-replace';
import type { Question } from '~/common/types';
import { Input } from '~/components/ui/input';
import parse from 'html-react-parser';

interface NoteCompletionQuestionProps {
  questions: Question[];
}
export function NoteCompletionQuestion({ questions }: NoteCompletionQuestionProps) {
  const { register } = useFormContext();

  return (
    <div className="w-full">
      {questions.map((question) => (
        <div className="prose w-full !max-w-none" key={JSON.stringify(question)}>
          {parse(question.body, {
            replace: (domNode) => {
              if (domNode.type === 'text') {
                return (
                  <>
                    {reactStringReplace(domNode.data, /{{(\d+)}}/g, (match) => (
                      <Input
                        className="inline-block w-[200px] mx-1 bg-white placeholder:text-center"
                        key={match}
                        {...register(match)}
                        placeholder={match}
                      />
                    ))}
                  </>
                );
              }
            },
          })}
        </div>
      ))}
    </div>
  );
}
