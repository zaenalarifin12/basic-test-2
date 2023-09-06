package main

import (
	"fmt"
	"log"
)

func plusMinus(numbers []int ) {

	if len(numbers) <= 0 || len(numbers) > 100 {
		log.Fatal("your data there is under 0 or upper of 100 number")
	}

	

	var dataPositif, dataNegatif, dataZero []int

	for i := 0; i < len(numbers); i++ {
		
		// check
		if numbers[i] <= -101 || numbers[i] > 100 {
			log.Fatal("your number there is under -100 and upper of 100")
		}


		// positif number
		if numbers[i] > 0 {
			dataPositif = append(dataPositif, numbers[i])
		}

		if numbers[i] == 0 {
			dataZero = append(dataZero, numbers[i])
		}

		if numbers[i] < 0 {
			dataNegatif = append(dataNegatif, numbers[i])
		}
	}

	proportionPositif := float64(float64(len(dataPositif))/float64(len(numbers)))

	proportionNegatife := float64(float64(len(dataNegatif))/float64(len(numbers)))

	proportionZero := float64(float64(len(dataZero))/float64(len(numbers)))

	fmt.Printf("PROPORTION OF POSITIF VALUE %.6f\n", proportionPositif)
	fmt.Printf("PROPORTION OF NEGATIF VALUE %.6f\n", proportionNegatife)
	fmt.Printf("PROPORTION OF ZERO VALUE %.6f\n", proportionZero)

}

func main()  {

	dataNumbers := []int{-4, 3, -9, 0, 4, 5}

	plusMinus(dataNumbers)

}