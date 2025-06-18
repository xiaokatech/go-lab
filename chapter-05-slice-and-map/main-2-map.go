package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 24
	ages["Mary"] = 21

	ages_v2 := map[string]int{
		"John": 37,
		"Mary": 21,
	}
	fmt.Println("ages_v2", ages_v2)
	fmt.Println("len(ages_v2)", len(ages_v2))

	hits_v1 := make(map[string]map[string]int)
	fmt.Println("hits_v1", hits_v1)
	add_v1(hits_v1, "/doc/", "au")
	add_v1(hits_v1, "/doc/", "en")
	add_v1(hits_v1, "/doc_v2/", "zh")
	fmt.Println("hits_v1", hits_v1)

	hits_v2 := make(map[Key]int)
	fmt.Println("hits_v2", hits_v2)
	add_v2(hits_v2, Key{"/doc/", "au"})
	add_v2(hits_v2, Key{"/doc/", "en"})
	add_v2(hits_v2, Key{"/doc_v2/", "zh"})
	fmt.Println("hits_v2", hits_v2)
}

func add_v1(m map[string]map[string]int, path, country string) {
	mm, ok := m[path]
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}

type Key struct {
	Path, Country string
}

func add_v2(m map[Key]int, keys Key) {
	mm, ok := m[keys]
	if !ok {
		mm = 0
		m[keys] = mm
	}
	mm++
}
