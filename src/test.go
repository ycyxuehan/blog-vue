package main

import (
	"reflect"
	"fmt"
)

var poll = make(chan int, 10)

func thread(){
	for {
		select {
		case i := <- poll :
				fmt.Println(i)
		}
	}
}

type I interface {
	Show()
}

type S struct {
	A string
}

func (s *S)Show(){
	fmt.Println("ssssssssssssssss")
}

func makeslice(i I){
	v := reflect.ValueOf(i)
	
	fmt.Println(v.Interface())
}

func main() {
	s := S{A:"sasasdadad"}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	slice := reflect.MakeSlice(reflect.SliceOf(t), 10, 10)
	ns := reflect.New(slice.Type())
	ns.Elem().Set(slice)
	ns.Index(1).Set(v)
	// sv := reflect.MakeSlice(t, 1, 2)
	// sv = reflect.AppendSlice(sv, v)
	fmt.Println(ns)
}