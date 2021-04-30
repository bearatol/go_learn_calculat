package operations

import "testing"

func TestRaznost(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
	}
	tests := []struct {
		name    string
		args    args
		wantRes float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Raznost(tt.args.a, tt.args.b, tt.args.c); gotRes != tt.wantRes {
				t.Errorf("Raznost() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
