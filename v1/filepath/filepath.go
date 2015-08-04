package filepath

import (
	"os"
	"path/filepath"

	"../config"
)

func GetPath(file string) string {
	if filepath.IsAbs(file) {
		return file
	} else {
		path, err := filepath.Abs(GetWorkspace() + "/" + file)
		if err != nil {
			panic(err)
		}
		return path
	}
}

func GetWorkspace() string {
	var path string
	var err error

	if config.Std.PassPad.Workspace != "" {
		path = config.Std.PassPad.Workspace
	} else {
		path = filepath.Dir(os.Args[0])
	}

	path, err = filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	return path
}
