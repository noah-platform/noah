import { Controller, useFormContext } from 'react-hook-form';
import type { Question } from '~/common/types';
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from '~/components/ui/accordion';
import { Label } from '~/components/ui/label';
import { RadioGroup, RadioGroupItem } from '~/components/ui/radio-group';

interface MultipleChoiceQuestionProps {
  questions: Question[];
}
export function MultipleChoiceQuestion({ questions }: MultipleChoiceQuestionProps) {
  return (
    <div className="w-full">
      <Accordion type="single" collapsible>
        {questions.map((question) => (
          <MultipleChoice key={JSON.stringify(question)} question={question} />
        ))}
      </Accordion>
    </div>
  );
}

interface MultipleChoiceProps {
  question: Question;
}
function MultipleChoice({ question }: MultipleChoiceProps) {
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
            <RadioGroup onValueChange={field.onChange} defaultValue={field.value}>
              {question.choices.map((choice) => (
                <div className="flex items-center space-x-2" key={choice.text}>
                  <RadioGroupItem value={choice.text} id={choice.text} />
                  <Label htmlFor={choice.text} className="text-base">
                    {choice.text}
                  </Label>
                </div>
              ))}
            </RadioGroup>
          )}
        />
      </AccordionContent>
    </AccordionItem>
  );
}
