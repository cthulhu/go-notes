package scanner

import (
	"context"
	"os"
	"path/filepath"
	"strings"
)

var skipDirs []string

func init() {
	skipDirs = []string{"vendor", ".git"}
}

// Scanner constructor
// It's execution can be terminated via context
// returns scanned paths channel and errors channel
func New(ctx context.Context, pathsToScan []string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errors := make(chan error, 1)
	go func() {
		defer func() {
			close(paths)
			close(errors)
		}()
		var err error
		pathsToScan, err = scanPaths(pathsToScan)
		if err != nil {
			errors <- err
			return
		}
		for _, path := range pathsToScan {
			select {
			case <-ctx.Done():
				return
			case paths <- path:
			}
		}
	}()
	return paths, errors
}

func scanPaths(pathsToScan []string) ([]string, error) {
	var newPathsToScan []string
	for _, pathToCheck := range pathsToScan {
		filepath.Walk(pathToCheck, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && contains(skipDirs, path) {
				return filepath.SkipDir
			}
			if !info.IsDir() && strings.HasSuffix(path, ".go") {
				newPathsToScan = append(newPathsToScan, path)
			}
			return nil
		})
	}
	return newPathsToScan, nil
}

func contains(paths []string, path string) bool {
	for _, p := range paths {
		if strings.Contains(path, p) {
			return true
		}
	}
	return false
}

func isDir(filename string) bool {
	fi, err := os.Stat(filename)
	return err == nil && fi.IsDir()
}
