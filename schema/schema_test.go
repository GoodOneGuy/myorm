package schema

import (
	"fmt"
	"github.com/GoodOneGuy/myorm/dialect"
	"testing"
)

type User struct {
	Name string `myorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	s := Parse(&User{}, TestDial)
	fmt.Println(s)
	for _, f := range s.Fields {
		fmt.Println(f)
	}
}
