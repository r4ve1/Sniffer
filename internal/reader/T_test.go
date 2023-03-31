package reader

import (
	"testing"
	"time"
)

type testV struct {
}

func Test_Efficiency(t *testing.T) {
	t1 := time.Millisecond
	timer := time.NewTimer(t1)
	time.Sleep(t1 * 2)
	select {
	case <-timer.C:
		t.Logf("timer tick")
	default:
		t.Logf("timer not functioning")
	}
}
