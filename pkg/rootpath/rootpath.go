package rootpath

import "os"

func Get() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
