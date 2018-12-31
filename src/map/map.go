package main

import "fmt"

type PersonInfo struct {
	ID string
	Name string
	Address string
}

func main() {
	var personDB map[string] PersonInfo // 
	personDB = make(map[string] PersonInfo)
	personDB["1"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDB["123"] = PersonInfo{"1", "Jack", "Room 101,..."}
	personDB["1"] = PersonInfo{"12", "Tom111", "Room 203,..."}
	// delete(personDB, "1")
	// fmt.Println(res)
	// 从这个map查找键为"1234"的信息 
	person, ok := personDB["1"] // ok是一个返回的bool型，返回true表示找到了对应的数据 
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.", person.Address)
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}
}