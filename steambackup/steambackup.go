package steambackup

import (
	"fmt"
	"time"
)

func CopyFromDestinationToTemp(source string, destination string) {
	// TODO: unzip latest backup
	CopyFolder(source, destination)
}

func CopyFromTempToDestination(source string, destination string) {
	// TODO: zip Folder
	CopyFolder(source, destination)
}

var (
	stop = false
)

func StartBackup(source string, backup string) {
	fmt.Println("0. Start Backup")
	CopyFromDestinationToTemp(backup, source)
	fmt.Println("1. Copy disk C to D DONE")

	go func() {
		for {
			if stop {
				break
			}

			fmt.Println("Backup disk C to D")

			CopyFromTempToDestination(source, backup)

			fmt.Println("Backup done")

			time.Sleep(10 * time.Second)
		}
	}()
}

func StopBackup() {
	stop = true
}
