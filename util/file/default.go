package file

import "os"

// spy 2018/6/28

/**
 * 创建文件
 */
func create(path string) error {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}

		defer file.Close()
	}
	return nil
}

/**
 * 删除文件
 */
func delete(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
