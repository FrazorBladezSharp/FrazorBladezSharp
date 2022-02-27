package main

import "fmt"

type position struct {
	x float32
	y float32
}

type person struct {
	name   string
	health int
	pos    position
}

func main() {

	var a int
	var pos position

	a = welcome("Basics") // nb passes a copy - avoids side effects
	fmt.Printf("value of a = %v \n", a)

	// we can use a ptr to avoid copying large amounts of data
	aPtr := &a
	addOne(aPtr) // by passing a pointer we avoid returning a value
	fmt.Printf("value of a = %v \n", a)

	pos.x = 4
	pos.y = 5
	fmt.Println(pos.x, pos.y)
	//can still pass a pointer
	var y int
	y = int(pos.y)
	addOne(&y)

	p := position{2, 3}
	fmt.Println(p)

	enemy := person{"Assassin", 10, p}
	player := person{"Bladez", 10, position{10, 11}}
	fmt.Println(player, enemy)

	// pointers &player == *player
	playerPtr := &player
	fmt.Println(playerPtr)
}

func errorCheck(err error) {
	if err != nil {
		fmt.Println("[welcome] : error")
	}
}

func addOne(num *int) {
	*num++
}

func welcome(name string) int {
	num, err := fmt.Printf("Welcome to %s \n", name)
	errorCheck(err)

	return num
}
