package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	jsonparser "github.com/vzx7/go-head-first/pkg/json_parser"
	"github.com/vzx7/go-head-first/pkg/reflection"
	S "github.com/vzx7/go-head-first/pkg/sort"
)

var data = []S.S2{
	{F1: 1, F2: "One", F3: S.S1{F1: 1, F2: "S1_1", F3: 10}},
	{F1: 2, F2: "Two", F3: S.S1{F1: 1, F2: "S1_1", F3: 20}},
	{F1: -1, F2: "Two", F3: S.S1{F1: -1, F2: "S1_1", F3: -20}},
}

func main() {
	reflection.DoReflect()
	reflection.DoModifyByReflect()
	doSort(data)
	jsonParse()
}

func doSort(data []S.S2) {
	fmt.Println("before:", data)
	sort.Sort(S.S2slice(data))
	fmt.Println("after:", data)
	sort.Sort(sort.Reverse(S.S2slice(data)))
	fmt.Println("reverse:", data)
}

func jsonParse() {
	JSONData := jsonparser.JSONRecord
	if len(os.Args) == 1 {
		fmt.Println("*** Using default JSON record.")
	} else {
		JSONData = os.Args[1]
	}

	JSONMap := make(map[string]any)
	err := json.Unmarshal([]byte(JSONData), &JSONMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonparser.ExploreMap(JSONMap)
	jsonparser.TypeSwitch(JSONMap)
}
