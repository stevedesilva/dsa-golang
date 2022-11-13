package recursion_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevedesilva/dsa-golang.git/commonsense/recursion"
)

func TestFindDirectories_AccumulatorReturned(t *testing.T) {
	// Given a Root directory and Sub directories
	rootName, expectedDirs := createRootAndSubdirectoriesWithTestFiles(t)
	defer os.Remove(rootName)
	// When
	res := recursion.FindDirectoriesWithAccumulator(rootName)
	// Then
	assert.Equal(t, expectedDirs, res)
}

func TestFindDirectories_AccumulatorPointer(t *testing.T) {
	// Given a Root directory and Sub directories
	rootName, expectedDirs := createRootAndSubdirectoriesWithTestFiles(t)
	defer os.Remove(rootName)
	// When
	res := recursion.FindDirectoriesWithSlicePointer(rootName)
	// Then
	assert.Equal(t, expectedDirs, res)
}

func createRootAndSubdirectoriesWithTestFiles(t *testing.T) (string, []string) {
	// root
	rootName, err := createDirectoryStructure("", "level_root_", "test_file_1", "test_file_2")
	assert.Nil(t, err)

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
	return rootName, expectedDirs
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
