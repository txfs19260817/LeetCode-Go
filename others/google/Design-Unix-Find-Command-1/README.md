# Design Unix Find Command+1

- Difficulty: Hard
- Companies: Google, Amazon
- Stages: Onsite
- Asked By: Google, Amazon
- Source: https://www.hack2hire.com/companies/google/coding-questions/683e7b308d3dfb7c9bd756aa/practice?questionId=683e86418d3dfb7c9bd756ab

## Problem

Design Unix Find Command
+1
Hard
Recursion
Interview Stages
Onsite
Frequency
Asked By
Last Reported
1 months ago

Design a miniature Unix find command inside an in-memory file system. The system should allow users to add files (with sizes and extensions) and later search any directory and sub-directory using size, extension, or name-prefix filters.

Given the following abstract Filter class:

abstract class Filter {
    // Returns true if the given file satisfies this filter's criterion
    abstract boolean apply(File file);
}

You need to implement three filter types and return a list of filenames that meet the provided filtering criteria. Your design should be extensible, allowing easy support for additional filter types in the future.

The system must support the following filter types:

Size Filter: Matches files whose size exceeds the threshold derived from filterParam. Valid units are "KB", "MB", and "GB".

For example, a filter with filterParam of "2MB" means files with size "5MB" and "1GB" are qualified, but "500KB" is not.

Extension Filter: Matches files whose extension (the substring after the dot) exactly equals filterParam.

For example, if filterParam is "xml", then "config.xml" qualifies, but "config.yml" does not.

Prefix Filter: Matches files whose base name (the substring before the dot) starts with filterParam.

For example, a filterParam of "doc" matches "doc.pdf" and "document.txt", but not "adoc.pdf".

Implement the FileSystem class:

FileSystem() Create an empty file system whose root path is "/".
String addFile(String fileName, String size, String path) Add a new file named fileName of the given size into a directory path.
On success, return the original fileName.
If a file with the same full name already exists in that directory, return an empty string.
Directory components that do not yet exist must be created automatically.
List<String> applyFilter(String path, String filterType, String filterParam) Recursively list all files under path that satisfy the filter. And you may return the qualified filenames in any order.
If filterType is "SIZE", then filterParam is a size string (e.g., "5MB"). Return files whose size is strictly greater than that value.
If filterType is "EXTENSION", then filterParam is an extension string (e.g., "xml"). Return files whose extension equals filterParam.
If filterType is "PREFIX", then filterParam is a filename prefix (ignore the extension). Return files whose base name starts with filterParam.

Constrains

File and directory names contain only lowercase letters plus exactly one dot (".") in each file name, separating name and extension.
Valid size units are "KB", "MB", "GB", and all size values are positive integers.
Path "/" denotes the root directory.

Example

Input:
["FileSystem","addFile","addFile","addFile","addFile","addFile","applyFilter","applyFilter","applyFilter","addFile"]

[[],["doc.pdf","5MB","/home/user/docs"],["img.jpg","2MB","/home/user/photos"],["script.sh","10KB","/home/user/bin"],["config.xml","1KB","/home/user/config"],["backup.zip","100MB","/home/user/backup"],["/home/user","SIZE","1MB"],["/home","EXTENSION","xml"],["/","PREFIX","doc"],["doc.pdf","10MB","/home/user/docs"]]

Output:
[null,"doc.pdf","img.jpg","script.sh","config.xml","backup.zip",["backup.zip","doc.pdf","img.jpg"],["config.xml"],["doc.pdf"],""]

Explanation:

FileSystem fs = new FileSystem();
fs.addFile("doc.pdf","5MB","/home/user/docs"); // Returns "doc.pdf".
fs.addFile("img.jpg","2MB","/home/user/photos"); // Returns "img.jpg".
fs.addFile("script.sh","10KB","/home/user/bin"); // Returns "script.sh".
fs.addFile("config.xml","1KB","/home/user/config"); // Returns "config.xml".
fs.addFile("backup.zip","100MB","/home/user/backup"); // Returns "backup.zip".
fs.applyFilter("/home/user","SIZE","1MB"); // Returns ["backup.zip","doc.pdf","img.jpg"] as all these files larger than 1 MB with the given directory path.
fs.applyFilter("/home","EXTENSION","xml"); // Returns ["config.xml]".
fs.applyFilter("/","PREFIX","doc"); // Returns ["doc.pdf"].
fs.addFile("doc.pdf","10MB","/home/user/docs"); // Returns "", since this is a duplicate filename in the given directory path.

## Python Template

```python
from typing import List, Optional

class FileSystem:
    def __init__(self, ):
        # TODO: Initialize FileSystem

    def addFile(self, fileName: str, size: str, path: str) -> str:
        # TODO: Implement addFile logic

    def applyFilter(self, path: str, filterType: str, filterParam: str) -> List[str]:
        # TODO: Implement applyFilter logic
```
