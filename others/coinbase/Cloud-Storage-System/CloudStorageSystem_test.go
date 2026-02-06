package cloudstoragesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloudStorageExample1(t *testing.T) {
	storage := NewCloudStorage()

	assert.True(t, storage.AddFile("/dir1/dir2/file.txt", 10))
	assert.False(t, storage.CopyFile("/not-existing.file", "/dir1/file.txt"))
	assert.True(t, storage.CopyFile("/dir1/dir2/file.txt", "/dir1/file.txt"))
	assert.False(t, storage.AddFile("/dir1/file.txt", 15))
	assert.False(t, storage.CopyFile("/dir1/file.txt", "/dir1/dir2/file.txt"))
	assert.Equal(t, 10, storage.GetFileSize("/dir1/file.txt"))
	assert.Equal(t, -1, storage.GetFileSize("/not-existing.file"))
}

func TestCloudStorageExample2(t *testing.T) {
	storage := NewCloudStorage()

	assert.True(t, storage.AddFile("file1.txt", 5))
	assert.True(t, storage.AddFile("file2.txt", 3))
	assert.Equal(t, 5, storage.GetFileSize("file1.txt"))
	assert.True(t, storage.CopyFile("file1.txt", "file1_copy.txt"))
	assert.Equal(t, 5, storage.GetFileSize("file1_copy.txt"))
}

func TestCloudStorageExample3(t *testing.T) {
	storage := NewCloudStorage()

	assert.True(t, storage.AddFile("/projects/app/main.java", 1500))
	assert.True(t, storage.AddFile("/projects/app/utils.java", 800))
	assert.True(t, storage.AddFile("/projects/config/settings.xml", 200))
	assert.True(t, storage.CopyFile("/projects/app/main.java", "/backup/main.java"))
	assert.True(t, storage.CopyFile("/projects/config/settings.xml", "/backup/settings.xml"))
	assert.Equal(t, 1500, storage.GetFileSize("/backup/main.java"))
	assert.Equal(t, 200, storage.GetFileSize("/backup/settings.xml"))
	assert.False(t, storage.CopyFile("/projects/app/utils.java", "/backup/main.java"))
}

func TestCloudStorageExample4(t *testing.T) {
	storage := NewCloudStorage()

	assert.True(t, storage.AddFile("", 10))
	assert.False(t, storage.AddFile("", 20))
	assert.Equal(t, 10, storage.GetFileSize(""))
	assert.True(t, storage.AddFile("empty.txt", 0))
	assert.Equal(t, 0, storage.GetFileSize("empty.txt"))
	assert.True(t, storage.CopyFile("empty.txt", "empty_copy.txt"))
	assert.Equal(t, 0, storage.GetFileSize("empty_copy.txt"))
	assert.False(t, storage.CopyFile("nonexistent1.txt", "nonexistent2.txt"))
	assert.Equal(t, -1, storage.GetFileSize("does_not_exist.txt"))
	assert.True(t, storage.AddFile("file@#$.txt", 100))
	assert.True(t, storage.CopyFile("file@#$.txt", "copy@#$.txt"))
	assert.Equal(t, 100, storage.GetFileSize("copy@#$.txt"))
}
