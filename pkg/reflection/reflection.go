package reflection

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

type T struct {
	F1 int
	F2 string
	F3 float64
}

func DoModifyByReflect() {
	A := T{1, "F1", 3.0}
	fmt.Println("A:", A)
	r := reflect.ValueOf(&A).Elem()

	fmt.Println("String value:", r.String())
	typeOfA := r.Type()
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		tOfA := typeOfA.Field(i).Name
		fmt.Printf("\t%d: %s %s = %v\n", i, tOfA, f.Type(), f.Interface())

		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k == reflect.Int {
			r.Field(i).SetInt(-100)
		} else if k == reflect.String {
			r.Field(i).SetString("Changed!")
		} else if k == reflect.Float64 {
			r.Field(i).SetFloat(-1.00)
		}
	}
	fmt.Println("A:", A)
}

func DoReflect() {
	A := Record{"String value", -12.123, Secret{"Mihalis", "Tsoucalos"}}
	r := reflect.ValueOf(A)
	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType)
	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("\t%s", iType.Field(i).Name)
		fmt.Printf("\twith type: %s ", r.Field(i).Type())
		fmt.Printf("\tand value _%v \n", r.Field(i).Interface())
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type())
		}
		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type())
		}
	}
}
