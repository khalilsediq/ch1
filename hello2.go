// package main
// import "fmt"
// func start() {
// fmt.Println("Hello world! This is khalilullah")
// }

package main

import "fmt"

func main() {
	// fmt.Println("Hello\nWorld")
	// fmt.Println("Greetings and\nSalutations")

	// var name = "khalilullah"
	// var age = 20
	// var num1 = 50
	// var num2 = 30
	// fmt.Println(num1 + num2)
	// fmt.Println(name, age)

	// var num1 float32 = 40.4
	// var num2 int8 = 60
	// fmt.Println(num1  , num2)

	// var text = `I am khalilullah,
	//  I am a software Engineer and the best programmer in the world.
	//   I am literally unstopable`
	// fmt.Println(text)

	// var num int8 = 20
	// var num2 = 10
	// var num3 = 4.5
	// var num4 = 5.5
	// fmt.Println(num3 - num4)

	// fmt.Println(num + int8(num2) + int8(num3))

	// var x = 18
	// if 10 >= 18 {
	// 	fmt.Println("you are eligible, for the club.")

	// } else {
	// 	fmt.Println("Sorry, you are not eligible for the Club")
	// }

	// var x = int  <= this is declared as a 0 value.

	// var x, y int = 20, 50 <= we can declare multiple variables at once and they can be of the same type.

	// var x, y int //All zero values of the same types

	// var x, y = 20, "Khalilullah"  // <= All values of different types

	// fmt.Println(x, y)

	// var (
	// 	x    int
	// 	y        = 20
	// 	z    int = 40
	// 	d, e     = 50, "Khalilullah"
	// 	f, g string
	// )
	// fmt.Println(x, y, z, d, e, f, g)

	// x := 40
	// // x, y := 30, "khalilThePro"
	// x, y := 50, `ХалилПрограммист` // <=  We can only do this on := this. ANd we can not achieve this by the = sign.

	// fmt.Println(x, y)
	//}

	// const x = 49

	// const (
	// 	idKey   = "id"
	// 	nameKey = "name"
	// )

	// const z = 3 * 3

	// func main() {
	// var y = "hello"
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
	// var b = x + 1
	// var f = "bye"
	// fmt.Println(b)
	// fmt.Println(f)
	// y = "bye"
	// fmt.Println(y)

	// const price = 80
	// var newPrice float64 = price * 2
	// fmt.Println(newPrice)

	// x := 10
	// x = 20
	// fmt.Println(x)
	// x = 30

	// var y = 89

	// _0 := 0_0
	// _𝟙 := 20
	// π := 3
	// ａ := "hello" // Unicode U+FF41
	// fmt.Println(_0)
	// fmt.Println(_𝟙)
	// fmt.Println(π)
	// fmt.Println(ａ)

	// var name = 0
	// fmt.Println(name)

	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i)
	// }
	checkArr()
}

// func calculateScope() {
// 	var basicInt int = 400
// 	var strictInt int64 = 400

// 	if basicInt == strictInt {
// 		fmt.Println("Match Found")
// 	}
// }

func checkArr() {
	// var v [3]int
	// v[1] = 99
	// fmt.Print(v[0])
	// fmt.Print(v[1])
	// var x = [3]int{10, 20, 30}
	// fmt.Println(x)
	// fmt.Println(v)

	// var x [3]int
	// x[1] = 99

	// fmt.Println(x[0])
	// fmt.Println(x[1])
	// fmt.Println(x)

	// var z = [...]int{1, 2, 3, 10: 7, 19: 100, 20}
	// fmt.Println(z)

	// var z = [...]int{10, 20, 30, 40}
	// fmt.Println(z[4])
	// var z = [...]int{10, 20, 30, 40}
	// fmt.Println(z[4])

	// var finalBoss = [...]int{1, 5: 10, 2, 99}
	// // var hero [7]int

	// // hero = finalBoss
	// fmt.Println(len(finalBoss))

	// var finalCheck = [...]int{4: 10, 9: 99}
	// fmt.Println(finalCheck)

	// var arr = [4]int{1, 2, 3}
	// var arr2 = [4]int{1, 2, 3, 4}
	// arr = arr2
	// fmt.Println(arr, arr2)
	// fmt.Println(arr)

	// var myArray = [3]int{1, 2, 3}
	// var mySlice = []int{1, 2, 3}
	// fmt.Println(myArray, mySlice)

	// var x = [...]int{1, 2, 3}
	// var y = [3]int{1, 2, 3}
	// fmt.Println(x == y) // prints true
	// fmt.Println(x != y) // prints false

	// var x [2][3]int
	// x[0] = 10
	// fmt.Println(x)

	// var scores = [4]int{10, 20, 30, 40}
	// fmt.Println(scores[3])

	// var scores = [4]int{10, 20, 30, 40}
	// i := 3
	// fmt.Println(scores[i])

	// var x [2][3]int
	// x[0][0] = 10
	// fmt.Println(x)

	// var y [10][3]int
	// y[7][0] = 100 // This adds the 100 number to the 0 index of the.
	// fmt.Println(y)

	// var x = [3]int{100, 200, 300}
	// x[1] = 199
	// fmt.Println(x)

	// var inventory = [...]int{2: 50, 100}
	// fmt.Println(inventory)

	// Part 1: The Blueprint
	// var inventory = [...]int{2: 50, 100}

	// // Part 2: The Grid
	// var grid [3][4]int
	// grid[1][2] = inventory[3]

	// // Part 3: The Target
	// const targetIndex = 3
	// fmt.Println(grid[0][targetIndex])

	// var config = [...]int{1: 10, 4: 40}
	// fmt.Println(config)

	// var center [2][5]int
	// fmt.Println(center)

	// Stage 1: The Config Array
	var config = [...]int{1: 10, 4: 40}

	// Stage 2: The Multi-Grid
	var dataCenter [2][5]int
	dataCenter[0][3] = config[2]
	dataCenter[1][1] = config[4]

	// Stage 3: The Access Phase
	readIndex := 5
	fmt.Println(dataCenter[0][readIndex]) /// This will cause a runtime error.

}
