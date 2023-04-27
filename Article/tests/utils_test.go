package test

import (
	"Article/utils"
	"reflect"
	"strconv"
	"testing"
)

func TestStrToInt(t *testing.T) {
	arr := []string{}
	want := []int{}
	for i := 1; i < 20; i++ {
		arr = append(arr, strconv.Itoa(i))
		want = append(want, i)
		i = i*10 + 1
	}
	got := []int{}
	for _, i := range arr {
		got = append(got, utils.StrToInt(i))
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("except:%v,got:%v", want, got)
	}

}
