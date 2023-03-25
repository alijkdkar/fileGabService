package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func statTimes(name string) (atime, mtime, ctime time.Time, err error) {
	fmt.Println("statTime:", name)
	fi, err := os.Stat(name)
	if err != nil {
		return
	}
	mtime = fi.ModTime()
	stat := fi.Sys().(*syscall.Stat_t)
	atime = time.Unix(int64(stat.Atimespec.Sec), int64(stat.Atimespec.Nsec))
	ctime = time.Unix(int64(stat.Ctimespec.Sec), int64(stat.Ctimespec.Nsec))
	return
}

func main() {

	fmt.Println("Please enter Full path of Folder")
	var path string
	var timeToDel int

	fmt.Scanln(&path)

	fmt.Println("Please enter Time for delete file(Second):")
	fmt.Scanln(&timeToDel)

	for {
		files, err := os.ReadDir(path)
		if err != nil {
			println("erorr " + err.Error())
		}

		for _, v := range files {
			filePath := path + v.Name()
			_, _, cTime, err := statTimes(filePath)

			if err != nil {
				fmt.Println("Error: ", err.Error())
				panic(err)
			}
			
			nowl := time.Now().After(cTime.Add(time.Second * time.Duration(timeToDel)))
			if nowl {
				println("deleted ")
				os.Remove(filePath)
			}
		}
	}

}
