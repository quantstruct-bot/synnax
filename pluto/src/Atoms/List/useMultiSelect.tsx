import { useEffect } from "react";
import { Key, useState } from "react";
import { useKeyHeld } from "../../Hooks/useKeys";
import SelectMultiple from "../Select/SelectMultiple";
import { TypedListEntry } from "./Types";

export interface useMultiSelectProps<
  K extends Key,
  E extends TypedListEntry<K>
> {
  data: E[];
  selected?: K[];
  selectMultiple?: boolean;
  onSelect?: (selected: K[]) => void;
}

export const useMultiSelect = <K extends Key, E extends TypedListEntry<K>>({
  data,
  selected: selectedProp,
  selectMultiple,
  onSelect: onSelectProp,
}: useMultiSelectProps<K, E>): {
  selected: K[];
  onSelect: (key: K) => void;
  clearSelected: () => void;
} => {
  const [selected, setSelected] = useState<K[]>([]);
  const [shiftSelected, setShiftSelected] = useState<K | undefined>(undefined);
  const shiftPressed = useKeyHeld("Shift");

  useEffect(() => {
    if (!shiftPressed) setShiftSelected(undefined);
  }, [shiftPressed]);

  const onSelect = (key: K) => {
    let nextSelected: K[] = [];
    if (!selectMultiple) {
      nextSelected = selected.includes(key) ? [] : [key];
    } else if (shiftPressed && shiftSelected !== undefined) {
      // We might select in reverse order, so we need to sort the indexes.
      const [start, end] = [
        data.findIndex((v) => v.key === key),
        data.findIndex((v) => v.key === shiftSelected),
      ].sort((a, b) => a - b);
      const nextKeys = data.slice(start, end + 1).map(({ key }) => key);
      // We already deselect the shiftSelected key, so we don't included it
      // when checking whether to select or deselect the entire range.
      if (
        nextKeys
          .slice(1, nextKeys.length - 1)
          .every((k) => selected.includes(k))
      ) {
        nextSelected = selected.filter((k) => !nextKeys.includes(k));
      } else {
        nextSelected = [...selected, ...nextKeys];
      }
      setShiftSelected(undefined);
    } else {
      if (shiftPressed) setShiftSelected(key);
      if (selected.includes(key))
        nextSelected = selected.filter((k) => k !== key);
      else nextSelected = [...selected, key];
    }
    setSelected(nextSelected);
    onSelectProp?.(nextSelected);
  };

  useEffect(() => {
    if (selectedProp) setSelected(selectedProp);
  }, [selectedProp]);

  const clearSelected = () => setSelected([]);
  return { selected, onSelect, clearSelected };
};
