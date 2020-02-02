package princess

import (
	"io"
	"io/ioutil"
	"os"
)

func Exists(path string) (bool, error) {
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	defer file.Close()
	return true, nil
}

func Save(src io.Reader, destpath string) error {
	dst, err := os.Create(destpath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

func All(dir string) ([]string, error) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	filenames := []string{}
	for _, fileInfo := range fileInfos {
		filenames = append(filenames, fileInfo.Name())
	}
	return filenames, nil
}

func Delete() {
	// delete princess
}
