package main

import (
	"flag"
	"fmt"
)

var apples, cakes int
var name string

func main() {
	flag.IntVar(&apples, "apple", 25, "how many apples")
	flag.IntVar(&cakes, "cake", 20, "how many cakes")
	flag.StringVar(&name, "name", "Ainun", "the owner name")
	flag.Parse()

	txt := "\n%s has %d cakes and %d apples. %s wants to bundle those cakes and apples into boxes and give them\n"
	fmt.Println(fmt.Sprintf(txt, name, cakes, apples, name))

	box := NewBox(apples, cakes)
	boxes := box.HowMany()
	apple, cake := box.CountItemsEachBox()

	txtManyBox := "- How many boxes that %s can make? %d boxes"
	txtManyCakesAndApples := "- How many cakes and apples every box have? %d apples and %d cakes"
	fmt.Println(fmt.Sprintf(txtManyBox, name, boxes))
	fmt.Println(fmt.Sprintf(txtManyCakesAndApples, apple, cake))
}

func findGCD(a int, b int) (gcd int) {
	// using Euclidean algorithm
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}

	gcd = a
	return
}

type Box struct {
	boxes int
	Items struct {
		Apple int
		Cake  int
	}
}

func NewBox(apple int, cake int) *Box {
	return &Box{
		Items: struct {
			Apple int
			Cake  int
		}{
			Apple: apple,
			Cake:  cake,
		},
	}
}

func (box *Box) HowMany() int {
	return findGCD(box.Items.Apple, box.Items.Cake)
}

func (box *Box) CountItemsEachBox() (apple int, cake int) {
	boxes := box.HowMany()
	apple = box.Items.Apple / boxes
	cake = box.Items.Cake / boxes
	return
}
