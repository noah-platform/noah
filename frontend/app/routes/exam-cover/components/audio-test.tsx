import { Play } from 'lucide-react';
import { useEffect, useRef, useState } from 'react';
import { Button } from '~/components/ui/button';

interface AudioTextProps {
  volume: number;
  next: () => void;
}
export function AudioText({ volume, next }: AudioTextProps) {
  const player = useRef<HTMLAudioElement>(null);
  const [isPlaying, setIsPlaying] = useState(false);

  useEffect(() => {
    if (!player.current) return;

    if (isPlaying) {
      player.current.play();
    } else {
      player.current.pause();
    }
  }, [isPlaying]);

  useEffect(() => {
    if (!player.current) return;

    player.current.volume = Math.max(0.1, volume / 100);
  }, [volume]);

  return (
    <div className="flex flex-col my-8">
      <h1 className="relative z-50 text-2xl font-bold mx-8">Test Sound</h1>
      <div className="relative top-[-14px] bg-gray-100 p-8 rounded-2xl">
        <p className="text-center my-3">
          Put on your headphones and click on the <span className="text-lg font-bold">Play Sound</span> button to play a
          sample sound.
        </p>
        <audio
          ref={player}
          src="https://s3-eu-west-1.amazonaws.com/oep2stt/sample-listening-multiple-choice-one-answer/sample-audio.ogg"
          loop
        />
        <div className="flex justify-center my-8">
          <Button
            className="flex items-center gap-2 min-w-[200px] h-12 text-lg font-medium"
            variant="outline"
            onClick={() => setIsPlaying((prev) => !prev)}
          >
            <Play strokeWidth={4} fill="black" />
            {!isPlaying ? 'Play Sound' : 'Stop Sound'}
          </Button>
        </div>
        <p className="text-center my-5">
          If you cannot hear the sound clearly, please check your device sound setting.
        </p>
        <div className="flex justify-center">
          <Button className="min-w-[200px] h-12 text-lg font-medium" onClick={next}>
            Continue
          </Button>
        </div>
      </div>
    </div>
  );
}
