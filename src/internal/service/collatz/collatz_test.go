package collatz

import "testing"

func TestCollatz(t *testing.T) {
	type args struct {
		factor int
	}
	tests := []struct {
		name      string
		factor    int
		wantSteps int
	}{
		{
			name:      "testing 0 as factor",
			factor:    0,
			wantSteps: 0,
		},
		{
			name:      "testing 1 as factor",
			factor:    1,
			wantSteps: 1,
		},
		{
			name:      "testing 2 as factor",
			factor:    2,
			wantSteps: 2,
		},
		{
			name:      "testing 3 as factor",
			factor:    3,
			wantSteps: 8,
		},
		{
			name:      "testing 4 as factor",
			factor:    4,
			wantSteps: 3,
		},
		{
			name:      "testing 5 as factor",
			factor:    5,
			wantSteps: 6,
		},
		{
			name:      "testing 6 as factor",
			factor:    6,
			wantSteps: 9,
		},
		{
			name:      "testing 19 as factor",
			factor:    19,
			wantSteps: 21,
		},
		{
			name:      "testing 27 as factor",
			factor:    27,
			wantSteps: 112,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSteps := Collatz(tt.factor); gotSteps != tt.wantSteps {
				t.Errorf("Collatz() = %v, want %v", gotSteps, tt.wantSteps)
			}
		})
	}
}
