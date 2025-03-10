import { COOKIE_NAME, requireAuth } from '~/common/auth';
import type { Route } from './+types/exam';
import ExamContextProvider, { useExam } from './context';
import cookie from 'cookie';
import { fetchClient } from '~/clients/client';
import { Layout } from './layout';
import { AudioPlayer } from './components/audio-player';
import { QuestionSet } from './components/question-set';
import { cn } from '~/lib/utils';
import { RichTextPreview } from '~/components/richtext-preview';

export async function loader({ request, params: { testId } }: Route.LoaderArgs) {
  const { sessionId } = await requireAuth(request);

  const response = await fetchClient.POST('/question-bank/v1/tests/{testID}', {
    params: { path: { testID: testId } },
    headers: { Cookie: cookie.serialize(COOKIE_NAME, sessionId) },
  });
  if (response.error) {
    throw new Error('Something went wrong');
  }
  return { session: response.data.data };
}

export default function ExamRoom({ params, loaderData }: Route.ComponentProps) {
  const { testId } = params;
  const { session } = loaderData;

  return (
    <ExamContextProvider testId={testId} session={session}>
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
        {section.questionSet.length === 1 && <p>{section.questionSet[0].instruction}</p>}
      </div>
      <div className={cn('', section.passage && 'grid grid-cols-1 lg:grid-cols-2 gap-4')}>
        {section.passage && (
          <div className="bg-gray-100 p-6 rounded-xl w-full">
            <RichTextPreview value={section.passage} />
          </div>
        )}
        <div className="flex flex-col gap-4">
          {hasAudio && <AudioPlayer />}
          {section.questionSet.map((questionSet) => (
            <QuestionSet
              key={questionSet.questionSetId}
              questionSet={questionSet}
              showInstruction={section.questionSet.length > 1}
            />
          ))}
        </div>
      </div>
    </div>
  );
}
