import { Button } from '~/components/ui/button';
import { Checkbox } from '~/components/ui/checkbox';
import { Label } from '~/components/ui/label';

interface TagSelectorProps {
  availableTags: string[];
  selectedTags: string[];
  onSelectTag: (tag: string) => void;
  onRandomize: () => void;
  onSelectAll: (checked: boolean) => void;
}

export function TagSelector({ availableTags, selectedTags, onSelectTag, onRandomize, onSelectAll }: TagSelectorProps) {
  return (
    <div className="flex flex-col gap-6 border-t border-b border-gray-300 py-6 my-3">
      <div className="flex gap-8 items-center">
        <p className="font-medium text-xl">Filter Options</p>
        <div className="flex items-center gap-6">
          <div className="flex items-center gap-2">
            <Checkbox
              checked={selectedTags.length === availableTags.length}
              onCheckedChange={(checked) => onSelectAll(checked && checked !== 'indeterminate')}
            />
            <Label className="text-lg">All</Label>
          </div>
          <Button className="rounded-full w-fit h-[32px]" variant="outline" onClick={onRandomize}>
            Random
          </Button>
        </div>
      </div>
      <div className="flex flex-wrap gap-4">
        {availableTags.map((tag) => (
          <Button
            className="rounded-full w-fit px-6 h-[32px]"
            key={tag}
            onClick={() => onSelectTag(tag)}
            variant={selectedTags.includes(tag) ? 'default' : 'outline'}
          >
            {tag}
          </Button>
        ))}
      </div>
    </div>
  );
}
