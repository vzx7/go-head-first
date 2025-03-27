package jsonparser

import "fmt"

var JSONRecord = `{
	"Flag": true,
	"Array": ["a", "b", "c"],
	"Entity": {
		"a1": "b1",
		"a2": "b2",
		"Value": -456,
		"Null": null
	},
	"Message": "Hello Go!"
}`

func TypeSwitch(m map[string]any) {
	for k, v := range m {
		switch c := v.(type) {
		case string:
			fmt.Println("Is a string!", k, c)
		case float64:
			fmt.Println("Is a float64!", k, c)
		case bool:
			fmt.Println("Is a boolean!", k, c)
		case map[string]any:
			fmt.Println("is a map!", k, c)
			TypeSwitch(v.(map[string]any))
		default:
			fmt.Printf("...Is %v: %T!\n", k, c)
		}
		return
	}
}

func ExploreMap(m map[string]any) {
	for k, v := range m {
		embMap, ok := v.(map[string]any)
		if ok {
			fmt.Printf("{\"%v\"}: \n", k)
			ExploreMap(embMap)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v: %v\n", k, v)
		}
	}
}
