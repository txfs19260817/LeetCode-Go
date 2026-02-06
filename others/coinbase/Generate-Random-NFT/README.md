# Generate Random NFT

You are provided with a JSON configuration containing various traits (e.g.,
"nose", "mouth", "eyes") and their respective possible values. You are asked to
randomly generate n random Non-fungible tokens (NFTs) based on these traits.
Each NFT must select exactly one value for every trait. In this scenario,
duplicates are permitted, meaning multiple generated NFTs can have the same
combination of traits.

## Constraints

- You can assume the JSON string is valid and includes a traits field.
- The number of traits can vary, and each trait can have one or more possible
  values.
- n can be any positive integer.

## Example

**Input:**

```
config = {
  "name": "config-1",
  "size": "large",
  "traits": {
    "nose": ["pointy", "tiny", "flat"],
    "mouth": ["small", "wide", "thin"],
    "eyes": ["blue", "green", "brown"]
  }
},
n = 5
```

**Output:**

```
[
  {"nose":"pointy","mouth":"small","eyes":"blue"},
  {"nose":"tiny","mouth":"wide","eyes":"green"},
  {"nose":"flat","mouth":"thin","eyes":"brown"},
  {"nose":"pointy","mouth":"small","eyes":"green"},
  {"nose":"tiny","mouth":"thin","eyes":"blue"}
]
```

## Follow-up 1

Keep all assumptions from the previous section, but now you are asked to
generate n unique NFTs, ensuring that no two NFTs share the same combination of
traits. If it is not possible to generate the specified number of unique NFTs
based on the given n and configuration, an exception should be thrown.

### Constraints

- You can assume the JSON string is valid and includes a traits field.
- The number of traits can vary, and each trait can have one or more possible
  values.
- n can be any positive integer.
- No duplicate NFT combinations are allowed in the output.

### Example 1

**Input:**

```
config = {
  "name": "config-1",
  "size": "large",
  "traits": {
    "nose": ["pointy", "tiny", "flat"],
    "mouth": ["small", "wide", "thin"],
    "eyes": ["blue", "green", "brown"]
  }
},
n = 5
```

**Output:**

```
[
  {"nose":"pointy","mouth":"small","eyes":"blue"},
  {"nose":"tiny","mouth":"wide","eyes":"green"},
  {"nose":"flat","mouth":"thin","eyes":"brown"},
  {"nose":"pointy","mouth":"small","eyes":"green"},
  {"nose":"tiny","mouth":"thin","eyes":"blue"}
]
```

**Explanation:**
The above output contains 5 possible NFTs generated randomly using the provided
traits. Each NFT's traits combination is unique.

### Example 2

**Input:**

```
config = {
  "name": "simple",
  "size": "small",
  "traits": {
    "color": ["red", "blue", "green"],
    "shape": ["circle", "square"]
  }
},
n = 10
```

**Output:**

```
An exception was thrown to indicate that n is too large for unique combinations
```

**Explanation:**
From the given traits, we can generate at most 3 \* 2 = 6 unique combinations.

## Follow-up 2

To represent the rarity of certain traits, the configuration now includes a
weight for each trait value. The weight determines the likelihood of selecting
a specific trait value. The probability of picking a particular value is
calculated as: probability = weight / (sum of all weights for the trait).

For example:

```
{
  ...
  "traits": {
    "nose": [
      { "name": "pointy", "weight": 1 },
      { "name": "tiny", "weight": 2 },
      { "name": "flat", "weight": 3 }
    ]
  }
}
```

For the trait "nose", the probabilities of each value being generated are:

- "pointy": 1 / (1 + 2 + 3) = 0.167 (16.7%)
- "tiny": 2 / (1 + 2 + 3) = 0.333 (33.3%)
- "flat": 3 / (1 + 2 + 3) = 0.5 (50%)

You are asked to generate n unique NFTs, ensuring that no two NFTs have
identical combinations of traits. If it is not possible to generate the
specified number of unique NFTs with the given configuration, an exception
should be thrown.

### Constraints

- Each trait's weight is an integer in the range [1, 1000].
- No duplicate NFT combinations are allowed in the output.

### Example

**Input:**

```
config = {
  "name": "config-1",
  "size": "large",
  "traits": {
    "nose": [
      {"name": "pointy", "weight": 1},
      {"name": "tiny", "weight": 2},
      {"name": "flat", "weight": 3}
    ],
    "mouth": [
      {"name": "small", "weight": 1000},
      {"name": "wide", "weight": 1},
      {"name": "thin", "weight": 1}
    ],
    "eyes": [
      {"name": "blue", "weight": 10},
      {"name": "green", "weight": 2},
      {"name": "brown", "weight": 1}
    ]
  }
},
n = 5
```

**Output:**

```
[
  {"nose": "pointy", "mouth": "small", "eyes": "blue"},
  {"nose": "tiny", "mouth": "small", "eyes": "brown"},
  {"nose": "tiny", "mouth": "small", "eyes": "green"},
  {"nose": "flat", "mouth": "small", "eyes": "blue"},
  {"nose": "tiny", "mouth": "small", "eyes": "blue"}
]
```

**Explanation:**
The above output contains 5 possible NFTs generated randomly, and each NFT's
trait combination is unique. The "mouth" trait with value "small" has a much
larger weight, making it appear more frequently than "wide" or "thin".
