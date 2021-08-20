package filesystem

import (
	"log"
	"os"
	"path/filepath"

	"github.com/h2non/filetype"
)

type FileInfo struct {
	FileName      string
	FileType      string
	FileMIME      string
	FileExtension string
}

func GetFileType(filename string) (FileInfo, error) {

	r, _ := filetype.MatchFile(filename)

	var f FileInfo
	f.FileMIME = r.MIME.Value
	f.FileType = r.MIME.Type
	f.FileExtension = r.Extension
	f.FileName = filename

	return f, nil

}

func GetFiles(directory string) ([]FileInfo, error) {
	var files []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	var foundFiles []FileInfo

	for _, f := range files {
		fileInfo, _ := GetFileType(f)
		foundFiles = append(foundFiles, fileInfo)
	}

	return foundFiles, err
}
