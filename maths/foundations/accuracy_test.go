package foundations

import "testing"

func TestDegreesOfAccuracy(t *testing.T) {
	tests := []struct {
		name      string
		target    float64
		rng       float64
		wantLower float64
		wantUpper float64
	}{
		{
			name:      "basic case",
			target:    10.0,
			rng:       2.0,
			wantLower: 9.0,
			wantUpper: 11.0,
		},
		{
			name:      "zero range",
			target:    5.0,
			rng:       0.0,
			wantLower: 5.0,
			wantUpper: 5.0,
		},
		{
			name:      "negative target",
			target:    -10.0,
			rng:       4.0,
			wantLower: -12.0,
			wantUpper: -8.0,
		},
		{
			name:      "small decimal",
			target:    10.0,
			rng:       0.1,
			wantLower: 9.95,
			wantUpper: 10.05,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DegreesOfAccuracy(tt.target, tt.rng)
			if got.LowerBound != tt.wantLower {
				t.Errorf("LowerBound = %v, want %v", got.LowerBound, tt.wantLower)
			}
			if got.UpperBound != tt.wantUpper {
				t.Errorf("UpperBound = %v, want %v", got.UpperBound, tt.wantUpper)
			}

			got.Print()
		})
	}
}
