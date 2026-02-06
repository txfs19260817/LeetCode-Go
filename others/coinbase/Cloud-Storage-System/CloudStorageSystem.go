package cloudstoragesystem

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type fileEntry struct {
	size  int64
	owner string
}

type userInfo struct {
	capacity  int64
	used      int64
	files     map[string]struct{}
	unlimited bool
}

type CloudStorage struct {
	files map[string]fileEntry
	users map[string]*userInfo
}

func NewCloudStorage() *CloudStorage {
	storage := &CloudStorage{
		files: make(map[string]fileEntry),
		users: make(map[string]*userInfo),
	}
	storage.users["admin"] = &userInfo{
		capacity:  math.MaxInt64,
		files:     make(map[string]struct{}),
		unlimited: true,
	}
	return storage
}

func (cs *CloudStorage) AddUser(userId string, capacity int) bool {
	if userId == "admin" {
		return false
	}
	if _, exists := cs.users[userId]; exists {
		return false
	}
	cs.users[userId] = &userInfo{
		capacity: int64(capacity),
		files:    make(map[string]struct{}),
	}
	return true
}

func (cs *CloudStorage) AddFile(name string, size int) bool {
	return cs.AddFileBy("admin", name, size) != -1
}

func (cs *CloudStorage) AddFileBy(userId string, name string, size int) int {
	if _, exists := cs.files[name]; exists {
		return -1
	}
	user, exists := cs.users[userId]
	if !exists {
		return -1
	}
	fileSize := int64(size)
	if !cs.hasCapacity(user, fileSize) {
		return -1
	}
	cs.files[name] = fileEntry{size: fileSize, owner: userId}
	user.used += fileSize
	user.files[name] = struct{}{}
	return cs.remainingCapacity(user)
}

func (cs *CloudStorage) CopyFile(nameFrom string, nameTo string) bool {
	if _, exists := cs.files[nameTo]; exists {
		return false
	}
	entry, exists := cs.files[nameFrom]
	if !exists {
		return false
	}
	user := cs.users[entry.owner]
	if user == nil {
		return false
	}
	if !cs.hasCapacity(user, entry.size) {
		return false
	}
	cs.files[nameTo] = fileEntry{size: entry.size, owner: entry.owner}
	user.used += entry.size
	user.files[nameTo] = struct{}{}
	return true
}

func (cs *CloudStorage) GetFileSize(name string) int {
	if entry, exists := cs.files[name]; exists {
		return int(entry.size)
	}
	return -1
}

func (cs *CloudStorage) FindFile(prefix string, suffix string) []string {
	type match struct {
		name string
		size int64
	}
	matches := make([]match, 0)
	for name, entry := range cs.files {
		if strings.HasPrefix(name, prefix) && strings.HasSuffix(name, suffix) {
			matches = append(matches, match{name: name, size: entry.size})
		}
	}
	sort.Slice(matches, func(i, j int) bool {
		if matches[i].size != matches[j].size {
			return matches[i].size > matches[j].size
		}
		return matches[i].name < matches[j].name
	})
	result := make([]string, len(matches))
	for i, item := range matches {
		result[i] = fmt.Sprintf("%s(%d)", item.name, item.size)
	}
	return result
}

func (cs *CloudStorage) UpdateCapacity(userId string, capacity int) int {
	user, exists := cs.users[userId]
	if !exists {
		return -1
	}
	if user.unlimited {
		return 0
	}
	user.capacity = int64(capacity)
	if user.used <= user.capacity {
		return 0
	}
	type fileInfo struct {
		name string
		size int64
	}
	files := make([]fileInfo, 0, len(user.files))
	for name := range user.files {
		entry, exists := cs.files[name]
		if exists {
			files = append(files, fileInfo{name: name, size: entry.size})
		}
	}
	sort.Slice(files, func(i, j int) bool {
		if files[i].size != files[j].size {
			return files[i].size > files[j].size
		}
		return files[i].name < files[j].name
	})
	removed := 0
	for _, file := range files {
		if user.used <= user.capacity {
			break
		}
		delete(cs.files, file.name)
		delete(user.files, file.name)
		user.used -= file.size
		removed++
	}
	return removed
}

func (cs *CloudStorage) CompressFile(userId string, name string) int {
	user, exists := cs.users[userId]
	if !exists {
		return -1
	}
	entry, exists := cs.files[name]
	if !exists {
		return -1
	}
	if entry.owner != userId {
		return -1
	}
	compressedName := name + ".COMPRESSED"
	if _, exists := cs.files[compressedName]; exists {
		return -1
	}
	newSize := entry.size / 2
	delete(cs.files, name)
	delete(user.files, name)
	user.used -= entry.size

	cs.files[compressedName] = fileEntry{size: newSize, owner: userId}
	user.files[compressedName] = struct{}{}
	user.used += newSize
	return cs.remainingCapacity(user)
}

func (cs *CloudStorage) DecompressFile(userId string, name string) int {
	user, exists := cs.users[userId]
	if !exists {
		return -1
	}
	entry, exists := cs.files[name]
	if !exists {
		return -1
	}
	if entry.owner != userId {
		return -1
	}
	if !strings.HasSuffix(name, ".COMPRESSED") {
		return -1
	}
	originalName := strings.TrimSuffix(name, ".COMPRESSED")
	if _, exists := cs.files[originalName]; exists {
		return -1
	}
	newSize := entry.size * 2
	if !cs.hasCapacity(user, newSize-entry.size) {
		return -1
	}
	delete(cs.files, name)
	delete(user.files, name)
	user.used -= entry.size

	cs.files[originalName] = fileEntry{size: newSize, owner: userId}
	user.files[originalName] = struct{}{}
	user.used += newSize
	return cs.remainingCapacity(user)
}

func (cs *CloudStorage) hasCapacity(user *userInfo, delta int64) bool {
	if user.unlimited {
		return true
	}
	return user.used+delta <= user.capacity
}

func (cs *CloudStorage) remainingCapacity(user *userInfo) int {
	if user.unlimited {
		return int(user.capacity - user.used)
	}
	return int(user.capacity - user.used)
}
