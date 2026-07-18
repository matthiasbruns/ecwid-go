package webhook

import "testing"

func TestClassify(t *testing.T) {
	tests := []struct {
		status        int
		wantDelivered bool
	}{
		// Ecwid's success set.
		{200, true}, {201, true}, {202, true}, {204, true}, {209, true},
		// 2xx traps that are NOT success.
		{203, false}, {208, false}, {206, false},
		// Every 3xx fails — the redirect trap.
		{301, false}, {302, false}, {307, false},
		// Ordinary failures.
		{400, false}, {404, false}, {500, false}, {199, false},
	}
	for _, tt := range tests {
		got := classify(tt.status)
		if got.Delivered != tt.wantDelivered {
			t.Errorf("classify(%d).Delivered = %t, want %t", tt.status, got.Delivered, tt.wantDelivered)
		}
		if got.Reason == "" {
			t.Errorf("classify(%d).Reason is empty; the reason must always be reported", tt.status)
		}
	}
}
