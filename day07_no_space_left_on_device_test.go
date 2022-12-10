package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_size(t *testing.T) {
	fileTree := fileTreeNode{
		files: map[string]int{
			"a.txt": 20,
			"b.txt": 30,
		},
		dirs: map[string]*fileTreeNode{
			"x": {
				files: map[string]int{
					"c.txt": 10,
				},
				dirs: map[string]*fileTreeNode{
					"y": {},
					"z": {
						files: map[string]int{
							"d.txt": 7,
						},
					},
				},
			},
		},
	}

	assert.Equal(t, 67, size(&fileTree))
}

func Test_cd(t *testing.T) {
	assert.Equal(t, getPath("a"), cd(getPath(), "a"))
	assert.Equal(t, getPath("a", "b"), cd(getPath("a"), "b"))
	assert.Equal(t, getPath("a"), cd(getPath("a", "b"), ".."))
	assert.Equal(t, getPath(), cd(getPath("a", "b"), "/"))
}

func Test_createDir(t *testing.T) {
	root := fileTreeNode{}

	workingDir := getPath()
	createDir(&root, workingDir, "a")
	assert.Equal(t,
		fileTreeNode{dirs: map[string]*fileTreeNode{
			"a": {},
		}}, root)

	workingDir = cd(workingDir, "a")
	createDir(&root, workingDir, "b")
	assert.Equal(t,
		fileTreeNode{dirs: map[string]*fileTreeNode{
			"a": {
				dirs: map[string]*fileTreeNode{
					"b": {},
				},
			},
		}}, root)

	createDir(&root, getPath(), "a") // should not overwrite existing dir
	assert.Equal(t,
		fileTreeNode{dirs: map[string]*fileTreeNode{
			"a": {
				dirs: map[string]*fileTreeNode{
					"b": {},
				},
			},
		}}, root)
}

func Test_createFile(t *testing.T) {
	root := fileTreeNode{}
	workingDir := getPath()
	createFile(&root, workingDir, "x.txt", 1000000)
	assert.Equal(t,
		fileTreeNode{
			files: map[string]int{
				"x.txt": 1000000,
			},
		}, root)

	createDir(&root, workingDir, "a")
	workingDir = cd(workingDir, "a")
	createFile(&root, workingDir, "y.txt", 999)
	assert.Equal(t,
		fileTreeNode{
			dirs: map[string]*fileTreeNode{
				"a": {
					files: map[string]int{
						"y.txt": 999,
					},
				},
			},
			files: map[string]int{
				"x.txt": 1000000,
			}}, root)
}

func Test_buildFileTreeFromTerminalOutput(t *testing.T) {
	input, err := os.Open("inputs/day07/default.txt")
	if err != nil {
		t.Fatal(err)
	}
	fileTree := buildFileTreeFromTerminalOutput(input)
	assert.Equal(t, fileTreeNode{
		dirs: map[string]*fileTreeNode{
			"a": {
				dirs: map[string]*fileTreeNode{
					"e": {
						files: map[string]int{
							"i": 584,
						},
					},
				},
				files: map[string]int{
					"f":     29116,
					"g":     2557,
					"h.lst": 62596,
				},
			},
			"d": {
				files: map[string]int{
					"j":     4060174,
					"d.log": 8033020,
					"d.ext": 5626152,
					"k":     7214296,
				},
			},
		},
		files: map[string]int{
			"b.txt": 14848514,
			"c.dat": 8504156,
		},
	}, fileTree)
}

func Test_getAllDirectoriesWithSize(t *testing.T) {
	fileTree := fileTreeNode{
		dirs: map[string]*fileTreeNode{
			"a": {
				dirs: map[string]*fileTreeNode{
					"e": {
						files: map[string]int{
							"i": 584,
						},
					},
				},
				files: map[string]int{
					"f":     29116,
					"g":     2557,
					"h.lst": 62596,
				},
			},
			"d": {
				files: map[string]int{
					"j":     4060174,
					"d.log": 8033020,
					"d.ext": 5626152,
					"k":     7214296,
				},
			},
		},
		files: map[string]int{
			"b.txt": 14848514,
			"c.dat": 8504156,
		},
	}
	result := getAllDirectoriesWithSize(&fileTree, getPath())
	assert.ElementsMatch(t, []DirectoryWithSize{
		{path: getPath(), size: 48381165},
		{path: getPath("a", "e"), size: 584},
		{path: getPath("a"), size: 94853},
		{path: getPath("d"), size: 24933642},
	}, result)
}
