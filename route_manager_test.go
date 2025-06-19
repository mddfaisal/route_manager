package routemanager

import "testing"

func TestRouteManager_AllFinePaths(t *testing.T) {
	tests := []struct {
		name   string
		routes map[string]string
	}{
		{
			name: "Test 1",
			routes: map[string]string{
				"/A/B/F/N":   "N",
				"/A/B/G/R":   "R",
				"/A/B/G/Q":   "Q",
				"/A/C/H/S":   "S",
				"/A/C/I/T":   "T",
				"/A/C/J":     "J",
				"/A/D/K/U":   "U",
				"/A/D/K/V":   "V",
				"/A/D/L/W":   "W",
				"/A/D/M/X":   "X",
				"/A/D/M/Y/Z": "Z",
				"/B/M/N":     "N",
				"/C/H/S/P":   "P",
				"/C/H/S/Q":   "Q",
				"/C/H/M/R/P": "P",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for route, expected := range tt.routes {
				r := &Node{}
				Insert(r, route)
				if got, _ := Search(r, route); got != expected {
					t.Errorf("Search(%q) = %q; want %q", route, got, expected)
				}
			}
		})
	}
}

func TestRouteManager_InvalidPaths(t *testing.T) {
	tests := []struct {
		name   string
		routes map[string]string
	}{
		{
			name: "Test Invalid Paths",
			routes: map[string]string{
				"/A/B/C/D/E/F/G/H/I/J/K/L/M/N/O/P/Q/R/S/T/U/V/W/X/Y/Z": "Z",
				"/A/B/C/D":    "D",
				"/A/B/C/:D":   "panic",
				"/A/B/C/:D/K": "panic",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for route := range tt.routes {
				r := &Node{}
				Insert(r, route)
			}
		})
	}
}

// Test for dynamic segments
func TestRouteManager_DynamicSegments(t *testing.T) {
	r := &Node{}
	if err := Insert(r, "/A/B/:C"); err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	if got, _ := Search(r, "/A/B/D"); got != "D" {
		t.Errorf("Search(/A/B/D) = %q; want D", got)
	}
}
