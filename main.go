package main

import (
	"Homework/teacher_example/009/fatrate_refactor/calc"
	"fmt"
	"runtime/debug"
)

func init() {
	fmt.Println("我是第一个init函数>>>>>")
}

func main() {
	for {
		mainFatRateBody()
		if cont := whetherContinue(); !cont {
			break

		}
	}
}

func init() {
	fmt.Println("我是第二个init函数<<<<<")

}

func recoverMainBody() {
	if re := recover(); re != nil {
		fmt.Printf("warning:catch critical error: %v\n", re)
		debug.PrintStack()
	}

}

func mainFatRateBody() {
	defer recoverMainBody() //defer语句,延迟语句被用于一个函数调用,在这个函数之前,延迟语句返回.
	//成功捕获
	weight, tall, age, _, sex := getMaterialsFromInput(58.0, 1.8, 30, 20, "男") //not enough arguments in call to getMaterialsFromInput
	//have ()
	//want (float64, float64, int, int, string)
	//58.0,1.8,30,20,"男"
	//too many arguments to return
	fatRate, err := calcFatRate(weight, tall, age, sex)
	if err != nil {//err !=nil应补充err error作为返回值
		fmt.Println("warning: 输入的数据有误,计算体脂率出错,不能再继续", err)

	}

	if fatRate <= 0 {
		panic("fat rate is not allowed to be negative脂肪率不能为负数!")
	}
	var checkHealthinessFunc func(age int, fatRate float64)
	if sex == "男" {
		checkHealthinessFunc = getHealthinessSuggestionsForMale
	} else {
		checkHealthinessFunc = getHealthinessSuggestionsForFemale
	}
	getHealthinessSuggestions(age, fatRate, checkHealthinessFunc)
}

func getHealthinessSuggestions(age int, fatRate float64, getSuggestion func(age int, fatRate float64)) {
	getSuggestion(age, fatRate)
}

func generateCheckHealthinessForMale() func(age int, fatRate float64) {
	//定制功能

	return func(age int, fatRate float64) {

	}

}

func getHealthinessSuggestionsForMale(age int, fatRate float64) {
	//编写男性的体脂率与体脂状态表
	if age >= 18 && age <= 39 {
		if fatRate <= 0.1 {
			fmt.Println("当前是:偏瘦.要多吃多锻炼,增强体脂")
		} else if fatRate > 0.1 && fatRate <= 0.16 {
			fmt.Println("当前是:标准.太棒了,要继续保持")
		} else if fatRate > 0.16 && fatRate <= 0.21 {
			fmt.Println("当前是:偏胖.吃饱后多多散步,有助于消化")
		} else if fatRate > 0.21 && fatRate <= 0.26 {
			fmt.Println("当前是:肥胖.少吃高热量食物,增加运动量")
		} else {
			fmt.Println("当前是:非常肥胖.需要多锻炼增加运动时间,每天多运动")
		}
	} else if age >= 40 && age <= 59 {
		if fatRate <= 0.1 {
			fmt.Println("当前是:偏瘦.要多吃多锻炼,增强体脂")
		} else if fatRate > 0.1 && fatRate <= 0.16 {
			fmt.Println("当前是:标准.太棒了,要继续保持")
		} else if fatRate > 0.16 && fatRate <= 0.21 {
			fmt.Println("当前是:偏胖.吃饱后多多散步,有助于消化")
		} else if fatRate > 0.21 && fatRate <= 0.26 {
			fmt.Println("当前是:肥胖.少吃高热量食物,增加运动量")
		} else {
			fmt.Println("当前是:非常肥胖.需要多锻炼增加运动时间,每天多运动")
		}
	} else if age >= 60 {
		if fatRate <= 0.1 {
			fmt.Println("当前是:偏瘦.要多吃多锻炼,增强体脂")
		} else if fatRate > 0.1 && fatRate <= 0.18 {
			fmt.Println("当前是:标准.太棒了,要继续保持")
		} else if fatRate > 0.18 && fatRate <= 0.23 {
			fmt.Println("当前是:偏胖.吃饱后多多散步,有助于消化")
		} else if fatRate > 0.23 && fatRate <= 0.28 {
			fmt.Println("当前是:肥胖.少吃高热量食物,增加运动量")
		} else {
			fmt.Println("当前是:非常肥胖.需要多锻炼增加运动时间,每天多运动")
		}
	} else {
		fmt.Println("我们不评估18岁以下,因为波动太大,效果没有意义.")
	}

}

