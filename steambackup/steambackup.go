package steambackup

import (
	"fmt"
	"time"
)

func CopyFromBackupToTemp(source, destination string) {
	// TODO: unzip latest backup
	ZipFolder("Backup.zip", source)

	CopyFolder("./Backup.zip", destination)

}

func CopyFromTempToBackup(source, destination string) {
	// TODO: zip Folder
	CopyFolder(source, destination)
	UnzipFolder("./Backup.zip", destination)
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

			fmt.Println("Backup done")

			time.Sleep(10 * time.Second)
		}
	}()
}

func StopBackup() {
	stop = true
}
