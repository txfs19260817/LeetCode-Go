from __future__ import annotations

from dataclasses import dataclass
from typing import Any, List, Optional


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


def validate(descriptor: MessageDescriptor, parsed_yaml: dict[str, Any]) -> List[str]:
    """
    Validate parsed YAML/JSON data against a protobuf-like schema.
    Returns a list of error messages.
    """
    errors: List[str] = []
    _validate_message(descriptor, parsed_yaml, "$", errors)
    return errors


def _validate_message(
    descriptor: MessageDescriptor, obj: dict[str, Any], path: str, errors: List[str]
) -> None:
    field_map = {f.name: f for f in descriptor.fields}

    for key in obj:
        if key not in field_map:
            errors.append(f"{path}.{key}: unknown field")

    for field in descriptor.fields:
        if field.name not in obj:
            if not field.optional and not field.repeated:
                errors.append(f"{path}.{field.name}: missing required field")
            continue
        
        if (value := obj[field.name]) is None and not field.optional:
            errors.append(f"{path}.{field.name}: null is not allowed")
            continue

        _validate_field(field, value, f"{path}.{field.name}", errors)


def _validate_field(field: FieldDescriptor, value: Any, path: str, errors: List[str]) -> None:
    if field.repeated:
        if not isinstance(value, list):
            errors.append(f"{path}: expected list, got {_type_name(value)}")
            return
        for i, item in enumerate(value):
            _validate_scalar_or_message(field, item, f"{path}[{i}]", errors)
        return

    _validate_scalar_or_message(field, value, path, errors)


def _validate_scalar_or_message(
    field: FieldDescriptor, value: Any, path: str, errors: List[str]
) -> None:
    ftype = field.field_type
    if ftype in ("int", "int32"):
        if not _is_int(value):
            errors.append(f"{path}: expected int, got {_type_name(value)}")
        return
    if ftype in ("float", "double"):
        if not _is_float(value):
            errors.append(f"{path}: expected float, got {_type_name(value)}")
        return
    if ftype == "string":
        if not isinstance(value, str):
            errors.append(f"{path}: expected string, got {_type_name(value)}")
        return
    if ftype == "bool":
        if not isinstance(value, bool):
            errors.append(f"{path}: expected bool, got {_type_name(value)}")
        return

    if ftype == "message":
        if not isinstance(value, dict):
            errors.append(f"{path}: expected object, got {_type_name(value)}")
            return
        if not isinstance(field.sub_descriptor, MessageDescriptor):
            errors.append(f"{path}: missing message descriptor")
            return
        _validate_message(field.sub_descriptor, value, path, errors)
        return

    errors.append(f"{path}: unknown field type '{ftype}'")


def _is_int(value: Any) -> bool:
    return isinstance(value, int) and not isinstance(value, bool)


def _is_float(value: Any) -> bool:
    return isinstance(value, (int, float)) and not isinstance(value, bool)


def _type_name(value: Any) -> str:
    return "null" if value is None else type(value).__name__


def _run_tests() -> None:
    Address = MessageDescriptor("Address", fields=[
        FieldDescriptor("street", "string"),
        FieldDescriptor("zip", "int32", optional=True),
    ])
    Person = MessageDescriptor("Person", fields=[
        FieldDescriptor("name", "string"),
        FieldDescriptor("age", "int", optional=True),
        FieldDescriptor("address", field_type="message", sub_descriptor=Address),
        FieldDescriptor("tags", "string", repeated=True),
        FieldDescriptor("scores", "double", repeated=True, optional=True),
    ])

    # 1) valid
    ok = {
        "name": "alice",
        "address": {"street": "1st ave"},
        "tags": ["ml", "compiler"],
        "scores": [1, 2.5],
    }
    assert validate(Person, ok) == []

    # 2) missing required nested
    bad1 = {"name": "bob"}
    errors = validate(Person, bad1)
    assert "$.address: missing required field" in errors

    # 3) wrong type scalar + wrong type nested field
    bad2 = {
        "name": 123,
        "address": {"street": 999},
        "tags": ["ok", 42],
    }
    errors = validate(Person, bad2)
    assert "$.name: expected string, got int" in errors
    assert "$.address.street: expected string, got int" in errors
    assert "$.tags[1]: expected string, got int" in errors

    # 4) unknown field
    bad3 = {
        "name": "c",
        "address": {"street": "x", "extra": 1},
        "unknown": True,
    }
    errors = validate(Person, bad3)
    assert "$.address.extra: unknown field" in errors
    assert "$.unknown: unknown field" in errors

    # 5) repeated provided but not a list
    bad4 = {
        "name": "d",
        "address": {"street": "x"},
        "tags": "not-a-list",
    }
    errors = validate(Person, bad4)
    assert "$.tags: expected list, got str" in errors

    print("done")


if __name__ == "__main__":
    _run_tests()
