# KV Store with Prefix Search

**Difficulty**: Medium
**Tags**: Trie, Hash Table

## Description

Design an in-memory key-value store that supports efficient value lookup by key prefix.

Implement the `KVStore` class:

- `KVStore()` Initialize an in-memory key-value store.
- `void set(String key, int value)` Insert a new `key-value` pair. Throw an error if the key already exists.
- `int get(String key)` Return the current value for `key`, or `-1` if absent.
- `void update(String key, int value)` Overwrite the value for an existing `key`. Throw an error if the key is missing.
- `void deleteKey(String key)` Remove `key` if it exists.
- `List<Integer> prefixSearch(String prefix)` Return a list of **all values** whose keys start with given `prefix`.

You may assume there are many prefix search queries, and all your methods must run in amortized _O(1)_.

**Constraints**

- 0 <= `key` <= $10^4$
- 0 <= `value` <= $10^5$

**Example 1**

> **Input:** > ["KVStore", "set", "set", "get", "update", "get", "prefixSearch", "deleteKey", "get", "prefixSearch"] > [[], ["aaple", 3], ["aap", 2], ["aap"], ["aap", 5], ["aap"], ["aap"], ["aap"], ["aap"], ["aap"], ["aaple"]]
>
> **Output:**
> [null, null, null, 2, null, 5, [5, 3], null, -1, [3]]
>
> **Explanation:**
>
> - `KVStore store = KVStore();` // initializes an empty KV store
> - `store.set("apple", 3);` // store = {"apple": 3}
> - `store.set("app", 2);` // store = {"apple": 3, "app": 2}
> - `store.get("app");` // returns 2
> - `store.update("app", 5);` // "app" value updated to 5
> - `store.get("app");` // returns 5
> - `store.prefixSearch("app");` // returns [5, 3]
> - `store.deleteKey("app");` // removes "app"; store = {"apple": 3}
> - `store.get("app");` // returns -1
> - `store.prefixSearch("app");` // returns [3]

**Example 2**

> **Input:** > ["KVStore", "set", "set", "set", "prefixSearch", "update", "prefixSearch", "deleteKey", "set", "prefixSearch", "get"] > [[], ["foo", 1], ["bar", 2], ["foobar", 3], ["foo"], ["foo", 4], ["foo"], ["bar"], ["baz", 5], ["ba"], ["bar"]]
>
> **Output:**
> [null, null, null, null, [1, 3], null, [4, 3], null, null, [5], -1]
