# Message Parser (Type Size)

Design a message parser `MessageParser` that reads a schema definition and can:

- Return the field type in the top-level `Message` struct.
- Compute the byte size of primitives, custom structs, or field names.

The schema format is:

- A struct starts with `TypeName :` (colon at the end).
- Following lines are fields: `field_name field_type`.
- Multiple structs may appear, separated by blank lines.

## Example

```text
primitives = {"int": 4, "float": 8, "char": 1}

schema =
Point :
x int
y int

Message :
id int
loc Point
weight float
```

Expected:

- `GetType("loc") -> "Point"`
- `GetSize("int") -> 4`
- `GetSize("Point") -> 8`
- `GetSize("loc") -> 8`
- `GetSize("Message") -> 20`

## Approach

Use DFS to compute struct sizes and memoize each computed type size to avoid
repeated work. A `messageFields` map provides O(1) lookup for `GetType`.

## Complexity

- Parsing: O(total lines)
- Size query with memoization: O(total fields) across all types
