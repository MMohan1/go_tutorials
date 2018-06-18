package main

import ("fmt";
	"strings"
)

func main () {
	var name string = "my name is manmohan"
	fmt.Println(strings.Contains(name, "manmohan"))
	fmt.Println(strings.Count(name, "manmohan"))
	fmt.Println(strings.HasPrefix(name, "my name"))
	fmt.Println(strings.HasSuffix(name, "mohan"))
	fmt.Println(strings.Index(name, "mohan"))
	fmt.Println(strings.Join([]string{"a","b"}, "-"))
	fmt.Println(strings.Replace(name, "m", "aaaaaa", 18))

}
