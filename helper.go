package meta

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ModuleSrcPath(inputPath string) (goModDirPath string, err error) {
	goModDirPath = inputPath
	for i := 0; i < 10; i++ {
		goModFilePath := strings.Join([]string{goModDirPath, "go.mod"}, string(filepath.Separator))
		_, err = os.Stat(goModFilePath)
		if err == nil {
			return
		} else if os.IsNotExist(err) {
			goModDirPath = path.Dir(goModDirPath)
			continue
		} else {
			return
		}
	}
	return
}

func HasGoFile(root string) (has bool, err error) {
	names, err := filepath.Glob(strings.Join([]string{root, "*.go"}, string(filepath.Separator)))
	has = len(names) > 0
	return
}
