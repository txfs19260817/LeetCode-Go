# User Features

Determine which feature flags apply to a user based on:

- Location restrictions
- A/B test participation (even vs odd user id)
- Explicit opt-in and opt-out lists

## Input

- `user`: object with `id`, `location`, optional `optIn` and `optOut`.
- `features`: list of feature definitions with optional `locations` and `abTest`.

## Output

- Set or list of feature IDs enabled for the user.

## Example

If a feature is limited to `US`, users outside the US never receive it even if
opted in.

## Constraints

- User IDs are integers.
- `optOut` overrides `optIn`.
