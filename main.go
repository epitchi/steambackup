package main

import "github.com/epitchi/steambackup/steambackup"

func main() {
	steambackup.StartBackup(
		"C:/Users/thinkmay/AppData/Local/HoYoverse",
		"D:/Backup",
	)
}
