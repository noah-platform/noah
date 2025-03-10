import { useQueryState } from 'nuqs';
import { useCallback, useMemo } from 'react';

import { TestSessionStatus, type Module, type TestInfo } from '~/common/types';

export function useTagFilter(module: Module, availableTests: TestInfo[]) {
  const [selectedTags, setSelectedTags] = useQueryState<string[]>('tags', {
    parse: (value) => value.split(',').filter((tag) => tag !== ''),
    defaultValue: [],
  });

  const availableTags = useMemo(
    () =>
      Array.from(new Set(availableTests.flatMap((test) => test.tags ?? []).sort((a, b) => a.localeCompare(b)) ?? [])),
    [availableTests]
  );

  const handleSelectTag = (tag: string) => {
    setSelectedTags((tags) => {
      if (tags.includes(tag)) {
        return tags.filter((t) => t !== tag);
      }
      return [...tags, tag];
    });
  };
  const handleSelectAll = useCallback(
    (checked: boolean) => {
      setSelectedTags(checked ? availableTags : []);
    },
    [availableTags]
  );
  const handleRandomize = useCallback(() => {
    const randomTags = [...availableTags]
      .sort(() => Math.random() - 0.5)
      .slice(0, Math.floor(availableTags.length / 2));
    setSelectedTags(randomTags);
  }, [availableTags]);

  useMemo(() => {
    setSelectedTags([]);
  }, [module]);

  const filteredTests = useMemo(() => {
    const filteredTests = availableTests.filter(
      (test) => selectedTags.length === 0 || (test.tags ?? []).some((tag) => selectedTags.includes(tag))
    );
    return filteredTests.sort((a, b) => {
      if (a.status === TestSessionStatus.COMPLETED && b.status !== TestSessionStatus.COMPLETED) {
        return -1;
      }
      if (a.status !== TestSessionStatus.COMPLETED && b.status === TestSessionStatus.COMPLETED) {
        return 1;
      }
      if (a.status === TestSessionStatus.IN_PROGRESS && b.status === TestSessionStatus.NOT_STARTED) {
        return -1;
      }
      if (a.status === TestSessionStatus.NOT_STARTED && b.status === TestSessionStatus.IN_PROGRESS) {
        return 1;
      }
      return 0;
    });
  }, [availableTests, selectedTags]);

  return { availableTags, selectedTags, handleSelectTag, handleRandomize, handleSelectAll, filteredTests };
}
