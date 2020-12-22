package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func init() {
	Days[21] = Day21
}

type Food struct {
	Allergens   map[string]int
	Ingredients map[string]int
}

func (f *Food) PrettyPrint() {
	fmt.Print("Ingredients: ")
	for k, _ := range f.Ingredients {
		fmt.Printf(" %s ", k)
	}
	fmt.Print("\nAllergens:")
	for k, _ := range f.Allergens {
		fmt.Printf(" %s ", k)
	}

	fmt.Println("\n--------------------------")

}

type Ingredients map[string]int
type AllergenList map[string][]string

func stringArrToMap(strArr []string) map[string]int {
	m := map[string]int{}
	for _, item := range strArr {
		m[strings.TrimSpace(item)] = 1
	}

	return m
}

func NewFood(allergens, ingredients []string) *Food {
	f := Food{
		Allergens:   stringArrToMap(allergens),
		Ingredients: stringArrToMap(ingredients),
	}

	return &f
}

type Foods []Food

func readIngredientList(path string) (Foods, Ingredients) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var foods Foods
	ingredientList := Ingredients{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "(contains", "_")
		line = strings.ReplaceAll(line, ")", "")
		line = strings.ReplaceAll(line, ",", "")

		items := strings.Split(line, " ")
		tpe := "I"
		var ingredients []string
		var allergens []string
		for _, item := range items {
			if item == "_" {
				tpe = "A"
				continue
			}

			switch tpe {
			case "A":
				allergens = append(allergens, item)
			case "I":
				ingredients = append(ingredients, item)
			}
		}

		for _, i := range ingredients {
			ingredientList[i] = 0
		}
		foods = append(foods, *NewFood(allergens, ingredients))
	}

	return foods, ingredientList
}

func getKeys(M map[string]int) []string {
	keys := make([]string, len(M))

	idx := 0

	for key, _ := range M {
		keys[idx] = key
		idx++
	}

	return keys
}

func containsString(str []string, target string) bool {
	//fmt.Println("target: ", target, " ||arr: ", str)
	for _, s := range str {

		if s == target {
			return true
		}
	}
	return false
}

func identifyPossibleAllergens(foods Foods) AllergenList {
	list := AllergenList{}
	for _, food := range foods {
		for al, _ := range food.Allergens {
			keys := getKeys(food.Ingredients)
			if val, ok := list[al]; ok {
				var items []string

				if len(val) == 0 {
					list[al] = keys
					continue
				}

				for _, key := range keys {
					if containsString(list[al], key) {
						items = append(items, key)
					}

				}
				list[al] = items
			} else {
				list[al] = keys
			}

		}
	}
	return list
}

func flagIngredients(inList Ingredients, list AllergenList) {
	for _, v := range list {
		for _, item := range v {
			inList[item] = 1
		}
	}
}

func appearsInFood(list []string, foods Foods) int {
	count := 0

	for _, item := range list {
		for _, food := range foods {
			if _, ok := food.Ingredients[item]; ok {
				count++
			}
		}
	}

	return count
}

func identifyAllergens(list AllergenList) {
	multipleCount := 0
	var singles []string

	// initial count
	for _, ls := range list {
		if len(ls) > 1 {
			multipleCount++
		} else {
			singles = append(singles, ls[0])
		}
	}

	// keep pruning the list's arrays until they
	// contains only 1 value
	for multipleCount > 0 {
		for _, single := range singles {
			for k, val := range list {
				if len(val) > 1 {
					var arr []string
					for _, i := range val {
						if i != single {
							arr = append(arr, i)
						}
					}
					list[k] = arr
				}
			}
		}
		multipleCount = 0
		singles = make([]string, 0)
		for _, ls := range list {
			if len(ls) > 1 {
				multipleCount++
			} else {
				singles = append(singles, ls[0])
			}
		}
	}
}

func getDangerousIngredients(list AllergenList) string {
	var danger []string
	var keys []string

	for k, _ := range list {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		danger = append(danger, list[key][0])
	}

	return strings.Join(danger, ",")
}

func Day21() {
	inputs := []string{"test01", "puzzle"} //
	for _, f := range inputs {
		path := fmt.Sprintf("input/day21/%s.input", f)

		foods, inList := readIngredientList(path)

		list := identifyPossibleAllergens(foods)
		flagIngredients(inList, list)

		var cannotBeAllergens []string
		for k, v := range inList {
			if v == 0 {
				cannotBeAllergens = append(cannotBeAllergens, k)
			}
		}
		identifyAllergens(list)

		ansP1 := appearsInFood(cannotBeAllergens, foods)
		ansP2 := getDangerousIngredients(list)

		fmt.Printf("%s: part 1: %d | part 2: %s\n", f, ansP1, ansP2)
	}
}
