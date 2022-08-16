package filesystem

import "os"

func MakeDirIfNotExists(dirPath string) error {
	if !existsDir(dirPath) {
		if err := os.Mkdir(dirPath, 0666); err != nil {
			return err
		}
	}
	return nil
}

func existsDir(dirPath string) bool {
	if f, err := os.Stat(dirPath); os.IsNotExist(err) || !f.IsDir() {
		return false
	}
	return true
}
