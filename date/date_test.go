package date

import (
	"testing"
	"time"
)

//TestGetFistDayOfMonth-
func TestGetFistDayOfMonth(t *testing.T){
	fistDayOfMonth := GetFistDayOfMonth(time.Now())
	t.Logf("%v",fistDayOfMonth)
}

//TestGetLastDateOfMonth-
func TestGetLastDateOfMonth(t *testing.T){
	lastDateOfMonth := GetLastDateOfMonth(time.Now())
	t.Logf("%v",lastDateOfMonth)
}
