/*
Simple fuzzy filename search.
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sqs/go-flags"
)

type Options struct {
	TopDir  string `long:"top" description:"top directory to search for directory names" default:"/home/beyang"`
	Verbose bool   `long:"verbose" short:"v"`
	Args    struct {
		Path string
	} `required:"yes" positional-args:"yes"`
}

var opt Options

func info(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s, args...)
}

func main() {
	_, err := flags.Parse(&opt)
	if err != nil {
		os.Exit(1)
	}

	// If path exists at $PWD, just return it
	if exists(opt.Args.Path) {
		fmt.Println(opt.Args.Path)
		return
	}

	// Otherwise, search for a path matching it from top-level dir
	found, err := find()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(found)
}

func find() (string, error) {
	// Get parent directories via find
	pathComponents := strings.Split(opt.Args.Path, string(filepath.Separator))
	out, err := exec.Command("find", opt.TopDir, "-name", pathComponents[0]).CombinedOutput()
	if err != nil {
		info(string(out))
		return "", err
	}
	parents := strings.Split(string(out), "\n")

	// Sort parents
	sort.Strings(parents)

	// Return first parent with matching child
	for _, parent := range parents {
		path := filepath.Join(parent, filepath.Join(pathComponents[1:]...))
		if exists(path) {
			return path, nil
		}
	}
	return "", fmt.Errorf("Didn't find matching directory")
}

func exists(path string) bool {
	_, err := os.Lstat(path)
	return err == nil
}
