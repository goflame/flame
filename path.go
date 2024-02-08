package flame

import "os"

func Root(path string) string {
	file, err := os.Getwd()
	if err != nil {
		return ""
	}

	if path[0] == '/' {
		return file + path
	}
	return file + "/" + path
}
