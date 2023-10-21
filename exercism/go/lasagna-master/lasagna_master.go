package lasagnamaster

// PreparationTime gives an estimate of the amount of prep time required based on layers
func PreparationTime(layers []string, avgTimePerLayer int) int {
	if avgTimePerLayer == 0 {
		avgTimePerLayer = 2
	}
	mins := 0
	for range layers {
		mins += avgTimePerLayer
	}
	return mins
}

// Quantites uses a 'naked' return, specifying the amount of noodles & sauce needed
func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0
	for _, value := range layers {
		if value == "noodles" {
			noodles += 50
		} else if value == "sauce" {
			sauce += 0.2
		}

	}
	return
}

// AddSecretIngredient replaces the last item of my list with my friend's list
func AddSecretIngredient(friendIngredients []string, myIngredients []string) {
	myIngredients[len(myIngredients)-1] = friendIngredients[len(friendIngredients)-1]
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	multiplier := float64(portions) / 2
	scaledAmounts := []float64{}
	for _, value := range amounts {
		scaledAmounts = append(scaledAmounts, value*multiplier)
	}
	return scaledAmounts
}
