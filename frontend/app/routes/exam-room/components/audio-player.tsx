import { useRef, useMemo, useEffect, useCallback } from 'react';
import { useExam } from '../context';

export function AudioPlayer() {
  const { exam, volume, isReviewing } = useExam();
  const player = useRef<HTMLAudioElement>(null);

  const audioUrls = useMemo(() => exam.sections.map((section) => section.audioUrl), [exam]);

  useEffect(() => {
    if (!player.current || isReviewing) return;

    player.current.src = audioUrls[0];
    player.current.play();
  }, [player]);

  useEffect(() => {
    if (!player.current) return;

    player.current.volume = Math.max(0.1, volume / 100);
  }, [volume]);

  const handleEnded = useCallback(() => {
    if (!player.current) return;

    const currentIndex = audioUrls.indexOf(player.current.src);
    if (currentIndex + 1 < audioUrls.length) {
      player.current.src = audioUrls[currentIndex + 1];
      player.current.play();
    }
  }, [player, audioUrls]);

  return <audio ref={player} onEnded={handleEnded} />;
}
