---
name: problem-scaffold
description: Create a coding-problem folder scaffold with README, Go source, Go tests, and a Python solution (with __main__ asserts). Use when the user provides a problem statement and wants a new problem folder created under a specified path.
---

# Problem Scaffold

## Quick start
When the user asks to scaffold a coding problem folder:
1. Gather inputs (base path, title, description, sample I/O, implementation vs skeleton, and any missing signature details).
2. Generate the folder name and filenames using the naming rules.
3. Create the folder and files using the templates and rules below.
4. If implementation is requested, implement working Go and Python solutions and make the tests pass.
5. If skeleton is requested, include TODO stubs and keep tests skippable by default.

## Inputs to gather
- Base path for the new folder (ask each time)
- Problem title
- Problem description
- Sample I/O (at least one input + expected output)
- Implementation mode: "implementation" or "skeleton"
- Required function/class signature if not explicit in the statement
- Optional: constraints, follow-ups, metadata (difficulty, topics, etc.)

## Naming rules
- Folder name: Title-Case words joined with hyphens from the problem title
  - Example: "design card game" -> `Design-Card-Game`
- File stem: remove hyphens from the folder name
  - Example: `Design-Card-Game` -> `DesignCardGame`
- Go package name: parent folder name (usually lower-case company name)
  - Example: `uber`

## Folder structure
Create:
```
<base_path>/<FolderName>/
├── README.md
├── <FileStem>.go
├── <FileStem>_test.go
└── solution.py
```

## README.md template
Only include sections that are provided by the user; do not invent details.
```
# <Problem Title>

<Optional metadata lines if provided>

<Problem description>

## Constraints
<Constraints if provided>

## Example
**Input:**
```
<sample input>
```

**Output:**
```
<sample output>
```

<Optional explanation if provided>

## Follow-up
<Follow-up if provided>
```

## Go source file rules
- Use `package <go_package>` where `<go_package>` is the lower-case parent folder.
- Implement the function/class specified in the problem. If missing, ask for the preferred signature.
- Keep helpers private and concise.
- In skeleton mode, include TODO stubs and return zero values or empty structs.

## Go test file rules
- Use table-driven tests in `<FileStem>_test.go`.
- Prefer `github.com/stretchr/testify/assert`.
- Build cases from the provided sample I/O as well as some edge cases under the constraints.

## Python solution rules
- Match the required interface from the problem statement.
- Provide `if __name__ == "__main__":` tests with asserts built from sample I/O.

## Sanity checks
- Folder name is Title-Case with hyphens.
- Files are named correctly and live under the requested base path.
- README includes title, description, and sample I/O.
- Go package name is lower-case parent folder.
- Python tests live under `__main__`.
