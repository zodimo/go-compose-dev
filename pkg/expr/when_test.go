package expr

import (
	"testing"
)

func TestWhen(t *testing.T) {
	tests := []struct {
		name       string
		value      interface{}
		conditions interface{}
		otherwise  interface{}
		want       interface{}
	}{
		{
			name:  "match integer",
			value: 1,
			conditions: map[int]string{
				1: "one",
				2: "two",
			},
			otherwise: "unknown",
			want:      "one",
		},
		{
			name:  "no match integer",
			value: 3,
			conditions: map[int]string{
				1: "one",
				2: "two",
			},
			otherwise: "unknown",
			want:      "unknown",
		},
		{
			name:  "match string",
			value: "hello",
			conditions: map[string]int{
				"hello": 5,
				"world": 6,
			},
			otherwise: 0,
			want:      5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got interface{}
			switch v := tt.value.(type) {
			case int:
				got = When(v, tt.conditions.(map[int]string), tt.otherwise.(string))
			case string:
				got = When(v, tt.conditions.(map[string]int), tt.otherwise.(int))
			}

			if got != tt.want {
				t.Errorf("When() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwitchEquivalence(t *testing.T) {
	// Scenario: HTTP Status Codes
	code := 200

	// Standard Switch
	var statusSwitch string
	switch code {
	case 200:
		statusSwitch = "OK"
	case 404:
		statusSwitch = "Not Found"
	default:
		statusSwitch = "Unknown"
	}

	// When Expression
	statusWhen := When(code, map[int]string{
		200: "OK",
		404: "Not Found",
	}, "Unknown")

	if statusSwitch != statusWhen {
		t.Errorf("Switch value %q != When value %q", statusSwitch, statusWhen)
	}
}
