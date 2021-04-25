package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// GetBytesForFileOrURL reads the file or filecontent from the http url
// returns the body
func GetBytesForFileOrURL(path string) ([]byte, error) {
	u, err := url.ParseRequestURI(path)
	if err == nil && u.Scheme != "" {
		return getBytesFromURL(path)
	}

	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	// Dir
	if stat.IsDir() {
		return nil, errors.New("please specify file path")
		// TODO: support decode multiple files
		if files, err := filePathWalkDir(path); err != nil {
			fmt.Println(files)
			return nil, err
		}
	}
	// File
	return ioutil.ReadFile(path)
}

// getBytesFromURL returns the body content of the given url
func getBytesFromURL(uri string) ([]byte, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// filePathWalkDir reads all the json files from the given dir
func filePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".json" {
			files = append(files, path)
		}
		// fmt.Println(filepath.Ext(path))
		return nil
	})
	return files, err
}
