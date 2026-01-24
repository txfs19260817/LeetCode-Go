# Find Robots Position

You are given a m * n board representing a position map and an array representing distances to the nearest blocker from a robot's position. The board is a 2D array where each cell can be:

- 'O': Represents a robot.
- 'E': Represents an empty space.
- 'X': Represents a blocker.

The boundary of the board is also considered a blocker. Additionally, you are provided with a distance array of four integers, which correspond to the distances to the closest blocker in the following order: left, top, bottom, and right.

Write a function that takes the position map and the distance array as inputs and returns the indices of all robots that match the given distance criteria.

## Constraints

- The board dimensions are at least 1x1.
- The distance array contains exactly four integers.
- The matrix only contains 'O', 'E' and 'X'

## Example 1

```
Input:
board = [["O","E","E","E","X"],
         ["E","O","X","X","X"],
         ["E","E","E","E","E"],
         ["X","E","O","E","E"],
         ["X","E","X","E","X"]]
distance = [2,2,4,1]

Output: [[1,1]]

Explanation: Only the robot at position (1,1) has a distance of 2 to the left blocker, 2 to the top blocker, 4 to the bottom blocker, and 1 to the right blocker, matching the distance array.
```

## Example 2

```
Input:
board = [["O","E","X","O","O"],
         ["E","O","X","O","X"],
         ["X","X","O","E","E"],
         ["E","O","E","O","E"],
         ["O","O","X","O","O"]]
distance = [2,1,2,4]

Output: [[3,1]]
```

## Example 3

```
Input:
board = [["O","X","O"],
         ["E","O","X"],
         ["O","X","O"]]
distance = [1,1,1,1]

Output: [[2,2],[0,2]]
```
