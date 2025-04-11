import { Module, TestSessionStatus } from '~/common/types';
import { useQueryState } from 'nuqs';
import { Suspense } from 'react';
import { BookDashed } from 'lucide-react';
import { client } from '~/clients/client';
import { Link } from 'react-router';
import { requireAuth } from '~/common/auth';
import type { Route } from '../+types';
import { ModuleSelector } from './components/module-selector';
import { match } from 'ts-pattern';
import { TagSelector } from './components/tag-selector';
import { useTagFilter } from './hooks/useTagFilter';

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
  const tests = data?.data ?? [];
  const { availableTags, selectedTags, handleSelectTag, handleRandomize, handleSelectAll, filteredTests } =
    useTagFilter(module, tests);

  return (
    <div className="flex flex-col gap-4">
      <TagSelector
        availableTags={availableTags}
        selectedTags={selectedTags}
        onSelectTag={handleSelectTag}
        onRandomize={handleRandomize}
        onSelectAll={handleSelectAll}
      />
      {filteredTests.length === 0 && (
        <div className="flex flex-col gap-4 justify-center items-center text-lg min-h-[400px]">
          <BookDashed className="w-22 h-22" strokeWidth={1} />
          No test available
        </div>
      )}
      {filteredTests.length > 0 && (
        <div className="grid grid-cols-3 gap-4">
          {filteredTests.map((test, index) => (
            <Link
              to={test.status !== TestSessionStatus.COMPLETED ? `/test/${test.testId}` : `/test/${test.testId}/start`}
              key={test.testId}
              className="flex flex-col gap-2 w-full h-[220px] bg-[#1310A5] border-2 border-gray-300 hover:border-2 hover:border-gray-400 rounded-xl p-4 shadow-sm"
            >
              <div className="flex justify-between items-center gap-2">
                <p className="text-xl font-medium text-white">Test {index + 1}</p>
                {match(test.status as TestSessionStatus)
                  .with(TestSessionStatus.IN_PROGRESS, () => (
                    <p className="px-4 py-1 border border-yellow-800 bg-yellow-600 text-white text-sm rounded-full">
                      In Progress
                    </p>
                  ))
                  .with(TestSessionStatus.COMPLETED, () => (
                    <p className="px-4 py-1 border border-green-800 bg-green-700 text-white text-sm rounded-full">
                      Completed
                    </p>
                  ))
                  .otherwise(() => null)}
              </div>
              <div className="border border-b border-gray-200" />
              {test.tags ?? [] ? (
                <div className="flex flex-wrap gap-2">
                  {(test.tags ?? []).map((tag) => (
                    <p key={tag} className="px-3 py-1 bg-gray-200 text-sm rounded-full">
                      {tag}
                    </p>
                  ))}
                </div>
              ) : (
                <p className="text-gray-500 text-sm">No tags</p>
              )}
            </Link>
          ))}
        </div>
      )}
    </div>
  );
}

function TestListFallback() {
  return (
    <div className="flex flex-col gap-6">
      <div className="w-full h-[160px] bg-gray-100 animate-pulse rounded-2xl"></div>
      <div className="grid grid-cols-3 gap-4">
        <div className="w-full h-[220px] bg-gray-100 animate-pulse rounded-2xl"></div>
        <div className="w-full h-[220px] bg-gray-100 animate-pulse rounded-2xl"></div>
        <div className="w-full h-[220px] bg-gray-100 animate-pulse rounded-2xl"></div>
        <div className="w-full h-[220px] bg-gray-100 animate-pulse rounded-2xl"></div>
        <div className="w-full h-[220px] bg-gray-100 animate-pulse rounded-2xl"></div>
        <div className="w-full h-[220px] bg-gray-100 animate-pulse rounded-2xl"></div>
      </div>
    </div>
  );
}
