package recursion_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevedesilva/dsa-golang.git/commonsense/recursion"
)

var (
	tempDir string
)

func TestFindDirectories(t *testing.T) {

	expectedDirs := make([]string, 0, 5)
	rootName, err := os.MkdirTemp("", "level_root_")
	if err != nil {
		return
	}
	defer os.Remove(rootName)

	createEmptyFile("test_file_1", rootName)
	createEmptyFile("test_file_2", rootName)

	levelOneName, err := os.MkdirTemp(rootName, "level_1_")
	expectedDirs = append(expectedDirs, levelOneName)
	createEmptyFile("test_file_1a", levelOneName)
	createEmptyFile("test_file_1b", levelOneName)

	levelTwoName, err := os.MkdirTemp(levelOneName, "level_2_")
	expectedDirs = append(expectedDirs, levelTwoName)
	createEmptyFile("test_file_2a", levelTwoName)

	levelThreeName, err := os.MkdirTemp(levelTwoName, "level_3_")
	expectedDirs = append(expectedDirs, levelThreeName)
	createEmptyFile("test_file_3a", levelThreeName)

	levelFourName, err := os.MkdirTemp(levelThreeName, "level_4_")
	expectedDirs = append(expectedDirs, levelFourName)
	levelFiveName, err := os.MkdirTemp(levelThreeName, "level_5_")
	expectedDirs = append(expectedDirs, levelFiveName)
	createEmptyFile("test_file_4a", levelFiveName)

	res := recursion.FindDirectories(rootName)
	assert.Equal(t, expectedDirs, res)
}

func createEmptyFile(name string, tempDir string) {
	filename := filepath.Join(tempDir, name)
	data := []byte("")
	os.WriteFile(filename, data, 0644)
}
