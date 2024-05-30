package main

import "fmt"

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

}

//Structs

//Composition or trait
type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
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
