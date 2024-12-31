package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var fileRegexp = regexp.MustCompile(` \w{32}\b`)

func processMarkdownFiles() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			substitutedContent := regexp.MustCompile(`\[.*?]\((.*?)\)`).ReplaceAllFunc(content, func(match []byte) []byte {
				matchStr := string(match)
				matches := regexp.MustCompile(`\[.*?]\((.*?)\)`).FindStringSubmatch(matchStr)
				if len(matches) > 1 {
					fileName := filepath.Base(matches[1])
					fileNameWithoutExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))
					return []byte(fmt.Sprintf("[%s](%s)", fileNameWithoutExt, fileNameWithoutExt))
				}
				return match
			})

			err = os.WriteFile(path, substitutedContent, 0644)
			if err != nil {
				return err
			}

			newPath := fileRegexp.ReplaceAllString(path, "")
			if newPath != path {
				err = os.Rename(path, newPath)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error processing markdown files:", err)
	}
}

func extractAndReplaceSrc() {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	rawFolder := filepath.Join(rootDir, "raw")
	files, err := os.ReadDir(rawFolder)
	if err != nil {
		fmt.Println("Error reading raw folder:", err)
		return
	}

	var zipFilePath string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".zip") {
			zipFilePath = filepath.Join(rawFolder, file.Name())
			break
		}
	}

	if zipFilePath == "" {
		fmt.Println("No zip files found in the 'raw' folder.")
		return
	}

	extractPath := filepath.Join(rootDir, "temp_extract")
	err = unzip(zipFilePath, extractPath)
	if err != nil {
		fmt.Println("Error extracting zip file:", err)
		return
	}

	srcPath := filepath.Join(rootDir, "src")
	if _, err := os.Stat(srcPath); !os.IsNotExist(err) {
		err = os.RemoveAll(srcPath)
		if err != nil {
			fmt.Println("Error removing existing src folder:", err)
			return
		}
	}

	extractedFolders, err := os.ReadDir(extractPath)
	if err != nil {
		fmt.Println("Error reading extracted folder:", err)
		return
	}

	var extractedFolderPath string
	for _, folder := range extractedFolders {
		if folder.IsDir() {
			extractedFolderPath = filepath.Join(extractPath, folder.Name())
			break
		}
	}

	if extractedFolderPath == "" {
		fmt.Println("No folders found in the extracted zip file.")
		return
	}

	err = os.Rename(extractedFolderPath, srcPath)
	if err != nil {
		fmt.Println("Error moving extracted folder to src:", err)
		return
	}

	err = os.RemoveAll(extractPath)
	if err != nil {
		fmt.Println("Error removing temporary extraction folder:", err)
	}
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func(r *zip.ReadCloser) {
		_ = r.Close()
	}(r)

	for _, f := range r.File {
		fPath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(fPath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			_ = outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)

		_ = outFile.Close()
		_ = rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	extractAndReplaceSrc()
	processMarkdownFiles()
}
