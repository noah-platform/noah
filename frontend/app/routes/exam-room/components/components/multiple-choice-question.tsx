import { Controller, useFormContext } from 'react-hook-form';
import type { Question } from '~/common/types';
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from '~/components/ui/accordion';
import { Label } from '~/components/ui/label';
import { RadioGroup, RadioGroupItem } from '~/components/ui/radio-group';
import { useExam } from '../../context';
import { cn } from '~/lib/utils';

interface MultipleChoiceQuestionProps {
  questions: Question[];
}
export function MultipleChoiceQuestion({ questions }: MultipleChoiceQuestionProps) {
  const { isReviewing } = useExam();

  return (
    <div className="w-full">
      {!isReviewing ? (
        <Accordion type="single" collapsible>
          {questions.map((question) => (
            <MultipleChoice key={question.questionId} question={question} />
          ))}
        </Accordion>
      ) : (
        <Accordion type="multiple" defaultValue={questions.map((question) => question.questionId)}>
          {questions.map((question) => (
            <MultipleChoice key={question.questionId} question={question} />
          ))}
        </Accordion>
      )}
    </div>
  );
}

interface MultipleChoiceProps {
  question: Question;
}
function MultipleChoice({ question }: MultipleChoiceProps) {
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
            <RadioGroup onValueChange={field.onChange} defaultValue={field.value}>
              {question.choices.map((choice) => (
                <div className="flex items-center space-x-2" key={choice.text}>
                  <RadioGroupItem
                    value={choice.text}
                    id={choice.text}
                    disabled={isReviewing}
                    className="disabled:opacity-100"
                  />
                  <Label
                    htmlFor={choice.text}
                    className={cn(
                      'text-base',
                      isReviewing && (answers[question.questionId] ?? []).includes(choice.text) && 'text-green-600',
                      isReviewing &&
                        field.value === choice.text &&
                        !(answers[question.questionId] ?? []).includes(choice.text) &&
                        'text-red-600'
                    )}
                  >
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
