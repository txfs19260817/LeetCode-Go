# Extract Cards

You are given an encoded string that represents a list of card IDs.
Each card is encoded as a two-digit length followed by the card string.

After extracting the cards, filter them against a list of supported cards.
If an extracted card is a generic prefix of a supported card, include the
supported card once.

## Input

- `card_string`: string of repeated `[len][card]` segments.
- `supported_cards`: list of supported card IDs.

## Output

- List of supported cards inferred from the input.

## Example

**Input**
```
card_string = "03abc02de04wxyz"
supported_cards = ["abc", "wxyz"]
```

**Output**
```
["abc", "wxyz"]
```

## Constraints

- Length prefixes are two digits.
- Card strings contain no digits.
