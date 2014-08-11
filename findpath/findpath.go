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
	TopDir   string `long:"top" description:"top directory to search for directory names"`
	Verbose  bool   `long:"verbose" short:"v"`
	FindArgs string `long:"find-args" description:"flags to pass to find"`
	Args     struct {
		Path string
	} `required:"yes" positional-args:"yes"`
}

var opt Options

func info(s string, args ...interface{}) {
	if opt.Verbose {
		fmt.Fprintf(os.Stderr, s, args...)
	}
}

func printerr(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s, args...)
}

func main() {
	_, err := flags.Parse(&opt)
	if err != nil {
		printerr("%s\n", err)
		os.Exit(1)
	}
	if opt.TopDir == "" {
		opt.TopDir = os.Getenv("HOME")
		if opt.TopDir == "" {
			printerr("must set --top flag or define $HOME\n")
			os.Exit(1)
		}
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
	pathComponents := tocomponents(opt.Args.Path)
	args := []string{opt.TopDir}
	args = append(args, strings.Fields(opt.FindArgs)...)
	args = append(args, "-name", pathComponents[0])
	out, err := exec.Command("find", args...).Output()
	if err != nil && len(out) == 0 {
		info(string(out))
		return "", fmt.Errorf("Error running find: %s", err)
	}
	parents := strings.Split(string(out), "\n")

	// Sort parents
	sort.Sort(PathSorter(parents))

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

func tocomponents(pathname string) []string {
	return strings.Split(pathname, string(filepath.Separator))
}

type PathSorter []string

func (s PathSorter) Len() int {
	return len(s)
}
func (s PathSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s PathSorter) Less(i, j int) bool {
	cmpi, cmpj := tocomponents(s[i]), tocomponents(s[j])

	hashiddeni := stringSliceContains(cmpi, func(s string) bool { return strings.HasPrefix(s, ".") })
	hashiddenj := stringSliceContains(cmpj, func(s string) bool { return strings.HasPrefix(s, ".") })
	if hashiddeni && !hashiddenj {
		return false
	} else if hashiddenj && !hashiddeni {
		return true
	}

	hasundi := stringSliceContains(cmpi, func(s string) bool { return strings.HasPrefix(s, "_") })
	hasundj := stringSliceContains(cmpj, func(s string) bool { return strings.HasPrefix(s, "_") })
	if hasundi && !hasundj {
		return false
	} else if hasundj && !hasundi {
		return true
	}

	if len(cmpi) < len(cmpj) {
		return true
	} else if len(cmpj) > len(cmpi) {
		return false
	} else {
		return len(cmpi) <= len(cmpj)
	}
}

func stringSliceContains(slice []string, fn func(string) bool) bool {
	for _, e := range slice {
		if fn(e) {
			return true
		}
	}
	return false
}
