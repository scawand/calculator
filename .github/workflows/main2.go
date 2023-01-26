package main

/*
func main() {

	/*
		anonyme := func() string {
			return "Mistère"
		}() // parenthe
		println("Début")
		println(anonyme)
		println("Fin")
*/
//testFor()
/*
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())
	if x := rand.Int(); x%2 == 0 {
		fmt.Println(x, "est paire")
	} else {
		fmt.Println(x, "est impair")
	}
*/
/*employees := []Employee{
	Employee{
		Person: Person{
			LastName: "Doe", FirstName: "John",
		},
	},
	Employee{
		Person: Person{
			LastName: "Campbell", FirstName: "David",
		},
	},
}

data, _ := json.Marshal(employees)
fmt.Printf("%s\n", data)

var decoded []Employee
json.Unmarshal(data, &decoded)
fmt.Printf("%v", decoded)
*/
/*
		count := 8
		seqFibonnacci := []int{1, 1}
		for nb := 2; nb < count; nb++ {
			seqFibonnacci = append(seqFibonnacci, seqFibonnacci[nb-2]+seqFibonnacci[nb-1])
		}
		fmt.Println(seqFibonnacci)

	romainNomber := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"v": 5,
		"I": 1,
	}
	fmt.Println(romainNomber["M"])

		var num int

		fmt.Print("What's the Fibonacci sequence you want? ")
		fmt.Scanln(&num)
		fmt.Println("The Fibonacci sequence is:", fibonacci(num))

	fmt.Println("MCLX is", romanToArabic("MCLX"))
	fmt.Println("MCMXCIX is ", romanToArabic("MCMXCIX"))
	fmt.Println("MCMZ is", romanToArabic("MCMZ"))
*/
//fmt.Println(monNom(""))
//fmt.Println(monNom("Seb"))
//packagee()
/*
	glucides100g := map[string]float32{
		"Nutella":   33.6,
		"cola zero": 0.01,
		"avocat":    1.1,
	}
	for key, value := range glucides100g {
		fmt.Println(key, value)
	}
*/
// age := 2 // variable par ref
// age = passValeur(age)
// println(age)
/*
}
/*
func passValeur(number int) (num int) {
	number = 5
	num = number
	return
}

func val(age int) {
	age = 18
}


func testFor() {


		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}


//équivalent du while

		nb := 0
		for nb < 5 {
			fmt.Println(nb)
			x++
		}


nb := 0
for {
	if nb > 5 {
		break
	}
	fmt.Println(nb)
	nb++

}

// continue = revenir directement au début de la boucle
count := 0
for ; count <= 10; count++ {
	if count%2 == 1 {
		continue
	}
	fmt.Println(count)
}

//var uneListe [n]type

		var ListeInt [4]int

		ListeInt[0] = 5
		ListeInt[1] = 10
		ListeInt[2] = 15

		fmt.Println(ListeInt)
		fmt.Println(ListeInt[0])
		fmt.Println(ListeInt[1])
		fmt.Println(ListeInt[2])

		listIntAuto := [...]int{40, 50} // compter nb élément
		fmt.Println(listIntAuto)
		fmt.Println(listIntAuto[0])
		fmt.Println(listIntAuto[1])

bigListInt := [...]int{40, 50, 33, 332, 462, 133, 555}

//boucle lire tableau
for pos, value := range bigListInt {
	fmt.Printf("Position %d a la valeur %d\n", pos, value)
}

	// si 2 retour on doit spécifier les 2 retours

}

//Je suis une fonction Public
func MonNom(nom string) (string, error) {
	if nom == "" {
		return "Oups", errors.New("Il y a pas de nom")
	}
	message := "Bonjour à toi " + nom
	return message, nil
}
//Je suis une fonction private
func monAge(age int) (int, error) {
	if age <0  {
		return 0, errors.New("age non valide")
	}

	return age, nil
}

/*
func fibonacci(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}

	nums := make([]int, n)
	nums[0], nums[1] = 1, 1

for InitSimplleStatement ; Condition ; UpdateStatement{

}

	for i := 2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}

	return nums
}
*/
/*
func romanToArabic(numeral string) int {
	romanMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	arabicVals := make([]int, len(numeral)+1)

	for index, digit := range numeral {
		if val, present := romanMap[digit]; present {
			arabicVals[index] = val
		} else {
			fmt.Printf("Error: The roman numeral %s has a bad digit: %c\n", numeral, digit)
			return 0
		}
	}

	total := 0

	for index := 0; index < len(numeral); index++ {
		if arabicVals[index] < arabicVals[index+1] {
			arabicVals[index] = -arabicVals[index]
		}
		total += arabicVals[index]
	}

	return total
}
*/

/*var nom, prenom string
var age int

var (
	article, description string
	prix                 float32
)

var (
	unArticle      string  = "Nutella"
	uneDescription string  = "Pot qui se vide rapidement"
	unPrix         float32 = 4.99
)
*/ /*
func lesVariables() {
	var (
		article, description string
		prix                 float32
	)

	var (
		unArticle      string  = "Nutella"
		uneDescription string  = "Pot qui se vide rapidement"
		unPrix         float32 = 4.99
	)
	var nom, prenom string
	var age int
	println(nom, prenom, age)

	monNom, monSurnom := "Sébastien", "Seb"
	monId := 1234
	println(monNom, monSurnom, monId)

}
*/
