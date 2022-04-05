package structs

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEmployee(t *testing.T) {
	employee := Employee{
		ID:        1,
		FirstName: "Li",
		LastName:  "Si",
		Address:   "ShenZhen",
	}
	fmt.Printf("%v", employee)
}

func TestJson(t *testing.T) {
	employee := Employee{
		ID:        1,
		FirstName: "Li",
		LastName:  "Si",
		Address:   "ShenZhen",
	}
	employees := []Employee{
		employee,
		{
			ID:        2,
			FirstName: "ZhangSan",
			LastName:  "Yin",
			Address:   "ChangSha",
		},
		{
			ID:        3,
			FirstName: "Wang",
			LastName:  "Wu",
			Address:   "JiuJiang",
		},
		{
			ID:        4,
			FirstName: "Ma",
			LastName:  "Zi",
			Address:   "JiuJiang",
		},
	}
	encoded, _ := json.Marshal(employees)
	fmt.Printf("%s\n", encoded)
	var decoded []Employee
	json.Unmarshal(encoded, &decoded)
	fmt.Println(decoded)
}

func TestPointer(t *testing.T) {
	employee := Employee{
		ID:        1,
		FirstName: "Li",
		LastName:  "Si",
		Address:   "ShenZhen",
	}
	copy := employee // 复制
	employeePointer := &employee
	employeePointerCopy := employeePointer // 指针副本

	copy.ID = 2 // 修改副本不会影响原对象
	fmt.Printf("%v\n%v\n%v\n%v\n", employee, *employeePointer, *employeePointerCopy, copy)
	fmt.Println()

	employeePointer.ID = 3 // 修改指针会影响原对象
	fmt.Printf("%v\n%v\n%v\n%v\n", employee, *employeePointer, *employeePointerCopy, copy)
	fmt.Println()

	employeePointerCopy.ID = 4 // 修改指针副本会影响原对象
	fmt.Printf("%v\n%v\n%v\n%v\n", employee, *employeePointer, *employeePointerCopy, copy)
	fmt.Println()
}
