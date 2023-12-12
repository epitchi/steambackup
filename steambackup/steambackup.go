package steambackup

import (
	"fmt"
	"time"
)

func CopyFromBackupToTemp(source, destination string) {
	err := UnzipFolder(source, destination)
	if err != nil {
		fmt.Printf("error backup folder : %s\n",err.Error())
	}
}

func CopyFromTempToBackup(source, destination string) {
	err := ZipFolder(destination, source)
	if err != nil {
		fmt.Printf("error backup folder : %s\n",err.Error())
	}
}

var (
	stop = false
)

func StartBackup(source, backup string) {
	CopyFromBackupToTemp(backup, source)

	go func() {
		for {
			if stop {
				break
			}
			CopyFromTempToBackup(source, backup)
			time.Sleep(10 * time.Second)
		}
	}()
}


func StopBackup() {
	stop = true
}
