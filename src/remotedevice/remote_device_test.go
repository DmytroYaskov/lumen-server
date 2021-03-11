package remotedevice

import "testing"

func TestHSL2RGB(t *testing.T) {

	type HSL struct {
		H, S, L float32
	}

	type RGB struct {
		R, G, B byte
	}

	var tests = []struct {
		HSL HSL
		RGB RGB
	}{
		{HSL{0, 0, 0}, RGB{0, 0, 0}},
		{HSL{0, 0, 1}, RGB{255, 255, 255}},
		{HSL{0, 1, 0.5}, RGB{255, 0, 0}},
		{HSL{120, 1, 0.5}, RGB{0, 255, 0}},
		{HSL{240, 1, 0.5}, RGB{0, 0, 255}},
		{HSL{60, 1, 0.5}, RGB{255, 255, 0}},
		{HSL{180, 1, 0.5}, RGB{0, 255, 255}},
		{HSL{300, 1, 0.5}, RGB{255, 0, 255}},
		{HSL{0, 0, 0.75}, RGB{191, 191, 191}},
		{HSL{0, 0, 0.5}, RGB{128, 128, 128}},
		{HSL{0, 1, 0.25}, RGB{128, 0, 0}},
		{HSL{60, 1, 0.25}, RGB{128, 128, 0}},
		{HSL{120, 1, 0.25}, RGB{0, 128, 0}},
		{HSL{300, 1, 0.25}, RGB{128, 0, 128}},
		{HSL{180, 1, 0.25}, RGB{0, 128, 128}},
		{HSL{240, 1, 0.25}, RGB{0, 0, 128}},
	}

	for _, test := range tests {

		var result RGB

		result.R, result.G, result.B = HSL2RGB(test.HSL.H, test.HSL.S, test.HSL.L)

		if test.RGB != result {
			t.Error("Test Failed: input {}, output{} - received {}", test.HSL, test.RGB, result)
		}
	}
}

func TestHSV2RGB(t *testing.T) {

	type HSV struct {
		H, S, V float32
	}

	type RGB struct {
		R, G, B byte
	}

	var tests = []struct {
		HSV HSV
		RGB RGB
	}{
		{HSV{0, 0, 0}, RGB{0, 0, 0}},
		{HSV{0, 0, 1}, RGB{255, 255, 255}},
		{HSV{0, 1, 1}, RGB{255, 0, 0}},
		{HSV{120, 1, 1}, RGB{0, 255, 0}},
		{HSV{240, 1, 1}, RGB{0, 0, 255}},
		{HSV{60, 1, 1}, RGB{255, 255, 0}},
		{HSV{180, 1, 1}, RGB{0, 255, 255}},
		{HSV{300, 1, 1}, RGB{255, 0, 255}},
		{HSV{0, 0, 0.75}, RGB{191, 191, 191}},
		{HSV{0, 0, 0.5}, RGB{128, 128, 128}},
		{HSV{0, 1, 0.5}, RGB{128, 0, 0}},
		{HSV{60, 1, 0.5}, RGB{128, 128, 0}},
		{HSV{120, 1, 0.5}, RGB{0, 128, 0}},
		{HSV{300, 1, 0.5}, RGB{128, 0, 128}},
		{HSV{180, 1, 0.5}, RGB{0, 128, 128}},
		{HSV{240, 1, 0.5}, RGB{0, 0, 128}},
	}

	for _, test := range tests {

		var result RGB

		result.R, result.G, result.B = HSV2RGB(test.HSV.H, test.HSV.S, test.HSV.V)

		if test.RGB != result {
			t.Error("Test Failed: input {}, output{} - received {}", test.HSV, test.RGB, result)
		}
	}
}
