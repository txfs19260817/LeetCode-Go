# Print Directory with Indentation

## Description

Given the root of a directory tree, print the directory structure in the following format, using indentation to represent depth:

```
dir
 subdir1
  file1.ext
  subsubdir1
 subdir2
  subsubdir2
   file2.ext
```

Each level of depth adds one leading space.

## Requirements

Implement a function that returns the directory listing as a slice of strings (one per line), in preorder (parent before children). Children should be printed in the order they appear in the input tree.

## Example

Tree:

- dir
  - subdir1
    - file1.ext
    - subsubdir1
  - subdir2
    - subsubdir2
      - file2.ext

Output:

```
dir
 subdir1
  file1.ext
  subsubdir1
 subdir2
  subsubdir2
   file2.ext
```
