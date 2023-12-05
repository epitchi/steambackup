package steambackup

import (
	"fmt"
	"time"
)

func CopyFromBackupToTemp(source, destination string) (error) {
	UnzipFolder(source+"Backup.zip", destination)
	return nil
}

func CopyFromTempToBackup(source, destination string) {
	ZipFolder(destination+"Backup.zip", source)
}

var (
	stop = false
)

func StartBackup(source, backup string) {
	fmt.Println("SteamBackup: Start Backup")
	err := CopyFromBackupToTemp(backup, source)
	if err != nil{
		fmt.Printf("%s\n",err.Error())
	}
	fmt.Println("SteamBackup: Copy disk D to C DONE")

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
