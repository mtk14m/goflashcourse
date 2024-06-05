package main

import (
	"fmt"
)

func main() {

	/*
		if len(os.Args) != 2 {
			os.Exit(1)
		}
		fmt.Println("It's over", os.Args[1])
	*/
	/*
		power := 1000
		fmt.Printf("default power is %d\n", power)

		name, power := "Goku", 9000
		fmt.Printf("%s's power is over %d\n", name, power)

		_, exists := Power("goku")

		if exists == false {
			fmt.Println("Ah lala")
		}
	*/

	//struct declaration & initialization
	//goku := Saiyan{}
	//goku := Saiyan{Name: "Goku"}
	//goku.Power = 9000
	//goku := Saiyan{"Goku", 9000}
	/*
		goku := Saiyan{
			Name:  "goku",
			Power: 9000,
		}
	*/

	//goku := Saiyan{"Goku", 9001, nil}
	//goku.Super() // est égale aussi à
	//Super(goku)
	//fmt.Println(goku.Power)

	//New functions
	/*
		gohan := new(Saiyan)
		gohan.Name = "gohan"
		gohan.Power = 9001
	*/
	// wich is the same as
	/*
		gohan:= &Saiyan{}
		gohan:= &Saiyan{
			Name: "gohan",
			Power: 90001,
		}
	*/

	//fields of Structure
	/*
		gohan := &Saiyan{
			Person: &Person{"gohan"},
			Power:  1000,
			Father: &Saiyan{
				Person: &Person{"gohan"},
				Power:  9001,
				Father: nil,
			},
		}

		gohan.Introduce()

		goku := &Saiyan{
			Person: &Person{"Goku"},
		}

		fmt.Println(goku.Name)
		fmt.Println(goku.Person.Name)
	*/

	//Maps, Arrays and Slices

	//Arrays
	/*
		var scores [10]int
		scores[0] = 339
		fmt.Println(len(scores))
		for index, value := range scores {
			fmt.Println(index + value)
		}*/

	//Slice
	//scores := []int{1, 4, 293, 4, 9} first way
	//scores := make([]int, 10) or

	/*
		scores := make([]int, 0, 10)
		//scores = append(scores, 5)
		scores = scores[0:8]
		scores[7] = 524552
		fmt.Println(scores)
	*/

	//4 way to create a slice
	//names:= []string{"Minou", "Leto"}
	//cheks:=make([]bool, 10)
	//var names[]string
	//scores := make([]int, 0, 20)

	//Examples
	/*
		func extarctPowers(Saiyan []*Saiyan) []int {
			powers:= make([]int, len(saiyans))
			for index, saiyan := range saiyans{
				powers[index]=saiyan.Power
			}
			return powers
		}
	*/

	/*
		func extractPowers(saiyans []*Saiyan) []int {
		powers := make([]int, 0, len(saiyans))
		for _, saiyan := range saiyans {
		powers = append(powers, saiyan.Power)
		}
		return powers
		}
	*/
	/*
		scores := []int{1, 2, 3, 4, 5}
		scores = removeAtIndex(scores, 2)
		fmt.Println(scores) // [1 2 5 4]
	*/

	//the copy built in function
	/*
		scores := make([]int, 100)
		for i := 0; i < 100; i++ {
			scores[i] = int(rand.Int31n(1000))
		}
		sort.Ints(scores)
		worst := make([]int, 5)
		copy(worst[2:4], scores[:9])
		fmt.Println(worst)
	*/

	//Maps
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]
	fmt.Println(power, exists)

	//total:= len(lookup)
	//delete(lookup, "goku")
	//lookup:= make(map[string]int, 100)
	//goku.Friends["Krillin"]= ...//to load or create Krillin

}

// won't preserve order
func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	//swap the last value and the value we want to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}

//Structs

type Saiyan2 struct {
	Name    string
	Friends map[string]*Saiyan
}

// Composition or trait
type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

// ==>overwriting this func for Saiyan
func (s *Saiyan) Introduce() {
	fmt.Printf("Hi, I'm %s\n", s.Name)
}

type Saiyan struct {
	*Person
	Power  int
	Father *Saiyan
}

func Super(s *Saiyan) {
	s.Power += 10000
}

func (s *Saiyan) Super() {
	s.Power += 10000
}

//Constructors using factory design pattern

func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Person: &Person{Name: name},
		Power:  power,
	}
}

//Functions

/*
func log(message string) {
}

	func add(a int, b int) int {
		return a + b
	}
	===>
	func add(a, b int) int {
		return a + b
	}

func Power(name string) (int, bool) {
	return 0, false
}
*/
