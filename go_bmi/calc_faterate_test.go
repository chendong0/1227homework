package go_bmi

import "testing"

func TestCalcFatRate(t *testing.T) {
	inputfatRate, inputage,inputsex := 1.0, 1,"1"
	expectedOutput := 1.0
	t.Logf("将要计算，输入：heightM: %f, age: %d, 期望结果：%f", inputfatRate, inputage, expectedOutput)
	actualOutput, err := CalcFatRate(inputfatRate,inputage,inputsex)
	t.Logf("实际得到：%f, error: %v", actualOutput, err)
	if err != nil {
		t.Fatalf("expecting no err, but got: %v", err)
	}
	if expectedOutput != actualOutput {
		t.Errorf("expecting %f, but got %f", expectedOutput, actualOutput)
	}
}
