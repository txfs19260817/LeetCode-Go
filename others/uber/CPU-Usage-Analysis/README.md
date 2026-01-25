# CPU Usage Analysis

Given a list of log events represented as a 2D string array `logs`, where each log entry contains three elements: `task_name`, an action ("enter" or "exit"), and a timestamp:

- "enter" indicates the start of a task.
- "exit" indicates the completion of a task.

Calculate the total time the CPU spends executing each task. Return the results as a list of strings in the format "task_name: total_time", sorted alphabetically by task names.

Note that the CPU is single-threaded, meaning it can only process one task at a time. Tasks may overlap, and when the CPU completes a task, if multiple tasks are waiting, it always resumes the most recently added task.

## Constraints

- Every "enter" action has a corresponding "exit" action.
- The log records are properly nested; a task cannot exit before it has entered.
- 1 <= logs.length <= 10^4
- 1 <= task_name.length <= 20
- 0 <= timestamp <= 10^9

## Example 1

**Input:** logs = [["print", "enter", "10"], ["malloc", "enter", "12"], ["malloc", "exit", "14"], ["write", "enter", "16"], ["write", "exit", "18"], ["write", "enter", "20"], ["write", "exit", "22"], ["print", "exit", "24"]]
**Output:** ["malloc: 2", "print: 8", "write: 4"]
**Explanation:**

- "malloc" runs on [12, 14], totaling 2 units of CPU time.
- "print" runs on [10, 12], [14, 16], [18, 20], and [22, 24], totaling 8 units of CPU time.
- "write" runs on [16, 18] and [20, 22], totaling 4 units of CPU time.

## Example 2

**Input:** logs = [["task1", "enter", "0"], ["task3", "exit", "6"], ["task2", "exit", "8"], ["task2", "enter", "2"], ["task3", "enter", "4"], ["task1", "exit", "10"]]
**Output:** ["task1: 4", "task2: 4", "task3: 2"]

## Example 3

**Input:** [["taskA", "enter", "0"], ["taskA", "exit", "5"], ["taskA", "enter", "6"], ["taskA", "exit", "10"], ["taskB", "enter", "10"], ["taskB", "exit", "15"]]
**Output:** ["taskA: 9", "taskB: 5"]
