package humanize

import (
	"testing"
	"time"
)

func TestLangZhCN(t *testing.T) {
	humanizer, err := New(`zh-CN`)
	if err != nil {
		t.Fatal(err)
	}
	timeStart, err := time.Parse(`2006-01-02 15:04:05`, `2002-09-01 09:30:18`)
	if err != nil {
		t.Fatal(err)
	}
	timeEnd, err := time.Parse(`2006-01-02 15:04:05`, `2012-02-01 20:00:10`)
	if err != nil {
		t.Fatal(err)
	}
	expected := `9年, 6个月, 2周, 6天, 10小时, 29分钟, 52秒以前`
	humanized := humanizer.TimeDiff(timeEnd, timeStart, -1)
	if humanized != expected {
		t.Errorf("Expected '%s', got '%s'.", expected, humanized)
	}
}
