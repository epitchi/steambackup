package steambackup

import "time"

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
	CopyFromDestinationToTemp(backup, source)

	go func() {
		for {
			if stop {
				break
			}

			CopyFromTempToDestination(source, backup)
			time.Sleep(60 * time.Second)
		}
	}()
}

func StopBackup() {
	stop = true
}
