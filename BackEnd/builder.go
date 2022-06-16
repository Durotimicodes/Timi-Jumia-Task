package main

import "fmt"

func main() {

	me := aHuman().setAge(age:31).setHeight(height:187.5).setWeigth(weight:90.7).setEyeColor(eyeColor:"Brown")
	
	log.Printf("The human ==> %v",me)

}

type human struct {
	age      int
	height   float64
	weight   float64
	eyeColor string
}

func aHuman() human {
	return human{}
}

func (h human) setAge(age int) human{
	h.age = age
	return h
}

func (h human) setHeight(height float64) human{
	h.height = height
	return h
}

func (h human) setWeigth(weight float64) human{
	h.weight = weight
	return h
}

func (h human) setEyeColor(eye string) human{
	h.eyeColor = eye
	return h
}

