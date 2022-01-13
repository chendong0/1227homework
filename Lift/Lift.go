package Lift

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

import (
"fmt"
"sort"
"time"
)
//案例一
const TotalFloors int = 5

//案例二
//创建乘梯人员Personnel
type Personnel struct {
	persent      int
	Destinations int //可以考虑slice []int

}

type Lift struct { //结构体创建
	//TotalFloors 楼层总数5层,整型
	TotalFloors int
	//CurrentFloor现在楼层
	CurrentFloor int
	//direction方向(orientation)
	Direction string
	// upsideRequest上行提示
	UpsideRequest []int
	//downsideRequest下行提示
	DownsideRequest []int
	Speed           time.Duration
	Up              bool //上
	Down            bool //下

}

func (L *Lift) Get(getFloor int) error {
	if getFloor >= 1 && getFloor <= L.TotalFloors {
		if getFloor == L.CurrentFloor {
			//为当前楼层
			//err:= fmt.Errorf("电梯为当前楼层")
			//return err
			return nil

		}
		if len(L.UpsideRequest)+len(L.DownsideRequest) == 0 {
			L.Up(getFloor)
		}
		if getFloor > L.CurrentFloor {
			L.UpsideRequest = append(L.UpsideRequest, getFloor)
			sort.Ints(L.UpsideRequest)
		} else {
			L.DownsideRequest = append(L.DownsideRequest, getFloor)
			sort.Slice(L.DownsideRequest, func(u, d int) bool {
				return L.DownsideRequest[d] < L.DownsideRequest[u]
			})
		}
		return nil
	} else {
		return errors.New("*****")
	}
}

func (L *Lift) Update(getFloor int) {
	if getFloor > L.CurrentFloor {
		L.Direction = "Up"
	} else if getFloor < L.CurrentFloor {
		L.Direction = "Down"
	} else {
		L.Direction = "**"

	}

}
func (L *Lift) LiftAct() {
	if L.Direction == "Up" {
		if len(L.UpsideRequest) == 0 {
			if len(L.DownsideRequest) != 0 {
				L.Update(L.DownsideRequest[0])
				L.LiftAct()
			} else {
				L.Direction = "****"
				return
			}
		} else if L.Direction == "Down" {
			if len(L.DownsideRequest) == 0 {
				if len(L.UpsideRequest) != 0 {
					L.Update(L.UpsideRequest[0])
					L.LiftAct()
				} else {
					L.Direction == "***"
					return
				}

			} else {
				L.LiftAct(L.DownsideRequest[0])
				L.DownsideRequest = L.DownsideRequest[1:]

			}

		}

	}

}
func (L *Lift) Open {
	fmt.Printf("Open the door at NO. %d %t \n", L.CurrentFloor)

}
func (L *Lift) Close {
	fmt.Printf("Close the door at NO. %d %t \n", L.CurrentFloor)

}
func (L *Lift) LiftAct(opera int) {
	for {
		if L.CurrentFloor == opera {
			L.Open()
			time.Sleep(5 * time.Second)
			L.Close()
			return
		} else {
			if L.CurrentFloor > opera {
				time.Sleep(5 * time.Second)
				if L.CurrentFloor -= 1
			} else {
				time.Sleep(5 % *time.Second)
				L.CurrentFloor += 1
			}

		}

	}

}
