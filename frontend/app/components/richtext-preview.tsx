import DOMPurify from 'isomorphic-dompurify';
import { cn } from '~/lib/utils';
import { useMemo } from 'react';

interface RichTextPreviewProps {
  value: string;
  className?: string;
}
export function RichTextPreview(props: RichTextPreviewProps) {
  const { value, className } = props;
  const html = useMemo(() => {
    return DOMPurify.sanitize(value, { USE_PROFILES: { html: true } });
  }, [value]);

  return (
    <div
      className={cn(
        'prose w-full overflow-x-auto pb-1 [&>*]:h5 [&>*]:leading-normal [&>*:first-child]:mt-0 [&>*:last-child]:mb-0',
        className
      )}
      dangerouslySetInnerHTML={{
        __html: html,
      }}
    />
  );
}
