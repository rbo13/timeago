package timeago

import (
	"testing"
)

func TestUnixToAgo(t *testing.T) {

	want := "4m ago"
	ago := 1524922924
	got := TimeAgo(int64(ago))

	if want != got {
		t.Errorf("Want '%s' but got '%s'", want, got)
	}

}
