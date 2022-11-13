package recursion

import (
	"fmt"
	"os"
)

func FindDirectoriesWithAccumulator(rootDirPath string) []string {
	foundDirs := make([]string, 0)
	return findDirectories(rootDirPath, foundDirs)
}

func findDirectories(directory string, acc []string) []string {
	directories, err := os.ReadDir(directory)
	if err != nil {
		return acc
	}

	for _, dir := range directories {
		if dir.IsDir() {
			name := fmt.Sprintf("%s/%s", directory, dir.Name())
			// add current dir to found list
			acc = append(acc, name)
			// search sub dirs
			acc = findDirectories(name, acc)
		}
	}
	return acc
}

func FindDirectoriesWithSlicePointer(rootDirPath string) []string {
	foundDirs := make([]string, 0)
	findDirectoriesPointer(rootDirPath, &foundDirs)
	return foundDirs
}

func findDirectoriesPointer(directory string, acc *[]string) {
	directories, err := os.ReadDir(directory)
	if err != nil {
		return
	}

	for _, dir := range directories {
		if dir.IsDir() {
			name := fmt.Sprintf("%s/%s", directory, dir.Name())
			*acc = append(*acc, name)
			findDirectoriesPointer(name, acc)
		}
	}
}
