package reflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateInsert(t *testing.T) {
	tests := []struct {
		name   string
		obj    any
		expect string
	}{
		{"struct",
			Order{Id: 60, CustId: 120},
			`INSERT INTO Order (Id, CustId); VALUES (60, 120)`},
		{
			"pointer",
			&Employee{
				Id:      120,
				Name:    "张三",
				Address: "广东省深圳市福田区",
				Salary:  "10000",
			},
			`INSERT INTO Employee (Id, Name, Address, Salary); VALUES (120, "张三", "广东省深圳市福田区", "10000")`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			insertSQL := CreateInsert(test.obj)
			assert.Equal(t, test.expect, insertSQL)
		})
	}
}

func TestCall(t *testing.T) {
	Invoke(TestCreateInsert, t)
}
