# Find Optimal Commute

**Difficulty:** Medium  
**Company:** Databricks  
**Interview Stages:** Screening, Onsite

You live in San Francisco city and want to minimize your commute time to
the Databricks HQ.

Given a 2D matrix of the San Francisco grid and the time as well as cost
matrix of all the available transportation modes, return the fastest mode
of transportation. If there are multiple such modes then return one with
the least cost.

## Rules

1. The input grid represents the city blocks, so the commuter is only
   allowed to travel along the horizontal and vertical axes.
   Diagonal traversal is not permitted.
2. The commuter can only move to neighboring cells with the same
   transportation mode.
3. Transportation modes in the grid are numbered 1-4 where
   1 = Walk, 2 = Bike, 3 = Car, 4 = Train.

## Example

**Input:**

```
2D Grid:              Legend:
|3|3|S|2|X|           X = Roadblock
|3|1|1|2|X|           S = Source
|3|1|1|2|2|           D = Destination
|3|1|1|1|D|           1 = Walk, 2 = Bike, 3 = Car, 4 = Train
|3|3|3|3|4|
|4|4|4|4|4|

Cost Matrix (Dollars/Block): [0, 1, 3, 2]
Time Matrix (Minutes/Block): [3, 2, 1, 1]
```

**Output:**

```
"Bike"
```

Walk reaches D in 5 blocks (time 15, cost 0). Bike also reaches D in 5 blocks
(time 10, cost 5). Car and Train cannot reach D. Bike is the fastest mode.
