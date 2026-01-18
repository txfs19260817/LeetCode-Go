# Protobuf-Like JSON/YAML Validator

You are given schema classes that describe a protobuf-like message. Each field
has a name, a type, and may be optional. Some fields can be nested messages or
repeated lists of elements.

Write a validator that checks whether a parsed YAML/JSON dictionary conforms to
the schema and returns **all** validation errors.

## Schema

```python
@dataclass(frozen=True)
class FieldDescriptor:
    name: str
    field_type: str  # "int", "int32", "float", "double", "string", "bool", "message"
    optional: bool = False
    repeated: bool = False
    sub_descriptor: Optional[object] = None


@dataclass(frozen=True)
class MessageDescriptor:
    name: str
    fields: List[FieldDescriptor]
```

Rules:

- Primitive types: `int`, `float`, `string`, `bool`
- `message` expects a dictionary and uses a nested `MessageDescriptor`
- `repeated` expects a list and uses a nested `FieldDescriptor` for elements
- Optional fields may be missing or `null`
- Unknown fields in the input are errors

## Task

Implement:

```python
def validate(descriptor: MessageDescriptor, parsed_yaml: dict[str, Any]) -> List[str]:
```

Return a list of error messages. Each message should clearly identify the field
path and the reason (missing, unknown, or type mismatch).

## Example

Schema:

```text
Message: Person
  id: int
  name: string
  email: string (optional)
  tags: repeated string (optional)
  address: message Address

Message: Address
  street: string
  zip: int (optional)
```

Input:

```json
{
  "id": "42",
  "name": "Ada",
  "age": 30,
  "address": { "street": "Main", "zip": "ABC" },
  "tags": ["ok", 7]
}
```

Possible errors:

- `$.id: expected int, got string`
- `$.age: unknown field`
- `$.address.zip: expected int, got string`
- `$.tags[1]: expected string, got int`

## Notes

This is not a classic algorithm question. It is a careful simulation of schema
validation with attention to optional fields and nested/repeated structures.
