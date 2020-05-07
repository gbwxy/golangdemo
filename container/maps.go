package container

import "fmt"

func main() {

	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)
	fmt.Println(len(m))

	m2 := make(map[string]int) //m2 = empty map
	var m3 map[string]int      //m3 = nil

	fmt.Println(m2)
	fmt.Println(m3)

	fmt.Println("key and value.")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("only value.")
	for _, v := range m {
		fmt.Println(v)
	}

	fmt.Println("only key.")
	for k, _ := range m {
		fmt.Println(k)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println("courseName = ", courseName)
	causeName := m["cause"]
	fmt.Println("causeName = ", causeName)

	if causeName, ok := m["cause"]; ok {
		fmt.Println("causeName = ", causeName)
	} else {
		fmt.Println("key 【causeName】 does not exists")
	}

	fmt.Println("Deleting element.")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

	fmt.Println(len(m))
}
