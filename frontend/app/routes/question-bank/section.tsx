import { Module } from '~/common/types';
import { useQueryState } from 'nuqs';
import { Suspense } from 'react';
import { BookDashed } from 'lucide-react';
import { client } from '~/clients/client';
import { Link } from 'react-router';
import { requireAuth } from '~/common/auth';
import type { Route } from '../+types';
import { ModuleSelector } from './components/module-selector';

export async function loader({ request }: Route.LoaderArgs) {
  await requireAuth(request);
}

export default function QuestionBank() {
  const [module, setModule] = useQueryState<Module>('module', {
    parse: (value) => Module[value as keyof typeof Module],
    defaultValue: Module.LISTENING,
  });

  return (
    <div className="flex flex-col gap-8 my-4">
      <ModuleSelector module={module} setModule={setModule} />
      <Suspense fallback={<TestListFallback />}>
        <TestList module={module} />
      </Suspense>
    </div>
  );
}

interface TestListProps {
  module: Module;
}
function TestList({ module }: TestListProps) {
  const { data } = client.useSuspenseQuery('get', '/question-bank/v1/tests', {
    params: { query: { module } },
  });
  const lists = data?.data ?? [];

  if (lists.length === 0) {
    return (
      <div className="flex flex-col gap-4 justify-center items-center text-lg min-h-[400px]">
        <BookDashed className="w-22 h-22" strokeWidth={1} />
        No test available
      </div>
    );
  }
  return (
    <div className="grid grid-cols-3 gap-4">
      {lists.map((test, index) => (
        <Link
          to={`/test/${test.testId}`}
          key={test.testId}
          className="flex flex-col gap-2 w-full h-[220px] bg-gray-100 border-2 border-gray-300 hover:border-2 hover:border-gray-400 rounded-xl p-4 shadow-sm"
        >
          <div className="flex justify-between items-center gap-2">
            <p className="text-xl font-medium">Test {index + 1}</p>
            {test.status === 'IN_PROGRESS' && (
              <p className="px-4 py-1 border-2 border-primary bg-gray-200 text-sm rounded-full">In Progress</p>
            )}
          </div>
          <div className="border border-b border-gray-200" />
          <p className="text-sm">{test.testId}</p>
        </Link>
      ))}
    </div>
  );
}

function TestListFallback() {
  return (
    <div className="grid grid-cols-3 gap-4">
      <div className="w-full h-[220px] bg-gray-100 animate-pulse rounded-2xl"></div>
      <div className="w-full h-[220px] bg-gray-100 animate-pulse rounded-2xl"></div>
      <div className="w-full h-[220px] bg-gray-100 animate-pulse rounded-2xl"></div>
      <div className="w-full h-[220px] bg-gray-100 animate-pulse rounded-2xl"></div>
      <div className="w-full h-[220px] bg-gray-100 animate-pulse rounded-2xl"></div>
      <div className="w-full h-[220px] bg-gray-100 animate-pulse rounded-2xl"></div>
    </div>
  );
}
