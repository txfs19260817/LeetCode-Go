# Shipping Routes and Costs

You are given a list of shipping routes in the form
`SOURCE:TARGET:METHOD:COST`. Implement helpers to:

- Find a direct route for a given method.
- Find any two-hop route.
- Find the cheapest route (possibly multi-hop).

Return the path, methods used, and total cost.

## Input

- `routes`: comma-separated list of route records.
- `sourceCountry`, `targetCountry`: origin and destination.
- optional `method` for direct route lookup.

## Output

- For direct: cost or error message.
- For indirect: route string, methods string, and total cost.

## Example

**Input**

```
US:UK:UPS:4,US:UK:DHL:5,UK:CA:FedEx:10
```

**Output**

```
US -> UK -> CA, cost = 14
```

## Constraints

- Costs are non-negative integers.
- Routes are directed edges.
