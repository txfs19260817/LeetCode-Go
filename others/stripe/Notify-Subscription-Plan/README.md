# Notify Subscription Plan

Generate subscription notification emails from a reusable `schedule`.

## Parts

### Part 1

Given `user_accounts` and `schedule`, print which email should be sent on each day in chronological order.
Code: `notify_subscription_plan_part1.py`

### Part 2

Add `subscription_changes`:

- a user may change plan on a specific day
- on that day, send one extra `[Updated]` email
- all notifications after that day should use the latest plan name
Code: `notify_subscription_plan_part2.py`

### Part 3

`subscription_changes` may also contain `extension`:

- it means extending subscription by `N` days
- on that day, send one extra `[Renewed]` email
- `upcoming expiry` and `expired` notifications should be emitted according to the extended end date
Code: `notify_subscription_plan_part3.py`

## Entrypoint

`notify_subscription_plan.py` is kept as a thin wrapper to Part 3 for backward compatibility.

## Input

- `user_accounts` (or backward-compatible `users`):
  list of `{name, plan, begin_date, duration}`
- `schedule` (or backward-compatible `send_schedule`): `Map<String, String>`
- optional `subscription_changes`: list of:
  `{name|user, day|date|change_date, new_plan?, extension?}`

## Schedule rules

- `"start"`: send on `begin_date`
- `"end"`: send on final end date
- other keys: numeric string offset relative to final end date
  (example: `"-15"` means `end_date - 15`)

## Output

- print notifications by day in ascending order
- keep generation order for same-day notifications
- message format:
  `day: [Label] <name>, subscribe in plan <plan>`
