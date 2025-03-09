import { Controller, useFormContext } from 'react-hook-form';
import type { Question } from '~/common/types';
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from '~/components/ui/accordion';
import { Checkbox } from '~/components/ui/checkbox';

interface MultipleSelectQuestionProps {
  questions: Question[];
}
export function MultipleSelectQuestion({ questions }: MultipleSelectQuestionProps) {
  return (
    <div className="w-full">
      <Accordion type="single" collapsible>
        {questions.map((question) => (
          <MultipleSelect key={JSON.stringify(question)} question={question} />
        ))}
      </Accordion>
    </div>
  );
}

interface MultipleSelectProps {
  question: Question;
}
function MultipleSelect({ question }: MultipleSelectProps) {
  const { control } = useFormContext();

  return (
    <AccordionItem className="my-2" value={question.title}>
      <AccordionTrigger className="bg-gray-200 px-4 rounded-none text-base">{question.title}</AccordionTrigger>
      <AccordionContent className="p-4">
        <Controller
          control={control}
          // TODO: Add questionId
          name={question.title.slice(0, 15)}
          render={({ field }) => (
            <div className="flex flex-col gap-3">
              {question.choices.map((choice) => (
                <div key={choice.text} className="flex items-center space-x-2 my-1">
                  <Checkbox
                    id={choice.text}
                    value={choice.text}
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
                    className="text-base text-md peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
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
