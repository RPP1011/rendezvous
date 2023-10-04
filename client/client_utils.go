package client

import (
	"time"

	"github.com/RPP1011/rendezvous/shared"
)

// Update drift every other minute
func (c *Client) update_drift() {
	select {
	case <-c.serverConnectionCancelChannel:
		// Cancel signal received, stop the goroutine
		return
	case <-c.lobbyConnectionCancelChannel:
		// Cancel signal received, stop the goroutine
		return
	default:
		c.timeDrift = shared.CalculateDrift(5)
		// Sleep for 2 minutes
		time.Sleep(time.Minute * 2)
	}

}
