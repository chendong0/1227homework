package Lift

import (
	"testing"
)

func Test1(t *testing.T) {
	L := &Lift{
		TotalFloors:  5,
		CurrentFloor: 3,
	}

	if L.Direction != "" {
		t.Fatalf("stop：%s", L.Direction)

	}
}

func Test2(t *testing.T) {
	L := &Lift{
		TotalFloors:  5,
		CurrentFloor: 1,
	}
	if err := L.Get(3); err != nil {
		t.Fatalf("error：%v", err)
	}
	if L.Direction != "Up" {
		t.Fatalf("Up side but：%s", L.Direction)
	}
	L.LiftAct()
	if L.CurrentFloor != 3 {
		t.Fatalf("3 floor but：%d", L.CurrentFloor)
	}
}

func TestCase3(t *testing.T) {
	L := &Lift{
		TotalFloors:  5,
		CurrentFloor: 3,
	}
	if err := L.Get(4); err != nil {
		t.Fatalf("error：%v", err)
	}

	if err := L.Get(2); err != nil {
		t.Fatalf("error：%v", err)
	}
	if L.Direction != "Up" {
		t.Fatalf("Up side but：%s", L.Direction)
	}
	L.LiftAct()
	if L.CurrentFloor != 4 {
		t.Fatalf("4floor but：%d", L.CurrentFloor)
	}
	L.LiftAct()
	if L.CurrentFloor != 2 {
		t.Fatalf("2 floor but：：%d", L.CurrentFloor)
	}
}
func TestCase4(t *testing.T) {
	L := &Lift{
		TotalFloors:  5,
		CurrentFloor: 3,
	}
	if err := L.Get(4); err != nil {
		t.Fatalf("error：%v", err)
	}
	if err := L.Get(5); err != nil {
		t.Fatalf("error：%v", err)
	}
	if err := L.Get(2); err != nil {
		t.Fatalf("error：%v", err)
	}
	if L.Direction != "Up" {
		t.Fatalf("Up side but：%s", L.Direction)
	}
	L.LiftAct()
	if L.CurrentFloor != 4 {
		t.Fatalf("4floor but：：%d", L.CurrentFloor)
	}
	L.LiftAct()
	if L.CurrentFloor != 5 {
		t.Fatalf("5 floor but：：%d", L.CurrentFloor)
	}
	L.LiftAct()
	if L.CurrentFloor != 2 {
		t.Fatalf("2 floor but：：%d", L.CurrentFloor)
	}
}
