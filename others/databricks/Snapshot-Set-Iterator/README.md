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

Multiple iterators can coexist independently. All primary set operations are amortized O(1). Space is O(N+M) where N is total entries and M is total snapshot size across live iterators.

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

- Maintain a slice of entries, each with `{value, addVersion, removeVersion}`.
- A global version counter incremented on each add/remove.
- A hash map from value to index in the entries slice.
- `getIterator` captures the current version as `snapVersion` and builds a snapshot by walking entries: include if `addVersion < snapVersion` AND (`removeVersion == -1` OR `removeVersion >= snapVersion`).
- Re-adding a previously removed element appends a new entry to preserve correct insertion order.
