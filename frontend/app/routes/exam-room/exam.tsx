import { COOKIE_NAME, requireAuth } from '~/common/auth';
import type { Route } from './+types/exam';
import ExamContextProvider, { useExam } from './context';
import cookie from 'cookie';
import { fetchClient } from '~/clients/client';
import { Layout } from './layout';
import { ResponseType, type Question, type QuestionSet } from '../../common/types';
import { match } from 'ts-pattern';
import { RichTextPreview } from '~/components/richtext-preview';
import { Controller, useFormContext } from 'react-hook-form';
import { AudioPlayer } from './components/audio-player';
import reactStringReplace from 'react-string-replace';
import parse from 'html-react-parser';
import { Input } from '~/components/ui/input';

export async function loader({ request, params: { testId } }: Route.LoaderArgs) {
  const { sessionId } = await requireAuth(request);

  const response = await fetchClient.POST('/question-bank/v1/tests/{testID}', {
    params: { path: { testID: testId } },
    headers: { Cookie: cookie.serialize(COOKIE_NAME, sessionId) },
  });
  if (response.error) {
    throw new Error('Something went wrong');
  }
  return { exam: response.data.data };
}

export default function ExamRoom({ params, loaderData }: Route.ComponentProps) {
  const { testId } = params;
  const { exam } = loaderData;

  return (
    <ExamContextProvider testId={testId} exam={exam.test} savedAnswers={exam.answers ?? {}}>
      <Layout>
        <Exam />
      </Layout>
    </ExamContextProvider>
  );
}

function Exam() {
  const { exam, hasAudio, currentSection } = useExam();

  const section = exam.sections[currentSection];
  return (
    <div className="flex flex-col gap-4 my-8 mb-22">
      <div className="flex flex-col gap-2 bg-gray-100 p-6 rounded-xl">
        <p className="font-bold text-lg">{section.title}</p>
        <p>{section.instruction}</p>
        {section.questionSet.map((questionSet) => (
          <p key={JSON.stringify(questionSet)}>{questionSet.instruction}</p>
        ))}
      </div>
      <div className="bg-gray-100 p-6 rounded-xl">
        {hasAudio && <AudioPlayer />}
        {section.questionSet.map((questionSet) => (
          <QuestionSet key={JSON.stringify(questionSet)} questionSet={questionSet} />
        ))}
      </div>
    </div>
  );
}

interface QuesitonSetProps {
  questionSet: QuestionSet;
}
function QuestionSet({ questionSet }: QuesitonSetProps) {
  return (
    <div className="flex flex-col gap-4">
      {match(questionSet.responseType)
        .with(ResponseType.FREE_TEXT, () => <FreeTextQuestion questions={questionSet.questions} />)
        .with(ResponseType.NOTE_COMPLETION, () => <NoteCompletionQuestion questions={questionSet.questions} />)
        .otherwise(() => (
          <p>Unsupported Response Type</p>
        ))}
    </div>
  );
}
interface FreeTextQuestionProps {
  questions: Question[];
}
function FreeTextQuestion({ questions }: FreeTextQuestionProps) {
  return (
    <div>
      {questions.map((question) => (
        <div className="grid grid-cols-2 gap-4" key={JSON.stringify(question)}>
          <div className="flex flex-col gap-2 relative top-[-8px]">
            <RichTextPreview value={question.body} />
            {question.imageUrls.map((url, index) => (
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

interface NoteCompletionQuestionProps {
  questions: Question[];
}
function NoteCompletionQuestion({ questions }: NoteCompletionQuestionProps) {
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
                      <Input className="inline-block w-[200px] mx-1 bg-white" key={match} {...register(match)} />
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
