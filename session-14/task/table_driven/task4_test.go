package table_driven

import (
	"testing"
)

type TestCases2 struct {
	testName        string
	testTemperature float64
	testScale       string
	expectedValue   float64
}

func TestConvertTemperature(t *testing.T) {
	testCases := []TestCases2{
		{
			testName:        "Convert to F, 0",
			testTemperature: 0,
			testScale:       "C",
			expectedValue:   32,
		},
		{
			testName:        "Convert to F, 100",
			testTemperature: 100,
			testScale:       "C",
			expectedValue:   212,
		},
		{
			testName:        "Convert to F, -40",
			testTemperature: -40,
			testScale:       "C",
			expectedValue:   -40,
		},
		{
			testName:  "Check the wrong scale",
			testScale: "wrong_scale",
		},
		{
			testName:        "Check the out of range temperatures",
			testTemperature: 200,
			testScale:       "C",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t *testing.T) {
			if res, err := ConvertTemperature(testCase.testTemperature, testCase.testScale); err != nil {
				t.Error("error ocurred:", err)
			} else {
				if res != testCase.expectedValue {
					t.Errorf("expected value not satisfied. should be %f, is %f", testCase.expectedValue, res)
				}
			}
		})
	}
}