func getHealthinessSuggestionsForFemale(age int, fatRate float64) {
	//编写男性的体脂率与体脂状态表
	if age >= 18 && age <= 39 {
		if fatRate <= 0.1 {
			fmt.Println("当前是:偏瘦.要多吃多锻炼,增强体脂")
		} else if fatRate > 0.1 && fatRate <= 0.16 {
			fmt.Println("当前是:标准.太棒了,要继续保持")
		} else if fatRate > 0.16 && fatRate <= 0.21 {
			fmt.Println("当前是:偏胖.吃饱后多多散步,有助于消化")
		} else if fatRate > 0.21 && fatRate <= 0.26 {
			fmt.Println("当前是:肥胖.少吃高热量食物,增加运动量")
		} else {
			fmt.Println("当前是:非常肥胖.需要多锻炼增加运动时间,每天多运动")
		}
	} else if age >= 40 && age <= 59 {
		if fatRate <= 0.1 {
			fmt.Println("当前是:偏瘦.要多吃多锻炼,增强体脂")
		} else if fatRate > 0.1 && fatRate <= 0.16 {
			fmt.Println("当前是:标准.太棒了,要继续保持")
		} else if fatRate > 0.16 && fatRate <= 0.21 {
			fmt.Println("当前是:偏胖.吃饱后多多散步,有助于消化")
		} else if fatRate > 0.21 && fatRate <= 0.26 {
			fmt.Println("当前是:肥胖.少吃高热量食物,增加运动量")
		} else {
			fmt.Println("当前是:非常肥胖.需要多锻炼增加运动时间,每天多运动")
		}
	} else if age >= 60 {
		if fatRate <= 0.1 {
			fmt.Println("当前是:偏瘦.要多吃多锻炼,增强体脂")
		} else if fatRate > 0.1 && fatRate <= 0.18 {
			fmt.Println("当前是:标准.太棒了,要继续保持")
		} else if fatRate > 0.18 && fatRate <= 0.23 {
			fmt.Println("当前是:偏胖.吃饱后多多散步,有助于消化")
		} else if fatRate > 0.23 && fatRate <= 0.28 {
			fmt.Println("当前是:肥胖.少吃高热量食物,增加运动量")
		} else {
			fmt.Println("当前是:非常肥胖.需要多锻炼增加运动时间,每天多运动")
		}
	} else {
		fmt.Println("我们不评估18岁以下,因为波动太大,效果没有意义.")
	}

}

func calcFatRate(weight float64, tall float64, age int, sex string) ( fatRate float64, err error){
	//计算体脂率
	bmi, err := calc.CalcBMI(tall, weight)
	if err != nil{
		return 0,err
	}

	fatRate = calc.CalcFatRate(bmi, age, sex)
	fmt.Printf("体脂率是:%.2f \n", fatRate)
	return fatRate, err
}

func getMaterialsFromInput(float64, float64, int, int, string) (float64, float64, int, int, string) { //(float64, float64, int, int, string)
	//返回参数列表未输入,报错.(float64, float64, int, int, string)
	//录入各项指标数据
	var name string
	fmt.Print("姓名:")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重(千克):")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高:")
	fmt.Scanln(&tall)
	var age int
	fmt.Print("年龄:")
	fmt.Scanln(&age)

	var sexWeight int
	sex := "男"
	fmt.Print("性别(男/女)")
	fmt.Scanln(&sex)

	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	return weight, tall, age, sexWeight, sex
}

func whetherContinue() bool {
	var whetherContinue string
	fmt.Print("是否要记录下一个(yes/no)?")
	fmt.Scanln(&whetherContinue)
	if whetherContinue != "yes" {
		return false
	}
	return true

}
