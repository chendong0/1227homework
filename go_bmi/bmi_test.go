package go_bmi

import "testing"

func TestBMI(t *testing.T) {
	inputheightM, inputweightKG := 1.0, 1.0
	expectedOutput := 1.0
	t.Logf("将要计算，输入：heightM: %f, weightKG: %f, 期望结果：%f", inputheightM, inputweightKG, expectedOutput)
	actualOutput, err := BMI(inputheightM, inputweightKG)
	t.Logf("实际得到：%f, error: %v", actualOutput, err)
	if err != nil {
		t.Fatalf("expecting no err, but got: %v", err)
	}
	if expectedOutput != actualOutput {
		t.Errorf("expecting %f, but got %f", expectedOutput, actualOutput)
	}
}
