package test

import (
	"github.com/dkys/etime"
	"log"
	"testing"
)

func TestAddRow(t *testing.T) {
	start, end := etime.Today()
	start1, end1 := etime.Month()
	start2, end2 := etime.MonthByNum(3)
	log.Printf("Today :%v %v\n", start, end)
	log.Printf("Month :%v %v\n", start1, end1)
	log.Printf("MonthByNum :%v %v\n", start2, end2)

}
