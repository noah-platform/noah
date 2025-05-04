import { Module, Section } from '~/common/types';
import { cn } from '~/lib/utils';
import { Suspense, type ReactNode } from 'react';
import { TableOfContents } from 'lucide-react';
import listening from '../listeningicon.svg';
import writing from '../writingicon.svg';
import reading from '../readingicon.svg';
import speaking from '../speakingicon.svg';
import { useState } from 'react';
import { requireAuth } from '~/common/auth';
import type { Route } from './+types/review';
import { client } from '~/clients/client';
import { Link } from 'react-router';
import dayjs from 'dayjs';
import { match } from 'ts-pattern';

export async function loader({ request }: Route.LoaderArgs) {
  await requireAuth(request);
}

interface ModuleButtonProps {
  section: Section;
  currentSection: Section;
  setSection: (section: Section) => void;
  icon: ReactNode;
  children: ReactNode;
}
function SectionButton({ section, currentSection, setSection, icon, children }: ModuleButtonProps) {
  return (
    <button
      className={cn(
        'flex items-center gap-3 text-lg font-semibold text-[#D9D9D9] shadow-md hover:text-primary border-2 border-transparent hover:border-gray-300 hover:bg-gray-100 px-5 py-2.5 rounded-full cursor-pointer',
        section === currentSection ? 'bg-primary hover:bg-primary hover:text-white text-white' : ''
      )}
      onClick={() => setSection(section)}
    >
      {icon}
      {children}
    </button>
  );
}

interface ModuleSelectorProps {
  section: Section;
  setSection: (section: Section) => void;
}
function ModuleSelector({ section, setSection }: ModuleSelectorProps) {
  return (
    <div>
      <div className="flex justify-center items-center gap-4">
        {/* <SectionButton section={Section.FULLTEST} currentSection={section} setSection={setSection} icon={<Globe />}>
          Full Test
        </SectionButton> */}
        <SectionButton
          section={Section.SECTIONTEST}
          currentSection={section}
          setSection={setSection}
          icon={<TableOfContents />}
        >
          Section Test
        </SectionButton>
      </div>
      <div className="border border-b border-gray-200 mt-8 mb-2" />
    </div>
  );
}

const Reviewbar = () => {
  const {
    data: { data: sections },
  } = client.useSuspenseQuery('get', '/question-bank/v1/sessions');

  // const sections = [
  //   {
  //     name: 'Listening',
  //     testNumber: 1,
  //     icon: <img className="w-12 h-12" src={listening} alt="Listening" />,
  //     score: 6.5,
  //   },
  //   { name: 'Writing', testNumber: 2, icon: <img className="w-12 h-12" src={writing} alt="Writing" />, score: 7.0 },
  //   { name: 'Reading', testNumber: 3, icon: <img className="w-12 h-12" src={reading} alt="Reading" />, score: 8.5 },
  //   { name: 'Speaking', testNumber: 4, icon: <img className="w-12 h-12" src={speaking} alt="Speaking" />, score: 7.5 },
  //   { name: 'Fulltest', testNumber: 1, icon: <img className="w-12 h-12" src={globe} alt="globe" />, score: 7.5 },
  // ];
  //
  // const maxScore = 9.0; // Max score for all sections

  return (
    <div className="w-full">
      {sections.map((section, index) => (
        <div
          key={section.test.testId}
          className="flex items-center justify-between p-4 border rounded-lg shadow-md mb-4 bg-white transform transition-transform hover:scale-105 hover:shadow-lg"
        >
          <div className="flex items-center gap-4">
            <div className="text-3xl">
              {match(section.test.module as Module)
                .with(Module.LISTENING, () => <img className="w-12 h-12" src={listening} alt="Listening" />)
                .with(Module.WRITING, () => <img className="w-12 h-12" src={writing} alt="Writing" />)
                .with(Module.READING, () => <img className="w-12 h-12" src={reading} alt="Reading" />)
                .with(Module.SPEAKING, () => <img className="w-12 h-12" src={speaking} alt="Speaking" />)
                .exhaustive()}
            </div>
            <div>
              <h2 className="text-xl font-semibold">
                {section.test.module} / Test {sections.length - index}
              </h2>
              <p className="text-sm text-gray-500">{dayjs(section.startedAt).format('DD MMM YYYY HH:mm')}</p>
            </div>
          </div>
          <div className="flex items-center gap-2">
            <div className="rounded-full px-4 py-2 text-lg text-black" style={{ backgroundColor: '#D9D9D9' }}>
              7.5/9
            </div>
            <Link to={`/test/${section.test.testId}/start`}>
              <button className="rounded-lg px-5 py-2 text-black" style={{ backgroundColor: '#D9D9D9' }}>
                Review
              </button>
            </Link>
          </div>
        </div>
      ))}
    </div>
  );
};

export default function QuestionBank() {
  const [section, setSection] = useState<Section>(Section.SECTIONTEST);
  return (
    <div className="flex flex-col gap-8 my-4">
      <ModuleSelector section={section} setSection={setSection} />
      <Suspense>
        <Reviewbar />
      </Suspense>
    </div>
  );
}
