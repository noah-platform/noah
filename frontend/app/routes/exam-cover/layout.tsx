import { Link } from 'react-router';
import logo from '~/assets/logo.png';
import home from '~/assets/home.svg';
import type { ReactNode } from 'react';
import { Volume2 } from 'lucide-react';
import { Slider } from '~/components/ui/slider';

interface LayoutProps {
  hasAudio: boolean;
  volume: number;
  setVolume: (value: number) => void;
  children: ReactNode;
}
export function Layout({ hasAudio, volume, setVolume, children }: LayoutProps) {
  return (
    <div className="flex flex-col">
      <div className="flex flex-1 justify-center bg-primary">
        <div className="flex w-5/6 justify-between">
          <Link to="/" className="flex items-center gap-2">
            <img className="h-[90px]" src={logo} />
            <h1 className="text-white text-2xl font-bold">NOAH ENGLISH</h1>
          </Link>
          <div className="flex items-center gap-6">
            <Link to="/section">
              <img className="w-10 h-10" src={home} />
            </Link>
            <button className="text-md font-semibold text-white bg-red-600 hover:bg-red-700 px-4 py-2 rounded-full">
              Help
            </button>
            {hasAudio && (
              <div className="flex items-center gap-2 ml-2">
                <Volume2 className="text-white h-8 w-8" strokeWidth={2} />
                <Slider
                  className="w-[120px]"
                  value={[volume]}
                  onValueChange={([value]) => setVolume(value)}
                  max={100}
                  step={1}
                />
              </div>
            )}
          </div>
        </div>
      </div>
      <div className="flex flex-col w-5/6 mx-auto my-4">{children}</div>
    </div>
  );
}
