package timeago

import (
	"fmt"
	"time"
)

func TimeAgo(ago int64) string {

	timeDiff := (time.Now().Unix() - ago)
	// TODO:: this should be fixed
	if timeDiff < 60 {
		return fmt.Sprintf("%ds", timeDiff)
	} else if timeDiff < 3600 {
		return fmt.Sprintf("%dm", int64(timeDiff/60))
	} else if timeDiff < (86400) {
		return fmt.Sprintf("%dh", int64(timeDiff/3600))
	} else if timeDiff < (604800) {
		return fmt.Sprintf("%dd", int64(timeDiff/86400))
	} else if timeDiff < (2628000) {
		return fmt.Sprintf("%dw", int64(timeDiff/604800))
	} else if timeDiff < (31536000) {
		return fmt.Sprintf("%dM", int64(timeDiff/2628000))
	}
	return fmt.Sprintf("%dy", int64(timeDiff/31536000))
}
