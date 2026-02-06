from typing import Dict, List, Set, Tuple
import sys


class _UserInfo:
    def __init__(self, capacity: int, unlimited: bool = False) -> None:
        self.capacity = capacity
        self.used = 0
        self.files: Set[str] = set()
        self.unlimited = unlimited


class CloudStorage:
    def __init__(self) -> None:
        self._files: Dict[str, Tuple[int, str]] = {}
        self._users: Dict[str, _UserInfo] = {
            "admin": _UserInfo(sys.maxsize, unlimited=True)
        }

    def addUser(self, userId: str, capacity: int) -> bool:
        if userId == "admin" or userId in self._users:
            return False
        self._users[userId] = _UserInfo(capacity)
        return True

    def addFile(self, name: str, size: int) -> bool:
        return self.addFileBy("admin", name, size) != -1

    def addFileBy(self, userId: str, name: str, size: int) -> int:
        if name in self._files:
            return -1
        user = self._users.get(userId)
        if user is None:
            return -1
        if not self._has_capacity(user, size):
            return -1
        self._files[name] = (size, userId)
        user.used += size
        user.files.add(name)
        return self._remaining_capacity(user)

    def copyFile(self, nameFrom: str, nameTo: str) -> bool:
        if nameTo in self._files:
            return False
        entry = self._files.get(nameFrom)
        if entry is None:
            return False
        size, owner = entry
        user = self._users.get(owner)
        if user is None or not self._has_capacity(user, size):
            return False
        self._files[nameTo] = (size, owner)
        user.used += size
        user.files.add(nameTo)
        return True

    def getFileSize(self, name: str) -> int:
        entry = self._files.get(name)
        return entry[0] if entry is not None else -1

    def findFile(self, prefix: str, suffix: str) -> List[str]:
        matches = [
            (name, entry[0])
            for name, entry in self._files.items()
            if name.startswith(prefix) and name.endswith(suffix)
        ]
        matches.sort(key=lambda item: (-item[1], item[0]))
        return [f"{name}({size})" for name, size in matches]

    def updateCapacity(self, userId: str, capacity: int) -> int:
        user = self._users.get(userId)
        if user is None:
            return -1
        if user.unlimited:
            return 0
        user.capacity = capacity
        if user.used <= user.capacity:
            return 0
        files = [
            (name, self._files[name][0]) for name in user.files if name in self._files
        ]
        files.sort(key=lambda item: (-item[1], item[0]))
        removed = 0
        for name, size in files:
            if user.used <= user.capacity:
                break
            del self._files[name]
            user.files.remove(name)
            user.used -= size
            removed += 1
        return removed

    def compressFile(self, userId: str, name: str) -> int:
        user = self._users.get(userId)
        if user is None:
            return -1
        entry = self._files.get(name)
        if entry is None:
            return -1
        size, owner = entry
        if owner != userId:
            return -1
        compressed_name = f"{name}.COMPRESSED"
        if compressed_name in self._files:
            return -1
        new_size = size // 2
        del self._files[name]
        user.files.remove(name)
        user.used -= size
        self._files[compressed_name] = (new_size, userId)
        user.files.add(compressed_name)
        user.used += new_size
        return self._remaining_capacity(user)

    def decompressFile(self, userId: str, name: str) -> int:
        user = self._users.get(userId)
        if user is None:
            return -1
        entry = self._files.get(name)
        if entry is None:
            return -1
        size, owner = entry
        if owner != userId:
            return -1
        if not name.endswith(".COMPRESSED"):
            return -1
        original_name = name[: -len(".COMPRESSED")]
        if original_name in self._files:
            return -1
        new_size = size * 2
        if not self._has_capacity(user, new_size - size):
            return -1
        del self._files[name]
        user.files.remove(name)
        user.used -= size
        self._files[original_name] = (new_size, userId)
        user.files.add(original_name)
        user.used += new_size
        return self._remaining_capacity(user)

    def _has_capacity(self, user: _UserInfo, delta: int) -> bool:
        return user.unlimited or user.used + delta <= user.capacity

    def _remaining_capacity(self, user: _UserInfo) -> int:
        return user.capacity - user.used


