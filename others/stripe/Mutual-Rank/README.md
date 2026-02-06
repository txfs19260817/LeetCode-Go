# Mutual Rank

Each user has a ranked wishlist of other users. Implement helpers to:

- Check if a user has a mutual match at a specific rank.
- Check if a user has a mutual first choice.
- Determine which mutual pairings change if a user bumps a rank up by one.
- Determine which mutual anti-pairings change when considering ranks from the end.

## Input

- `wishlists`: mapping from user to ordered list of users.
- `username`: user to evaluate.
- `rank`: 0-based rank.

## Output

- Boolean or list of affected users depending on the function.

## Example

If `a` ranks `c` first and `c` ranks `a` first, then `a` has a mutual first
choice.

## Constraints

- Wishlists contain unique user IDs.
- Rankings are 0-based and contiguous.
