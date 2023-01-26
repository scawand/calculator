package main

/*
import "github.com/scawand/demo/module"


type Animal interface {
	Noise() string
	NumberOfLeg() int
}
func PrintAnimalInfo(a Animal) {
	println("Le bruit est ", a.Noise(), " et il a ", a.NumberOfLeg(), " pattes")

}
func (c *Cat) Noise() string    { return "miaou" }
func (c *Cat) NumberOfLeg() int { return 4 }

func (s *Spider) Noise() string    { return "tt" }
func (s *Spider) NumberOfLeg() int { return 8 }
type Cat struct {
	Name  string
	Breed string
}

type Spider struct {
	Name     string
	Breed    string
	Venomous bool
}






cat := Cat{
	Name:  "kitty",
	Breed: "Siamoi",
}
PrintAnimalInfo(&cat)

func CalculateValue(intChan chan int) {
	randonNumber := module.GenerateRandomNumber(50)
	intChan <- randonNumber
}
func CalculateValue2(intd int) int {
	randonNumber := module.GenerateRandomNumber(50)
	return randonNumber
}

func main() {
	/*foo := make(chan int)

	defer close(foo)
	go CalculateValue(foo)
	num := module.GenerateRandomNumber(50)
	println(num)

	/*
		cat := Cat{
			Name:  "kitty",
			Breed: "Siamoi",
		}
		PrintAnimalInfo(&cat)

		spider := Spider{
			Name:     "spi",
			Breed:    "Tarentule",
			Venomous: true,
		}
		PrintAnimalInfo(&spider)

}

*/
