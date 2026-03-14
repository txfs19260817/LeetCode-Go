# Encode & Decode Morse Code

- Company: Google
- Stage: Onsite
- Type: Interview
- Reported By: Anonymous User
- Reported On: Jun 10, 2020

Assuming a mapping for morse code to char (and vice versa) is available, write two functions:

- `encode`: Encodes a regular string into a morse code string.
- `decode`: Returns all possible decoded message strings for a given morse code string.

Hypothetical mapping:

```text
{ 'A': '.', 'B': '-', 'C': '.-' }
```

## Example
**Input:**
```text
encode("AB") with { 'A': '.', 'B': '-', 'C': '.-' }
encode("C") with { 'A': '.', 'B': '-', 'C': '.-' }
decode(".-") with { '.': 'A', '-': 'B', '.-': 'C' }
```

**Output:**
```text
".-"
".-"
["AB", "C"]
```
