package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type fileTreeNode struct {
	dirs  map[string]*fileTreeNode
	files map[string]int
}

func size(node *fileTreeNode) int {
	result := 0
	for _, fileSize := range node.files {
		result += fileSize
	}
	for _, dir := range node.dirs {
		result += size(dir)
	}
	return result
}

func getPath(dirs ...string) []string {
	return dirs
}
func cd(workingDir []string, target string) []string {
	if target == "/" {
		return getPath()
	}
	if target == ".." {
		return workingDir[0 : len(workingDir)-1]
	}
	return append(workingDir, target)
}

func createDir(fileSystem *fileTreeNode, workingDir []string, name string) {
	currentNode := fileSystem
	for _, dir := range workingDir {
		currentNode = currentNode.dirs[dir]
	}
	if currentNode.dirs == nil {
		currentNode.dirs = map[string]*fileTreeNode{}
	}
	if currentNode.dirs[name] == nil {
		newNode := &fileTreeNode{}
		currentNode.dirs[name] = newNode
	}
}

func createFile(fileSystem *fileTreeNode, workingDir []string, name string, size int) {
	currentNode := fileSystem
	for _, dir := range workingDir {
		currentNode = currentNode.dirs[dir]
	}
	if currentNode.files == nil {
		currentNode.files = map[string]int{}
	}
	currentNode.files[name] = size
}

func buildFileTreeFromTerminalOutput(input *os.File) fileTreeNode {
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	fileTree := fileTreeNode{}
	workingDir := getPath()
	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		tokens := strings.Split(currentLine, " ")
		if tokens[0] == "$" {
			cmd := tokens[1]
			if cmd == "cd" {
				workingDir = cd(workingDir, tokens[2])
			}
		} else if tokens[0] == "dir" {
			createDir(&fileTree, workingDir, tokens[1])
		} else {
			size, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic(err)
			}
			createFile(&fileTree, workingDir, tokens[1], size)
		}
	}

	return fileTree
}

type DirectoryWithSize struct {
	path []string
	size int
}

func getAllDirectoriesWithSize(node *fileTreeNode, path []string) []DirectoryWithSize {
	result := make([]DirectoryWithSize, 0)
	result = append(result, DirectoryWithSize{
		path: path, size: size(node),
	})
	for subdir, subtree := range node.dirs {
		result = append(result, getAllDirectoriesWithSize(subtree, append(path, subdir))...)
	}
	return result
}

func SumSizeOfAllDirsWithTotalSizeAtMost100000(input *os.File) (string, error) {
	fileTree := buildFileTreeFromTerminalOutput(input)
	sum := 0
	for _, dirWithSize := range getAllDirectoriesWithSize(&fileTree, getPath()) {
		if dirWithSize.size <= 100000 {
			sum += dirWithSize.size
		}
	}
	return strconv.Itoa(sum), nil
}

func FreeUpSpaceForUpdate(input *os.File) (string, error) {
	fileTree := buildFileTreeFromTerminalOutput(input)
	totalSize := size(&fileTree)
	println("TOTAL SIZE IS " + strconv.Itoa(totalSize))
	unusedSpace := 70000000 - totalSize
	println("UNUSED SPACE IS " + strconv.Itoa(unusedSpace))
	minDirSizeToRemove := totalSize
	for _, dirWithSize := range getAllDirectoriesWithSize(&fileTree, getPath()) {
		if unusedSpace+dirWithSize.size > 30000000 {
			if dirWithSize.size < minDirSizeToRemove {
				minDirSizeToRemove = dirWithSize.size
			}
		}
	}
	return strconv.Itoa(minDirSizeToRemove), nil
}
