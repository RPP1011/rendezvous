package shared

import (
	"github.com/beevik/ntp"
	"time"
)

func RemoteTime() time.Time {
	server := "time.google.com"

	remoteTime, err := ntp.Time(server)
	if err != nil {
		return time.Now()
	}
	return remoteTime
}

func CalculateDrift(trials int) int64 {
	// Calculate the average time drift
	var drift int64 = 0
	for i := 0; i < trials; i++ {
		// Get the local time
		local_time := time.Now()

		// Get the remote time
		remote_time := RemoteTime()

		// Calculate the time drift
		drift += remote_time.UnixNano() - local_time.UnixNano()

		// Sleep for 1 second
		time.Sleep(time.Millisecond * 500)
	}

	return drift / int64(trials)
}
