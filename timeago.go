package timeago

import (
	"fmt"
	"time"
)

func UnixToAgo(unix int64) string {

	ago := (time.Now().Unix() - unix)
	if ago < 60 {
		return fmt.Sprintf("%ds ago", ago)
	} else if ago < 3600 {
		return fmt.Sprintf("%dm ago", int64(ago/60))
	} else if ago < (86400) {
		return fmt.Sprintf("%dh ago", int64(ago/3600))
	} else if ago < (604800) {
		return fmt.Sprintf("%dd ago", int64(ago/86400))
	} else if ago < (2628000) {
		return fmt.Sprintf("%dw ago", int64(ago/604800))
	} else if ago < (31536000) {
		return fmt.Sprintf("%dM ago", int64(ago/2628000))
	}
	return fmt.Sprintf("%dy ago", int64(ago/31536000))
}
