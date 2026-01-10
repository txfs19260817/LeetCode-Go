# Dynamic LRU Cache

**Difficulty:** Medium  
**Topics:** Hash Table, Design, Linked List

**Interview Stages:** Onsite  
**Frequency:** 80%  
**Asked By:** Snowflake  
**Last Reported:** 15 hours ago

*(This question is a variation of the LeetCode question [146. LRU Cache](https://leetcode.com/problems/lru-cache/description/). If you haven't completed that question yet, it is recommended to solve it first.)*

Design a data structure that implements a **Least Recently Used (LRU) Cache** with support for **dynamic resizing**. The cache should allow retrieval, insertion, and capacity adjustment while maintaining the LRU eviction policy.

Implement the `LRUCache` class:

- `LRUCache(int capacity)` Initializes the cache with the given capacity.
- `int get(int key)` Returns the value associated with `key` if present; otherwise, returns `-1`.
- `void put(int key, int value)` Inserts or updates the key with the given value. If the cache exceeds capacity, evict the least recently used item.
- `void resize(int newCapacity)` Changes the cache's capacity to `newCapacity`.
    - If `newCapacity` is larger than the current number of items, no eviction occurs.
    - If `newCapacity` is smaller, repeatedly evict the least recently used items until the number of items fits within the new capacity.

The functions `get` and `put` must each run in `O(1)` average time complexity.

**Constraints:**
- `1 <= capacity <= 10^4`
- `0 <= key <= 10^4`
- `0 <= value <= 10^4`
- `1 <= newCapacity <= 10^4`
- At most `2 * 10^5` calls to `get`, `put`, and `resize` combined.

**Example:**

> **Input:**  
> `["LRUCache", "put", "put", "put", "get", "put", "get", "resize", "put", "put", "resize", "get", "get", "put", "get"]`  
> `[[3], [1, 1], [2, 2], [3, 3], [2], [4, 4], [1], [5], [5, 5], [6, 6], [2], [5], [4], [6, 66], [6]]`
>
> **Output:**  
> `[null, null, null, null, 2, null, -1, null, null, null, null, 5, -1, null, 66]`
>
> **Explanation:**
> - `LRUCache cache = new LRUCache(3);`   // Initialize cache with capacity 3
> - `cache.put(1, 1);`                    // Cache is {1=1}
> - `cache.put(2, 2);`                    // Cache is {1=1, 2=2}
> - `cache.put(3, 3);`                    // Cache is {1=1, 2=2, 3=3}
> - `cache.get(2);`                       // Returns 2. Cache order: {2=2, 1=1, 3=3}
> - `cache.put(4, 4);`                    // Evicts key 1 (LRU). Cache is {2=2, 3=3, 4=4}
> - `cache.get(1);`                       // Returns -1 (key 1 was evicted)
> - `cache.resize(5);`                    // Resize to capacity 5. Cache remains {2=2, 3=3, 4=4}
> - `cache.put(5, 5);`                    // Cache is {2=2, 3=3, 4=4, 5=5}
> - `cache.put(6, 6);`                    // Cache is {2=2, 3=3, 4=4, 5=5, 6=6}
> - `cache.resize(2);`                    // Shrinks capacity to 2. Evicts keys 2, 3, 4. Cache is {5=5, 6=6}
> - `cache.get(5);`                       // Returns 5
> - `cache.get(4);`                       // Returns -1 (key 4 was evicted)
> - `cache.put(6, 66);`                   // Updates key 6 value to 66. Cache is {5=5, 6=66}
> - `cache.get(6);`                       // Returns 66  
