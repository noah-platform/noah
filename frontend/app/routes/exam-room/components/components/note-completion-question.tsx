import { Controller, useFormContext } from 'react-hook-form';
import reactStringReplace from 'react-string-replace';
import type { Question } from '~/common/types';
import { Input } from '~/components/ui/input';
import parse from 'html-react-parser';
import { useExam } from '../../context';
import { cn } from '~/lib/utils';

interface NoteCompletionQuestionProps {
  questions: Question[];
}
export function NoteCompletionQuestion({ questions }: NoteCompletionQuestionProps) {
  const { isReviewing, answers } = useExam();
  const { control } = useFormContext();

  return (
    <div className="w-full">
      {questions.map((question) => (
        <div className="prose w-full !max-w-none" key={question.questionId}>
          {parse(question.body, {
            replace: (domNode) => {
              if (domNode.type === 'text') {
                return (
                  <>
                    {reactStringReplace(domNode.data, /{{(\d+)}}/g, (match) => (
                      <Controller
                        control={control}
                        name={`${question.questionId}-${match}`}
                        render={({ field }) =>
                          !isReviewing ? (
                            <Input
                              className={
                                'inline-block w-[200px] mx-1 bg-white placeholder:text-center disabled:opacity-100'
                              }
                              key={match}
                              disabled={isReviewing}
                              value={field.value}
                              onChange={field.onChange}
                              placeholder={match}
                            />
                          ) : (
                            <p
                              className={cn(
                                'h-9 inline-block w-[200px] rounded-sm border border-input bg-transparent px-3 py-1 text-base',
                                (answers[`${question.questionId}-${match}`] ?? []).includes(field.value) &&
                                  'border-2 border-green-600 text-green-600',
                                !(answers[`${question.questionId}-${match}`] ?? []).includes(field.value) &&
                                  'border-2 border-red-600'
                              )}
                            >
                              <span
                                className={cn(
                                  !(answers[`${question.questionId}-${match}`] ?? []).includes(field.value) &&
                                    'line-through'
                                )}
                              >
                                {field.value}
                              </span>
                              {!(answers[`${question.questionId}-${match}`] ?? []).includes(field.value) && (
                                <span className="text-green-600 pl-1">
                                  {answers[`${question.questionId}-${match}`]}
                                </span>
                              )}
                            </p>
                          )
                        }
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
