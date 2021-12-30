package go_bmi

import "fmt"

func BMI(weightKG, heightM float64) (bmi float64, err error) {
	if weightKG < 0 {
		err = fmt.Errorf("weight cannot be negative体重不能为负数")
		return
	}
	if weightKG == 0 {
		err = fmt.Errorf("weight cannot be 0体重不能为0")//按作业要求,返回错误
	}
	if heightM < 0 {
		err = fmt.Errorf("height cannot be negative身高不能为负数")
		return
	}
	if heightM == 0 {
		err = fmt.Errorf("height cannot be 0体重不能为0")
		return
	}
	bmi = weightKG / (heightM * heightM)
	return
}