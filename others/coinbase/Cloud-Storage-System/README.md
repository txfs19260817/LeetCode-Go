# Cloud Storage System

The cloud storage system should support operations to add files, copy files, and
retrieve files stored on the system.

boolean addFile(String name, int size) — should add a new file name to the
storage. size is the amount of memory required in bytes.

The current operation fails if a file with the same name already exists.
Returns true if the file was added successfully or false otherwise.

boolean copyFile(String nameFrom, String nameTo) — should copy the file at
nameFrom to nameTo.

The operation fails if nameFrom points to a file that does not exist or points
to a directory.
The operation fails if the specified file already exists at nameTo.
Returns true if the file was copied successfully or false otherwise.

int getFilesize(String name) — should return the size of the file name if it
exists, or -1 otherwise.

## Constraints

- Each file name name is a non-empty string and may include any visible ASCII
  characters (including /).
- There is no separate concept of directories; every name is treated as a
  distinct file.
- At most 10^5 operations will be performed.
- All file sizes size fit within a 32-bit signed integer.

## Example

**Input:**

```
["CloudStorage", "addFile", "copyFile", "copyFile", "addFile", "copyFile", "getFileSize", "getFileSize"]

[[], ["/dir1/dir2/file.txt", 10], ["/not-existing.file", "/dir1/file.txt"], ["/dir1/dir2/file.txt", "/dir1/file.txt"], ["/dir1/file.txt", 15], ["/dir1/file.txt", "/dir1/dir2/file.txt"], ["/dir1/file.txt"], ["/not-existing.file"]]
```

**Output:**

```
[null, true, false, true, false, false, 10, -1]
```

**Explanation:**

CloudStorage cloudStorage = new CloudStorage();
cloudStorage.addFile("/dir1/dir2/file.txt", 10); // Returns true; adds file "/dir1/dir2/file.txt" of 10 bytes.
cloudStorage.copyFile("/not-existing.file", "/dir1/file.txt"); // Returns false; the file "/not-existing.file" does not exist.
cloudStorage.copyFile("/dir1/dir2/file.txt", "/dir1/file.txt"); // Returns true; adds file "/dir1/file.txt" of 10 bytes.
cloudStorage.addFile("/dir1/file.txt", 15); // Returns false; the file "/dir1/file.txt" exists already.
cloudStorage.copyFile("/dir1/file.txt", "/dir1/dir2/file.txt"); // Returns false; the file "/dir1/dir2/file.txt" exists already.
cloudStorage.getFileSize("/dir1/file.txt"); // Returns 10. The file size.
cloudStorage.getFileSize("/not-existing.file"); // Returns -1; the file "/not-existing.file" does not exist.

## Follow-up 1: Search by Prefix and Suffix

List<String> findFile(String prefix, String suffix) — should search all stored
files for names starting with prefix and ending with suffix.

Returns a list of strings in this format:
["<name1>(<size1>)", "<name2>(<size2>)", ...]

The output is sorted by descending file size, and lexicographically for ties.
If no files match, return an empty list. Scanning all stored files is acceptable.

**Example:**

```
CloudStorage cloudStorage = new CloudStorage();
cloudStorage.addFile("/root/dir/another_dir/file.mp3", 10);
cloudStorage.addFile("/root/file.mp3", 5);
cloudStorage.addFile("/root/music/file.mp3", 7);
cloudStorage.copyFile("/root/music/file.mp3", "/root/dir/file.mp3");
cloudStorage.findFile("/root", ".mp3");
// Returns ["/root/dir/another_dir/file.mp3(10)",
//          "/root/dir/file.mp3(7)",
//          "/root/music/file.mp3(7)",
//          "/root/file.mp3(5)"]
```

## Follow-up 2: Users and Capacities

All users share the same filesystem, but each user has a capacity limit.
The user "admin" always exists and has unlimited storage. All existing addFile
calls are performed by "admin".

boolean addUser(String userId, int capacity) — add a new user with a capacity
limit. Returns false if the user already exists or userId is "admin".

int addFileBy(String userId, String name, int size) — same as addFile, but the
file is owned by userId. Returns remaining capacity if added, otherwise -1.

int updateCapacity(String userId, int capacity) — update the capacity. If the
user's total file sizes exceed the new capacity, remove the largest files
(lexicographically for ties) until the usage fits. Returns the number of
removed files, or -1 if userId does not exist.

Copying a file preserves the original file's owner and must respect that
owner's capacity.

## Follow-up 3: Compression

int compressFile(String userId, String name) — replace the file with
<name>.COMPRESSED owned by userId and of size size / 2. Returns remaining
capacity, or -1 if the file does not exist or is not owned by userId.

int decompressFile(String userId, String name) — revert a compressed file back
to the original name. Fails if the decompressed file already exists or the
user's capacity would be exceeded. Returns remaining capacity, or -1 on error.
