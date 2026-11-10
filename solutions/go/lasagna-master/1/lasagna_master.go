package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2	
	}
	return len(layers) * avgTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0 // Mengubah nama variabel agar lebih selaras

	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" { // PERBAIKAN: dari "sauces" menjadi "sauce"
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsIngredients []string, myIngredients []string) {
	lastFriendIndex := len(friendsIngredients) - 1
	lastMyIndex := len(myIngredients) - 1

	myIngredients[lastMyIndex] = friendsIngredients[lastFriendIndex]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaled := make([]float64, len(amounts))
	factor := float64(portions) / 2.0

	for i, v := range amounts {
		scaled[i] = v * factor
	}
	return scaled
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.