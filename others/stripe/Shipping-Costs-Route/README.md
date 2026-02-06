# Shipping Costs Route

Routes are given as a comma-separated string in the form:

```
SOURCE:TARGET:METHOD:COST
```

Each route is directed and costs are non-negative integers.

## Part 1: Direct route with method

Given `routes`, `sourceCountry`, `targetCountry`, and `method`, return the cost
if a direct route with the given method exists.

## Part 2: Up to one middle point (ignore method)

Ignore method and find **any** valid route from source to target with at most
one intermediate country. Return:

- `route`: "A -> B -> C"
- `method`: "M1 -> M2" (or a single method if direct)
- `cost`: total cost

## Part 3: Cheapest route with up to one middle point

Same as Part 2, but return the **minimum cost** route among direct or one-stop
paths.

## Part 4: Cheapest route with any number of hops (not implemented)

Model the routes as a weighted directed graph and run Dijkstra from the source.
Track the predecessor country and method to reconstruct the path.

## Input

- `routes`: comma-separated route string.
- `sourceCountry`, `targetCountry`.
- `method` (Part 1 only).

## Output

- Part 1: cost as an integer (or "no route").
- Part 2/3: route summary with path, methods, and total cost.

## Example

**Input**

```
US:UK:UPS:4,US:UK:DHL:5,UK:CA:FedEx:10
```

**Output (Part 2/3)**

```
US -> UK -> CA, cost = 14
```
