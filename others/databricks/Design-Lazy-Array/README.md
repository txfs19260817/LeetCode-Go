# Design Lazy Array

Design a class `LazyArray` that supports deferred operations on an integer array. The class allows chaining multiple transformations via `map` and `filter`, but defers execution until `indexOf` is called.

## API

- `LazyArray(arr)` — initializes with the given integer array.
- `Map(fn)` — returns a **new** `LazyArray` that accumulates `fn` into its pipeline without applying it.
- `Filter(predicate)` — returns a **new** `LazyArray` that accumulates a predicate into its pipeline. Elements that do not satisfy the predicate are skipped.
- `IndexOf(target) → int` — applies all accumulated operations in order to each original element and returns the 0-based index in the original array of the first transformed value equal to `target`. Returns `-1` if not found.

It is guaranteed that any valid query yields at most one matching index.

## Example

```
arr = [10, 20, 30, 40, 50]

arr.map(n*2)        → [20, 40, 60, 80, 100]   → indexOf(40)  = 1
arr.map(n*2).map(n*3) → [60, 120, 180, 240, 300] → indexOf(240) = 3

arr2 = [1, 2, 3, 4, 5]
arr2.map(n+10)      → [11, 12, 13, 14, 15]    → indexOf(100) = -1

arr3 = [1, 2, 3, 4, 5]
arr3.filter(n%2==1) → [1, 3, 5]               → indexOf(3) = 1
arr3.map(n*3).filter(n%2==0) → [6, 12]        → indexOf(12) = 3
arr3.filter(n%2==1).map(n*10) → [10, 30, 50]  → indexOf(50) = 4
```
