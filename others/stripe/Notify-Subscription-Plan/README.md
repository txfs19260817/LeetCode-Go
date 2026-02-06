# Notify Subscription Plan

Generate subscription email notifications for users based on plan start,
expiration, and plan changes. Users receive:

- `[Welcome]` at begin date
- `[Upcoming expiration]` 15 days before expiration
- `[Expired]` at expiration date
- `[Changed]` when switching plans
- `[Renewed]` when extending duration

Plan changes can modify future expiration notices.

## Input

- `users`: list of user records with `name`, `plan`, `begin_date`, `duration`.
- `changes`: list of change records with `new_plan` or `extension`.

## Output

- Email notifications in chronological order.

## Example

**Input**

```
user A: plan X, begin 0, duration 30
change: A switches to plan Y at day 5
```

**Output**

```
0: [Welcome] A, subscribe in plan X
5: [Changed] A, subscribe in plan Y
15: [Upcoming expiration] A, subscribe in plan Y
30: [Expired] A, subscribe in plan Y
```

## Constraints

- Dates are non-negative integers.
- A change applies from its change date forward.
