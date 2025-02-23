import { Headphones, BookText, Pencil, Mic } from 'lucide-react';
import { Module } from '~/common/types';
import type { ReactNode } from 'react';
import { cn } from '~/lib/utils';

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
export function ModuleSelector({ module, setModule }: ModuleSelectorProps) {
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
