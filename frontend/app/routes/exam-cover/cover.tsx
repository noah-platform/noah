import { fetchClient } from '~/clients/client';
import { useCallback, useState } from 'react';
import { COOKIE_NAME, requireAuth } from '~/common/auth';
import type { Route } from './+types/cover';
import { CoverState } from './types';
import { match } from 'ts-pattern';
import { ExamInstruction } from './components/instruction';
import { AudioText } from './components/audio-test';
import cookie from 'cookie';
import { Layout } from './layout';
import { LocalStorageKey, Module } from '~/common/types';

export async function loader({ request, params: { testId } }: Route.LoaderArgs) {
  const { sessionId } = await requireAuth(request);

  const response = await fetchClient.GET('/question-bank/v1/tests/{testID}', {
    params: { path: { testID: testId } },
    headers: { Cookie: cookie.serialize(COOKIE_NAME, sessionId) },
  });
  if (response.error) {
    throw new Error('Something went wrong');
  }
  return { cover: response.data.data };
}

export default function Cover({ loaderData }: Route.ComponentProps) {
  const { cover } = loaderData;
  const hasAudio = cover.module === Module.LISTENING;

  const [state, setState] = useState<CoverState>(() => (hasAudio ? CoverState.AUDIO_TEST : CoverState.INSTRUCTION));

  const [volume, setVolume] = useState(() =>
    Number(typeof window !== 'undefined' ? localStorage.getItem(LocalStorageKey.VOLUME) ?? 100 : 100)
  );
  const handleSetVolume = useCallback((volume: number) => {
    localStorage.setItem(LocalStorageKey.VOLUME, volume.toString());
    setVolume(volume);
  }, []);

  return (
    <Layout hasAudio={hasAudio} volume={volume} setVolume={handleSetVolume}>
      {match(state)
        .with(CoverState.AUDIO_TEST, () => <AudioText volume={volume} next={() => setState(CoverState.INSTRUCTION)} />)
        .with(CoverState.INSTRUCTION, () => <ExamInstruction cover={cover} />)
        .exhaustive()}
    </Layout>
  );
}
