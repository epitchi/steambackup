package steambackup

import (
	"fmt"
	"time"
)

func CopyFromBackupToTemp(source, destination string) {
	// TODO: unzip latest backup
	CopyFolder(source+"/Zip/Backup.zip", destination)
	UnzipFolder("Backup.zip", destination)
	DeleteFile(destination + "/Backup.zip")
}

func CopyFromTempToBackup(source, destination string) {

	ZipFolder("Backup.zip", source)
	err := CopyFile("D:/code/steambackup/Backup.zip", destination + "/Backup.zip")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File backup successfully")
}
var (
	stop = false
)

func StartBackup(source, backup string) {
	fmt.Println("0. Start Backup")
	CopyFromBackupToTemp(backup, source)
	fmt.Println("1. Copy disk C to D DONE")

	go func() {
		for {
			if stop {
				break
			}

			fmt.Println("Backup disk C to D")

			CopyFromTempToBackup(source, backup)

			time.Sleep(10 * time.Second)
		}
	}()
}

func StopBackup() {
	stop = true
}
