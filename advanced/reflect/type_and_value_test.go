package reflect

import (
	"fmt"
	"testing"
)

func TestCreateInsert(t *testing.T) {
	o := Order{
		Id:     60,
		CustId: 120,
	}
	e := Employee{
		Id:      120,
		Name:    "张三",
		Address: "广东省深圳市福田区",
		Salary:  "10000",
	}
	fmt.Println(CreateInsert(&o))
	fmt.Println(CreateInsert(e))
}

func TestCall(t *testing.T) {
	Invoke(TestCreateInsert, t)
}
