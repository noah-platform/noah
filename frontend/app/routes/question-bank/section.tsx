import { Module } from '~/common/types';
import { useQueryState } from 'nuqs';
import { cn } from '~/lib/utils';
import { Suspense, type ReactNode } from 'react';
import { BookDashed, BookText, Headphones, Mic, Pencil } from 'lucide-react';
import { client } from '~/clients/client';
import { Link } from 'react-router';

interface ModuleButtonProps {
  module: Module;
  currentModule: Module;
  setModule: (module: Module) => void;
  icon: ReactNode;
  children: ReactNode;
}
function ModuleButton({ module, currentModule, setModule, icon, children }: ModuleButtonProps) {
  return (
    <button
      className={cn(
        'flex items-center gap-3 text-lg font-semibold text-[#D9D9D9] shadow-md hover:text-primary border-2 border-transparent hover:border-gray-300 hover:bg-gray-100 px-5 py-2.5 rounded-full cursor-pointer',
        module === currentModule ? 'bg-primary hover:bg-primary hover:text-white text-white' : ''
      )}
      onClick={() => setModule(module)}
    >
      {icon}
      {children}
    </button>
  );
}

interface ModuleSelectorProps {
  module: Module;
  setModule: (module: Module) => void;
}
function ModuleSelector({ module, setModule }: ModuleSelectorProps) {
  return (
    <div className="flex justify-center items-center gap-4">
      <ModuleButton module={Module.LISTENING} currentModule={module} setModule={setModule} icon={<Headphones />}>
        Listening
      </ModuleButton>
      <ModuleButton module={Module.READING} currentModule={module} setModule={setModule} icon={<BookText />}>
        Reading
      </ModuleButton>
      <ModuleButton module={Module.WRITING} currentModule={module} setModule={setModule} icon={<Pencil />}>
        Writing
      </ModuleButton>
      <ModuleButton module={Module.SPEAKING} currentModule={module} setModule={setModule} icon={<Mic />}>
        Speaking
      </ModuleButton>
    </div>
  );
}

interface TestListProps {
  module: Module;
}
export function TestList({ module }: TestListProps) {
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
          className="flex flex-col gap-2 w-full h-[220px] bg-gray-100 border border-gray-300 rounded-xl p-4 shadow-md"
        >
          <p className="text-xl font-medium">Test {index + 1}</p>
          <div className="border border-b border-gray-200" />
          <p className="text-sm">{test.testId}</p>
        </Link>
      ))}
    </div>
  );
}

export function TestListFallback() {
  return (
    <div className="grid grid-cols-3 gap-4">
      <div className="w-full h-[220px] bg-gray-200 animate-pulse rounded-2xl"></div>
      <div className="w-full h-[220px] bg-gray-200 animate-pulse rounded-2xl"></div>
      <div className="w-full h-[220px] bg-gray-200 animate-pulse rounded-2xl"></div>
      <div className="w-full h-[220px] bg-gray-200 animate-pulse rounded-2xl"></div>
      <div className="w-full h-[220px] bg-gray-200 animate-pulse rounded-2xl"></div>
      <div className="w-full h-[220px] bg-gray-200 animate-pulse rounded-2xl"></div>
    </div>
  );
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
