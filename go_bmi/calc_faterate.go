package go_bmi

import "fmt"

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64, err error) {
	//返回值列表要err error与err对应
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	if sex != "男" && sex != "女" {
		err = fmt.Errorf("sex is female or male,非男女的性别输入")
	}

	if bmi <= 0 {
		err = fmt.Errorf("bmi cannot be 0/negative体重不能为0/负数") //unresolved reference 'err'
	} //录入非法 BMI、年龄、性别（0、负数、超过 150 的年龄、非男女的性别输入），返回错误；
	if age <= 0 && age >= 150 {
		err = fmt.Errorf("age cannot be 0 or 150,年龄不能为0或超过150年")
	}

	fatRate = (1.2*bmi + getAgeWeight(age)*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}

func getAgeWeight(age int) (ageWeight float64) {
	ageWeight = 0.23
	if age >= 30 && age <= 40 {
		ageWeight = 0.22
	}
	return
}
