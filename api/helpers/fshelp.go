package helpers

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/pootytang/danielleapi/types"
)

type File struct {
	SRC  string `json:"src"`
	ALT  string `json:"alt"`
	PATH string `json:"path,omitempty"`
}

func GetFiles(dir string, request_path string) ([]File, error) {
	slog.Info("GetFiles(): CALLED...")

	var exts = []string{"JPEG", "JPG", "PNG", "MP4"}
	var files []File
	endpoint := filepath.Base(request_path)
	folder_name := types.ENDPOINT_TO_FOLDER[endpoint]

	slog.Info(fmt.Sprintf("GetFiles(): Request URI: %s, Endpoint requested: %s, Matching folder: %s", request_path, endpoint, folder_name))
	if len(folder_name) <= 0 {
		slog.Error("GetFiles(): folder value is empty")
		return files, fmt.Errorf("GetFiles(): not folder mapping")
	}

	if !dirExists(dir) {
		slog.Error(fmt.Sprintf("GetFiles(): Directory %s does not exist", dir))
		return files, fmt.Errorf("GetFiles(): directory not found %s", dir)
	}

	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		slog.Debug(fmt.Sprintf("GetFiles->walkFunc(): Path set to %s and Entry set to %s", path, entry.Name()))

		if entry.IsDir() && entry.Name() == folder_name {
			slog.Debug(fmt.Sprintf("GetFiles->walkFunc(): Directory %s is valid", entry.Name()))
			return nil
		}

		if entry.IsDir() {
			slog.Info(fmt.Sprintf("GetFiles->walkFunc(): %s is a Directory - skipping...", entry.Name()))
			return filepath.SkipDir
		}

		slog.Debug("GetFiles->walkFunc(): Performing case-incensitive check on the files extention")
		hostURL, _ := types.GetHost()
		for _, ext := range exts {
			file_ext := strings.ToUpper(filepath.Ext(entry.Name()))
			no_dot_file_ext := strings.Split(file_ext, ".")[1]
			slog.Debug(fmt.Sprintf("GetFiles->walkFunc(): file_name = %s, file_ext = %s, ext_checked = %s", entry.Name(), no_dot_file_ext, ext))
			if no_dot_file_ext == ext {
				slog.Info(fmt.Sprintf("GetFiles->walkFunc(): file extention %s MATCH with ext %s. Adding image to array", entry.Name(), ext))
				files = append(files, File{
					SRC: hostURL + request_path + entry.Name(),
					ALT: entry.Name(),
				})
				break
			}
		}

		return nil
	})

	if err != nil {
		slog.Error("GetFiles(): error caught " + err.Error())
		return nil, err
	}

	return files, err
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func dirExists(path string) bool {
	stat, err := os.Stat(path)
	if err == nil && stat.IsDir() {
		return true
	}

	if errors.Is(err, fs.ErrNotExist) {
		slog.Error(fmt.Sprintf("dirExists(): %s does not exist", path))
		return false
	} else {
		slog.Error("dirExists(): problem searching for path")
		return false
	}
}
