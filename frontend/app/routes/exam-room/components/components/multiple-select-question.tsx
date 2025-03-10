import { Controller, useFormContext } from 'react-hook-form';
import type { Question } from '~/common/types';
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from '~/components/ui/accordion';
import { Checkbox } from '~/components/ui/checkbox';
import { useExam } from '../../context';
import { cn } from '~/lib/utils';

interface MultipleSelectQuestionProps {
  questions: Question[];
}
export function MultipleSelectQuestion({ questions }: MultipleSelectQuestionProps) {
  const { isReviewing } = useExam();

  const content = questions.map((question) => <MultipleSelect key={question.questionId} question={question} />);
  return (
    <div className="w-full">
      {!isReviewing ? (
        <Accordion type="single" collapsible>
          {content}
        </Accordion>
      ) : (
        <Accordion type="multiple" defaultValue={questions.map((question) => question.questionId)}>
          {content}
        </Accordion>
      )}
    </div>
  );
}

interface MultipleSelectProps {
  question: Question;
}
function MultipleSelect({ question }: MultipleSelectProps) {
  const { isReviewing, answers } = useExam();
  const { control } = useFormContext();

  return (
    <AccordionItem className="my-2" value={question.questionId}>
      <AccordionTrigger className="bg-gray-200 px-4 rounded-none text-base">{question.title}</AccordionTrigger>
      <AccordionContent className="p-4">
        <Controller
          control={control}
          name={question.questionId}
          render={({ field }) => (
            <div className="flex flex-col gap-3">
              {question.choices.map((choice) => (
                <div key={choice.text} className="flex items-center space-x-2 my-1">
                  <Checkbox
                    id={choice.text}
                    value={choice.text}
                    className="disabled:opacity-100"
                    disabled={isReviewing}
                    checked={JSON.parse(field.value ?? '[]').includes(choice.text)}
                    onCheckedChange={(checked) => {
                      const values = new Set(JSON.parse(field.value ?? '[]'));
                      if (checked) {
                        if (values.size + 1 <= question.selectCount) {
                          values.add(choice.text);
                        }
                      } else {
                        values.delete(choice.text);
                      }
                      field.onChange(JSON.stringify(Array.from(values)));
                    }}
                  />
                  <label
                    htmlFor={choice.text}
                    className={cn(
                      'text-base',
                      isReviewing && (answers[question.questionId] ?? []).includes(choice.text) && 'text-green-600',
                      isReviewing &&
                        JSON.parse(field.value ?? '[]').includes(choice.text) &&
                        !(answers[question.questionId] ?? []).includes(choice.text) &&
                        'text-red-600'
                    )}
                  >
                    {choice.text}
                  </label>
                </div>
              ))}
            </div>
          )}
        />
      </AccordionContent>
    </AccordionItem>
  );
}
