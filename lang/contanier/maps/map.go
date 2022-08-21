package main

import (
	"fmt"
)

func declare() {
	m := map[string]string{
		"hello": "world",
	}

	s := m["hello"]

	fmt.Printf("s[\"hello\"] =  %s\n", s)

	// 说明
	s = m["world"]
	fmt.Printf("s[\"world\"] = %s\r\n", s)

	m2 := map[string]map[string]string{
		"a": {
			"b": "c",
		},
	}

	fmt.Printf("m2[\"a\"][\"b\"]  = %s\r\n\n", m2["a"]["b"])

	m3 := make(map[string]int) // m2 == empty map

	var m4 map[string]int // m3 == nil

	fmt.Printf("m3 is %s ; m4 is %s \r\n", m3, m4)

}

func mapCurd() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	fmt.Println("Getting Value")
	fmt.Printf("m[\"name\"] = %s\r\n", m["name"])
	fmt.Println()

	fmt.Println("Getting Value")
	if value, ok := m["what"]; ok {
		fmt.Printf("m[\"name\"] = %s\r\n", value)
	} else {
		fmt.Println("key does not exists!")
	}
	fmt.Println()

	mapForeach(m)

	fmt.Println("Adding Value")
	m["winter"] = "fell"
	mapForeach(m)

	fmt.Println("Deleting Value")
	delete(m, "name")
	mapForeach(m)

	fmt.Println("Changing Value")
	m["winter"] = "winterfell"
	mapForeach(m)

}

func mapForeach(m map[string]string) {
	fmt.Println("foreach:")
	for k, v := range m {
		fmt.Printf("key = %s \t value = %s \r\n", k, v)
	}
	fmt.Println()
}

func main() {
	declare()
	mapCurd()
}