if __name__ == "__main__":
    storage = CloudStorage()
    assert storage.addFile("/dir1/dir2/file.txt", 10) is True
    assert storage.copyFile("/not-existing.file", "/dir1/file.txt") is False
    assert storage.copyFile("/dir1/dir2/file.txt", "/dir1/file.txt") is True
    assert storage.addFile("/dir1/file.txt", 15) is False
    assert storage.copyFile("/dir1/file.txt", "/dir1/dir2/file.txt") is False
    assert storage.getFileSize("/dir1/file.txt") == 10
    assert storage.getFileSize("/not-existing.file") == -1

    storage = CloudStorage()
    assert storage.addFile("file1.txt", 5) is True
    assert storage.addFile("file2.txt", 3) is True
    assert storage.getFileSize("file1.txt") == 5
    assert storage.copyFile("file1.txt", "file1_copy.txt") is True
    assert storage.getFileSize("file1_copy.txt") == 5

    storage = CloudStorage()
    assert storage.addFile("/projects/app/main.java", 1500) is True
    assert storage.addFile("/projects/app/utils.java", 800) is True
    assert storage.addFile("/projects/config/settings.xml", 200) is True
    assert storage.copyFile("/projects/app/main.java", "/backup/main.java") is True
    assert (
        storage.copyFile("/projects/config/settings.xml", "/backup/settings.xml")
        is True
    )
    assert storage.getFileSize("/backup/main.java") == 1500
    assert storage.getFileSize("/backup/settings.xml") == 200
    assert storage.copyFile("/projects/app/utils.java", "/backup/main.java") is False

    storage = CloudStorage()
    assert storage.addFile("", 10) is True
    assert storage.addFile("", 20) is False
    assert storage.getFileSize("") == 10
    assert storage.addFile("empty.txt", 0) is True
    assert storage.getFileSize("empty.txt") == 0
    assert storage.copyFile("empty.txt", "empty_copy.txt") is True
    assert storage.getFileSize("empty_copy.txt") == 0
    assert storage.copyFile("nonexistent1.txt", "nonexistent2.txt") is False
    assert storage.getFileSize("does_not_exist.txt") == -1
    assert storage.addFile("file@#$.txt", 100) is True
    assert storage.copyFile("file@#$.txt", "copy@#$.txt") is True
    assert storage.getFileSize("copy@#$.txt") == 100

    storage = CloudStorage()
    assert storage.addFile("/root/dir/another_dir/file.mp3", 10) is True
    assert storage.addFile("/root/file.mp3", 5) is True
    assert storage.addFile("/root/music/file.mp3", 7) is True
    assert storage.copyFile("/root/music/file.mp3", "/root/dir/file.mp3") is True
    assert storage.findFile("/root", ".mp3") == [
        "/root/dir/another_dir/file.mp3(10)",
        "/root/dir/file.mp3(7)",
        "/root/music/file.mp3(7)",
        "/root/file.mp3(5)",
    ]
    assert storage.findFile("/root", "file.txt") == []
    assert storage.findFile("/dir", "file.mp3") == []

    storage = CloudStorage()
    assert storage.addFile("", 10) is True
    assert storage.addFile("file", 20) is True
    assert storage.addFile("file.ext", 30) is True
    assert storage.addFile(".hidden", 40) is True
    assert storage.findFile("", "") == [
        ".hidden(40)",
        "file.ext(30)",
        "file(20)",
        "(10)",
    ]
    assert storage.findFile("file", "") == ["file.ext(30)", "file(20)"]
    assert storage.findFile("", "file") == ["file(20)"]
    assert storage.addFile("test@file#.txt", 100) is True
    assert storage.addFile("test$file%.txt", 90) is True
    assert storage.findFile("test", ".txt") == [
        "test@file#.txt(100)",
        "test$file%.txt(90)",
    ]
    assert storage.findFile("nonexistent", ".pdf") == []
    assert storage.findFile("/root", "nonexistent") == []
    assert storage.addFile("abc", 60) is True
    assert storage.findFile("ab", "bc") == ["abc(60)"]
    assert storage.findFile("abc", "abc") == ["abc(60)"]

    storage = CloudStorage()
    assert storage.addFile("zebra.txt", 100) is True
    assert storage.addFile("apple.txt", 100) is True
    assert storage.addFile("banana.txt", 100) is True
    assert storage.addFile("cherry.txt", 200) is True
    assert storage.addFile("date.txt", 50) is True
    assert storage.findFile("", ".txt") == [
        "cherry.txt(200)",
        "apple.txt(100)",
        "banana.txt(100)",
        "zebra.txt(100)",
        "date.txt(50)",
    ]
    assert storage.addFile("app_zebra.log", 150) is True
    assert storage.addFile("app_apple.log", 150) is True
    assert storage.addFile("app_banana.log", 150) is True
    assert storage.findFile("app_", ".log") == [
        "app_apple.log(150)",
        "app_banana.log(150)",
        "app_zebra.log(150)",
    ]

    storage = CloudStorage()
    assert storage.addUser("user1", 125) is True
    assert storage.addUser("user1", 100) is False
    assert storage.addUser("user2", 100) is True
    assert storage.addFileBy("user1", "/dir/file.big", 50) == 75
    assert storage.addFileBy("user1", "/file.med", 30) == 45
    assert storage.addFileBy("user2", "/file.med", 40) == -1
    assert storage.copyFile("/file.med", "/dir/another/file.med") is True
    assert storage.copyFile("/file.med", "/dir/another/another/file.med") is False
    assert storage.addFileBy("user1", "/dir/file.small", 10) == 5
    assert storage.addFile("/dir/admin_file", 200) is True
    assert storage.addFileBy("user1", "/dir/file.small", 5) == -1
    assert storage.addFileBy("user1", "/my_folder/file.huge", 100) == -1
    assert storage.addFileBy("user3", "/my_folder/file.huge", 100) == -1
    assert storage.updateCapacity("user1", 300) == 0
    assert storage.updateCapacity("user1", 50) == 2
    assert storage.updateCapacity("user2", 1000) == 0

    storage = CloudStorage()
    assert storage.addUser("user1", 1000) is True
    assert storage.addUser("user2", 500) is True
    assert storage.addFileBy("user1", "/dir/file.mp4", 500) == 500
    assert storage.compressFile("user2", "/dir/file.mp4") == -1
    assert storage.compressFile("user3", "/dir/file.mp4") == -1
    assert storage.compressFile("user1", "/folder/non_existing_file") == -1
    assert storage.compressFile("user1", "/dir/file.mp4") == 750
    assert storage.getFileSize("/dir/file.mp4.COMPRESSED") == 250
    assert storage.getFileSize("/dir/file.mp4") == -1
    assert storage.copyFile("/dir/file.mp4.COMPRESSED", "/file.mp4.COMPRESSED") is True
    assert storage.addFileBy("user1", "/dir/file.mp4", 500) == 0
    assert storage.decompressFile("user1", "/dir/file.mp4.COMPRESSED") == -1
    assert storage.updateCapacity("user1", 2000) == 0
    assert storage.decompressFile("user2", "/dir/file.mp4.COMPRESSED") == -1
    assert storage.decompressFile("user3", "/dir/file.mp4.COMPRESSED") == -1
    assert storage.decompressFile("user1", "/dir/file.mp4.COMPRESSED") == -1
    assert storage.decompressFile("user1", "/file.mp4.COMPRESSED") == 750
