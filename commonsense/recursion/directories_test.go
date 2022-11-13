package recursion_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevedesilva/dsa-golang.git/commonsense/recursion"
)

func TestFindDirectories(t *testing.T) {

	// Given a Root directory and files
	rootName, err := createDirectoryStructure("", "level_root_", "test_file_1", "test_file_2")
	assert.Nil(t, err)
	defer os.Remove(rootName)

	// And Sub directories
	expectedDirs := make([]string, 0, 5)
	levelOneName, err := createDirectoryStructure(rootName, "level_1_", "test_file_1a", "test_file_1b")
	assert.Nil(t, err)
	expectedDirs = append(expectedDirs, levelOneName)

	levelTwoName, err := createDirectoryStructure(levelOneName, "level_2_", "test_file_2a")
	assert.Nil(t, err)
	expectedDirs = append(expectedDirs, levelTwoName)

	levelThreeName, err := createDirectoryStructure(levelTwoName, "level_3_", "test_file_3a")
	assert.Nil(t, err)
	expectedDirs = append(expectedDirs, levelThreeName)

	levelFourName, err := createDirectoryStructure(levelThreeName, "level_4_")
	assert.Nil(t, err)
	expectedDirs = append(expectedDirs, levelFourName)

	levelFiveName, err := createDirectoryStructure(levelFourName, "level_5_", "test_file_5a")
	assert.Nil(t, err)
	expectedDirs = append(expectedDirs, levelFiveName)

	// When
	res := recursion.FindDirectoriesWithAccumulator(rootName)

	// Then
	assert.Equal(t, expectedDirs, res)
}

func createEmptyFile(name string, tempDir string) {
	filename := filepath.Join(tempDir, name)
	data := []byte("")
	os.WriteFile(filename, data, 0644)
}

func createDirectoryStructure(dirName, dirPattern string, files ...string) (string, error) {
	dirPath, err := os.MkdirTemp(dirName, dirPattern)
	if err != nil {
		return dirPath, err
	}
	for _, file := range files {
		createEmptyFile(file, dirPath)
	}
	return dirPath, nil
}
