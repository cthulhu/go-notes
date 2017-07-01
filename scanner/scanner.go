package scanner

import (
	"context"
	"os"
	"path/filepath"
	"strings"
)

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
			if (!info.IsDir() && strings.HasSuffix(path, ".go")) && !contains(pathsToScan, path) {
				newPathsToScan = append(newPathsToScan, path)
			}
			return nil
		})
	}
	return newPathsToScan, nil
}

func contains(paths []string, path string) bool {
	for _, p := range paths {
		if p == path {
			return true
		}
	}
	return false
}

func isDir(filename string) bool {
	fi, err := os.Stat(filename)
	return err == nil && fi.IsDir()
}
