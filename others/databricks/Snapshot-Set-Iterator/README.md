# Snapshot Set Iterator

Design a `SnapshotSet` that functions as a set of integers supporting: `add`, `remove`, `contains`, `getIterator`.

When an iterator is created, it iterates over exactly the elements present at that time, in insertion order. Changes made to the set after an iterator is created do not affect that iterator.

## API

- `Add(n int) bool` — Add `n` if not present. Returns `true` if newly added, `false` if it already existed.
- `Remove(n int) bool` — Remove `n` if present. Returns `true` if removed, `false` if not found.
- `Contains(n int) bool` — Returns `true` if `n` is currently in the set.
- `GetIterator() Iterator` — Returns an iterator over the current elements in insertion order.

### Iterator

- `HasNext() bool` — Returns `true` if there are more elements.
- `Next() int` — Returns the next element. Panics if no next element.

Multiple iterators can coexist independently. All primary set operations are amortized O(1). Iterator extra space is O(1), and total space is O(N) for the entry log.

## Example

```
set := NewSnapshotSet()
set.Add(1)    // true
set.Add(2)    // true
set.Add(3)    // true
set.Add(4)    // true
set.Add(1)    // false (already exists)
it1 := set.GetIterator() // snapshot: [1,2,3,4]
set.Remove(1)  // true
set.Remove(3)  // true
set.Remove(5)  // false
it2 := set.GetIterator() // snapshot: [2,4]

Iterating it1: [1,2,3,4]
Iterating it2: [2,4]
```

## Approach

- Use a single map with tagged keys:
  - `("ent", addVersion) -> (value, removeVersion)`
  - `("idx", value) -> latest addVersion`
- A global version counter increments on each successful add/remove.
- `getIterator` captures `snapVersion` and lazily scans add versions in `[0, snapVersion)`.
- An entry is visible if `removeVersion == -1` or `removeVersion >= snapVersion`.
- Re-adding a removed value creates a new addVersion, so insertion order remains correct.
