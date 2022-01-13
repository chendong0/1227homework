package main

/*
作业：实现电梯核心功能
电梯在没有人按按钮的时候停留在当前所在楼层不动
*/
/*案例一
楼层有 5 层，所有电梯楼层没有人请求电梯，电梯不动
案例二
楼层有 5 层，电梯在 1 层。三楼按电梯。电梯向三楼行进，并停在三楼。
案例三
楼层有 5 层，电梯在 3 层。上来一些人后，目标楼层： 4 楼、2 楼。电梯先向上到 4 楼，然后转头到 2 楼，最后停在 2 楼。
案例四
楼层有 5 层，电梯在 3 层。上来一些人后，目标楼层： 4 楼、5 楼、2 楼。电梯先向上到 4 楼，然后到 5 楼，之后转头到 2 楼，最后停在 2 楼。*/案例一
*/

import( "fmt"
"Lift"
"time"
)

func main()  {
	L1 :=Lift.Lift{
		3,
		[]int{},
		[]int{},
		time.Duration(time.Second),
		false,
		false,

	}
	Personnel:=[]Lift.Lift.Personnel{
		{
			3,
			4,
		},
		{
			3,
			5,
		},
		{
			3,
			2,
		},
	}
	fmt.Println("电梯在 3 层。上来一些人后，目标楼层： 4 楼、5 楼、2 楼")
	L1.Get(Personnel)
	r:= L1.Update（）
	fmt.Println("***********")
	for _, b := range r {
		fmt.Printf("%d Floor \n", b)
	}


}

