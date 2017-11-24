package util

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

// CurrentDir returns the current working directory
func CurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

// HomeDir returns the home directory
func HomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return home
}

// Exists checks whether a file is exists or not
func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
    return true
}

// FilePath returns the full path (with extension) to file
func FilePath(dir, fileName, ext string) string {
	return filepath.Join(dir, fileName + "." + ext)
}