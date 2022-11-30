package day21

import (
	"fmt"
	"sort"
	"strings"
)

type food struct {
	ingredients []string
	allergens   []string
}

func ParseInput(input []string) []food {
	foods := make([]food, 0, len(input))
	for _, line := range input {
		var f food
		f.ingredients = make([]string, 0)
		f.allergens = make([]string, 0)
		inAllergens := false

		for _, part := range strings.Split(line, " ") {
			if part == "(contains" {
				inAllergens = true
				continue
			}

			if inAllergens {
				f.allergens = append(f.allergens, part[:len(part)-1])
			} else {
				f.ingredients = append(f.ingredients, part)
			}
		}

		foods = append(foods, f)
	}
	return foods
}

var possibilities struct {
	allergen    string
	ingredients []string
}

func Intersection[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}

type empty interface{}

var isEmpty empty

func IterateAllergens(pos map[string][]string, knownAllergen *map[string]string) {
	nextIngredients := make(map[string][]string)
	for allergen, possibleIngredients := range pos {
		pis := make([]string, 0)

		for _, ing := range possibleIngredients {
			found := false
			for _, allergenIngredient := range *knownAllergen {
				if ing == allergenIngredient {
					found = true
					break
				}
			}

			if !found {
				pis = append(pis, ing)
			}
		}

		if len(pis) == 1 {
			(*knownAllergen)[allergen] = pis[0]
		} else {
			nextIngredients[allergen] = pis
		}
	}

	for allergen, _ := range *knownAllergen {
		delete(nextIngredients, allergen)
	}

	if len(nextIngredients) > 0 {
		IterateAllergens(nextIngredients, knownAllergen)
	}
}
func Day21(input []string) []string {
	foods := ParseInput(input)
	possibilities := make(map[string][]string)
	ingredients := make(map[string]interface{})
	for _, f := range foods {
		for _, ingredient := range f.ingredients {
			ingredients[ingredient] = isEmpty
		}
		for _, allergen := range f.allergens {
			if _, ok := possibilities[allergen]; !ok {
				possibilities[allergen] = f.ingredients
				continue
			}

			possibilities[allergen] = Intersection(possibilities[allergen], f.ingredients)
		}
	}

	for _, possibleIngredients := range possibilities {
		for _, pi := range possibleIngredients {
			delete(ingredients, pi)
		}
	}

	found := 0
	for _, f := range foods {
		for _, ing := range f.ingredients {
			if _, ok := ingredients[ing]; ok {
				found++
			}
		}
	}
	knownAllergens := make(map[string]string)
	IterateAllergens(possibilities, &knownAllergens)
	allergenList := make([]string, 0)
	for allergen, _ := range knownAllergens {
		allergenList = append(allergenList, allergen)
	}
	sort.Strings(allergenList)

	result := make([]string, len(allergenList))

	for i, allergen := range allergenList {
		result[i] = knownAllergens[allergen]
	}

	return []string{
		fmt.Sprintf("Found %d appearances of ingredients that can't be allergens", found),
		fmt.Sprintf("Allergen-containing ingredients sorted by allergen: %s", strings.Join(result, ",")),
	}
}
