package timeago

import (
	"fmt"
	"time"
)

// TimeAgo returns a string
// that calculates time ago
func TimeAgo(ago int64) string {
	timeDiff := (time.Now().Unix() - ago)

	// TODO:: this should be fixed
	if timeDiff < 60 {
		return fmt.Sprintf("%d s", timeDiff)
	}

	if timeDiff < 3600 {
		return fmt.Sprintf("%d min", int64(timeDiff/60))
	}

	if timeDiff < (86400) {
		return fmt.Sprintf("%d hr", int64(timeDiff/3600))
	}

	if timeDiff < (604800) {
		return fmt.Sprintf("%d d", int64(timeDiff/86400))
	}

	if timeDiff < (2628000) {
		return fmt.Sprintf("%d wk", int64(timeDiff/604800))
	}

	if timeDiff < (31536000) {
		return fmt.Sprintf("%d Mo", int64(timeDiff/2628000))
	}

	return fmt.Sprintf("%d y", int64(timeDiff/31536000))
}
