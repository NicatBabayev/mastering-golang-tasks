package table_driven

import "testing"

type TestCases struct {
	name      string
	testValue string
}

func TestReverseString(t *testing.T) {
	testCases := []TestCases{
		{
			name:      "Check Empty String",
			testValue: "",
		},
		{
			name:      "Check Palindrome Case",
			testValue: "radar",
		},
		{
			name:      "Special character",
			testValue: "D@@R",
		},
		{
			name:      "Ordinary word reversal",
			testValue: "koroglu",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if _, err := ReverseString(testCase.testValue); err != nil {
				t.Error("Special case occurred: ", err)
			}
		})
	}

}
