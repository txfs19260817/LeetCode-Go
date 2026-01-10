
# IP Routing Lookup

**Difficulty:** Hard  
**Topics:** Trie, Bit Manipulation

---

## Metadata

- **Interview Stages:** Onsite  
- **Frequency:** ~80%  
- **Asked By:** Snowflake  
- **Last Reported:** 2 months ago  

---

## Problem Description

An **IP address** is a formatted 32-bit unsigned integer where each group of 8 bits is printed as a decimal number and the dot character `'.'` splits the groups.

- For example, the binary number `11000000 10101000 00000001 01111011` (spaces added for clarity) formatted as an IP address would be `"192.168.1.123"`.

A **CIDR block** is a format used to denote a specific set of IP addresses. It is a string consisting of a base IP address, followed by a slash, followed by a prefix length *k*. The addresses it covers are all the IPs whose **first k bits** are the same as the base IP address.

- For example, `"192.168.1.0/24"` is a CIDR block with a prefix length of *24*. Any IP address whose binary representation matches  
  `11000000 10101000 00000001 xxxxxxxx`,  
  where *x* can be either `0` or `1`, is in the set covered by the CIDR block.

---

## Input

You are given a list of routing entries, where each entry is represented as a pair `(prefix, nextHop)`:

- `prefix` is a string containing an IPv4 address and subnet mask in CIDR notation (for example, `"192.168.1.0/24"`).
- `nextHop` is a string representing the destination URL for that prefix.

You are also given a list of queries, where each `query` is a string representing an IPv4 address (for example, `"192.168.1.123"`).

---

## Output

For each `query`, determine the corresponding `nextHop` value by matching it to the **longest routing prefix** that contains it.

- If no prefix matches the query IP, return an empty string `""` for that query.
- Return a list of `nextHop` strings for the queries, in the same order as the queries.

---

## Constraints

- `1 ≤ routes.length, queries.length ≤ 10^5`
- Each `prefix` is a valid IPv4 address with a subnet mask in CIDR notation (`"a.b.c.d/x"`, `0 ≤ x ≤ 32`)
- Each `nextHop` and `query` is a non-empty string
- All IP addresses are valid IPv4 addresses

---

## Examples

### Example 1

**Input:**
```text
routes = [
  ["10.0.0.0/8", "a.com"],
  ["10.1.0.0/16", "b.com"],
  ["10.1.2.0/24", "c.com"]
]
queries = ["10.1.2.7", "10.2.3.4", "192.168.1.1"]
````

**Output:**

```text
["c.com", "a.com", ""]
```

**Explanation:**

* `"10.1.2.7"` is covered by all three prefixes, but `"10.1.2.0/24"` is the longest matching prefix, so `"c.com"` is returned.
* `"10.2.3.4"` matches only `"10.0.0.0/8"`, so `"a.com"` is returned.
* `"192.168.1.1"` does not match any prefix.

---

### Example 2

**Input:**

```text
routes = [
  ["192.168.1.0/24", "url1.com"],
  ["192.168.0.0/16", "url2.com"]
]
queries = ["192.168.1.123", "192.168.2.10"]
```

**Output:**

```text
["url1.com", "url2.com"]
```

---

### Example 3

**Input:**

```text
routes = [
  ["0.0.0.0/0", "default.com"],
  ["10.0.0.0/8", "corporate.com"],
  ["10.1.1.0/24", "subnet.com"],
  ["10.1.1.100/30", "server_farm.com"],
  ["10.1.1.101/32", "critical_server.com"]
]
queries = ["10.1.1.101", "10.1.1.100", "10.1.1.104", "10.2.0.1", "8.8.8.8"]
```

**Output:**

```text
[
  "critical_server.com",
  "server_farm.com",
  "subnet.com",
  "corporate.com",
  "default.com"
]
```
