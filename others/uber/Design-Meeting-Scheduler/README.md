# Design Meeting Scheduler

Design a meeting room scheduler for a predefined set of meeting rooms identified by their unique room IDs (e.g., `["roomA", "roomB", ...]`).

Implement the `MeetingScheduler` class:

- `MeetingScheduler(List<String> roomList)` Initializes the scheduler system with the given list of room IDs.
- `String schedule(int start, int end)` Schedule a meeting in one of the available rooms between the start and end time.
  - If multiple rooms are available, assign the room with the lexicographically smallest room ID.
  - If no rooms are available for the specified time slot, return an empty string.

## Constraints

- The number of rooms will be between 1 and 10^4.
- `start` and `end` times are integers where `0 <= start < end <= 10^9`.
- The number of `schedule` calls will not exceed 10^5.

## Example 1

**Input:**
```
["MeetingScheduler", "schedule", "schedule", "schedule", "schedule", "schedule", "schedule"]
[["roomB", "roomA", "roomC"], [1, 5], [1, 5], [2, 6], [2, 3], [5, 10], [8, 10]]
```

**Output:**
```
[null ,"roomA", "roomB", "roomC", "", "roomA", "roomB"]
```

**Explanation:**

```
MeetingScheduler scheduler = new MeetingScheduler(rooms) // Initialize with rooms ["roomB", "roomA", "roomC"].
scheduler.schedule(1, 5) // Return "roomA", because "roomA" is available and alphabetically first.
scheduler.schedule(1, 5) // Return "roomB".
scheduler.schedule(2, 6) // Return "roomC".
scheduler.schedule(2, 3) // Return "", since no rooms are available.
scheduler.schedule(5, 10) // Return "roomA" as its previous meeting has ended, making it available again.
scheduler.schedule(8, 10) // Return "roomB".
```

## Example 2

**Input:**
```
["MeetingScheduler", "schedule", "schedule", "schedule", "schedule"]
[["roomA"], [1, 5], [1, 5], [5, 10], [2, 6]]
```

**Output:**
```
[null, "roomA", "",  "roomA", ""]
```

## Example 3

**Input:**
```
["MeetingScheduler", "schedule", "schedule", "schedule", "schedule", "schedule", "schedule"]
[["roomA", "roomB"], [1, 3], [2, 4], [3, 5], [1, 2], [4, 6], [8, 10]]
```

**Output:**
```
[null, "roomA", "roomB", "roomA", "roomB", "roomB", "roomA"]
```
