package main

import (
	"math"
	"time"

	"github.com/epitchi/steambackup/steambackup"
)

func main() {
	steambackup.StartBackup(
		"C:/Users/thinkmay/AppData/Local/HoYoverse",
		"D:/Backup",
	)

	time.Sleep(time.Duration(math.MaxInt64))
}

