package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type food struct {
	ingredients []string
	allergen    []string
}

type dangerFood struct {
	ingredients string
	allergen    string
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	foods := []food{}
	for s.Scan() {
		l := s.Text()
		ls := strings.Split(l, "(")
		tmpb := []byte(ls[0])
		ings := strings.Split(string(tmpb[0:len(tmpb)-1]), " ")

		tmpb = []byte(ls[1])
		algs := strings.Split(string(tmpb[9:len(tmpb)-1]), ", ")

		foods = append(foods, food{
			ingredients: ings,
			allergen:    algs,
		})
	}

	allergeneList := make(map[string][]string)
	for _, f := range foods {
		for _, alg := range f.allergen {
			if ings, ok := allergeneList[alg]; ok {
				// remove ingredients if don't exist in both lists
				allergeneList[alg] = []string{}
				for _, ingNew := range f.ingredients {
					for _, ingOld := range ings {
						if ingNew == ingOld {
							allergeneList[alg] = append(allergeneList[alg], ingNew)
						}
					}
				}
			} else {
				allergeneList[alg] = f.ingredients
			}
		}
	}
	for alg, ings := range allergeneList {
		if len(ings) == 1 {
			for algComp, ingsComp := range allergeneList {
				if alg != algComp {
					allergeneList[algComp] = []string{}
					for _, ingComp := range ingsComp {
						if ingComp != ings[0] {
							allergeneList[algComp] = append(allergeneList[algComp], ingComp)
						}
					}
				}
			}
		}
	}

	result := 0
	for _, f := range foods {
		for _, ing := range f.ingredients {
			hasAlg := false
			for _, ingsComp := range allergeneList {
				if ing == ingsComp[0] {
					hasAlg = true
					break
				}
			}
			if !hasAlg {
				result++
			}
		}
	}
	fmt.Println(result)

	// puzzle two
	dangerousIngredientList := []dangerFood{}
	for alg, ing := range allergeneList {
		dangerousIngredientList = append(dangerousIngredientList, dangerFood{
			ingredients: ing[0],
			allergen:    alg,
		})
	}
	// sort.Strings(dangerousIngredientList)

	sort.SliceStable(dangerousIngredientList, func(i, j int) bool {
		return dangerousIngredientList[i].allergen < dangerousIngredientList[j].allergen
	})
	for i, di := range dangerousIngredientList {
		if i == 0 {
			fmt.Printf("%s", di.ingredients)
		} else {
			fmt.Printf(",%s", di.ingredients)
		}
	}
	fmt.Printf("\n")
}
