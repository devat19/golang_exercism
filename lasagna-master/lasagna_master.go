package lasagna_master

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int) int {

	var layerPrepTime int = prepTime
	if prepTime == 0 {
		layerPrepTime = 2
	}
	// else   <<<<<<< Check why this doesn't work
	// 	layerPrepTime = prepTime
	// }

	// fmt.Println("Data -layerPrepTime, layerscount", layerPrepTime, prepTime, len(layers))

	return len(layers) * layerPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {

	const noodlesPerLayerInGrams = 50
	const saucePerLayerInLitres = 0.2

	var totalNoodlesNeededInGrams int = 0
	var totalSauceNeededInLitres float64 = 0.0

	for _, layer := range layers {

		if layer == "noodles" {
			totalNoodlesNeededInGrams = totalNoodlesNeededInGrams + noodlesPerLayerInGrams
		} else if layer == "sauce" {
			totalSauceNeededInLitres = totalSauceNeededInLitres + saucePerLayerInLitres
		}
	}

	//  return noodlesPerLayerInGrams * len(layers), float64(saucePerLayerInLitres * saucePerLayerInLitres)

	return totalNoodlesNeededInGrams, totalSauceNeededInLitres
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendsList, myList []string) []string {

	// Was trying to use map to make intersection of two arrays - Retry later! - also , anonymus functions, naked returns?
	// m := make(map[string]int)

	// for index, ingredient := range myList {
	// 	m[ingredient] = index
	// }

	myList[len(myList)-1] = friendsList[len(friendsList)-1]

	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {

	var scaledRecipe = []float64{}

	var scalingFactor float64 = float64(portions) / 2 // Is that typecasting in GO?

	for _, portion := range quantities {
		scaledRecipe = append(scaledRecipe, portion*float64(scalingFactor)) // why is this float?
	}

	// fmt.Println("DATA >> 00", scalingFactor, portions)

	return scaledRecipe

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
