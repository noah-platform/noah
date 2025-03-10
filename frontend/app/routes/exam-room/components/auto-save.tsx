import { useEffect } from 'react';
import { useWatch } from 'react-hook-form';
import { client } from '~/clients/client';
import useDebounce from '~/hooks/useDebounce';
import { useWarnBeforeUnload } from '~/hooks/useWarnBeforeUnload';

interface AutoSaveProps {
  testId: string;
}

export function AutoSave({ testId }: AutoSaveProps) {
  const data = useWatch();
  const debouncedData = useDebounce(data, 1000);
  const { isPending, mutateAsync } = client.useMutation('put', '/question-bank/v1/tests/{testID}');

  useWarnBeforeUnload(data !== debouncedData || isPending);
  useEffect(() => {
    if (isPending) return;
    const save = async () => {
      await mutateAsync({ params: { path: { testID: testId } }, body: { answers: debouncedData } });
    };
    save();
  }, [debouncedData]);

  return null;
}
