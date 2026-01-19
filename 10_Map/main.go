package main

import (
	"fmt"
	"maps"
	// "slices"
)

func main(){
	var mp = make(map[string]string)
	mp["name"] = "Ankesh"
	mp["dept"] = "IT"
	fmt.Println(mp)

	// another creation technique
	mp1 := map[string]int{}
	mp1["name"] = 1
	mp1["age"] = 2
	fmt.Println(mp1)

	k , ok := mp1["name"]
	if ok {
		fmt.Printf("All ok and value is:%d\n",k)
	} else {
		fmt.Println("Not Ok")
	}
	mp2 := map[string]int{"name":2,"age":3}
	fmt.Println(maps.Equal(mp1,mp2))
}