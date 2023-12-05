package steambackup

import (
	"testing"
	"math"
	"time"

)


func TestBackup(t *testing.T) {
	StartBackup(
		"C:/AtlasA",
		"D:/",
	)

	time.Sleep(time.Duration(math.MaxInt64))
}