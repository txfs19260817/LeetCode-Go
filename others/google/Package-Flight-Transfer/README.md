# Package Flight Transfer

## Problem

Given an origin airport, a destination airport, and a series of flights, determine whether it is possible for a package to travel from the origin to the destination.

Each flight is represented by:

- departure airport
- arrival airport
- departure time
- arrival time

The package starts at the origin airport at time `0`.

During transportation, the package may take a flight only if the flight's departure time is greater than or equal to the time the package arrives at that airport.

Determine whether the package can be transferred from `s` to `t`.

## Example 1

**Input:**
```text
origin = "NYC"
destination = "SFO"
flights = [
    ("NYC", "LAX", 0, 4),
    ("LAX", "SFO", 5, 7),
]
```

**Output:**
```text
True
```

## Example 2

**Input:**
```text
origin = "NYC"
destination = "SFO"
flights = [
    ("NYC", "LAX", 0, 4),
    ("LAX", "SFO", 3, 5),
]
```

**Output:**
```text
False
```
