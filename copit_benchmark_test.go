package copit_test

import (
	"encoding/json"
	"testing"

	"github.com/pkgng/copit"
)

func BenchmarkCopyStruct(b *testing.B) {
	var fakeAge int32 = 12
	user := User{Name: "Zhangsan", Nickname: "zhangsan", Age: 18, FakeAge: &fakeAge, Role: "Admin", Notes: []string{"hello world", "welcome"}, flags: []byte{'x'}}
	for x := 0; x < b.N; x++ {
		copit.Copy(&Employee{}, &user)
	}
}

func BenchmarkNamaCopy(b *testing.B) {
	var fakeAge int32 = 12
	user := User{Name: "Zhangsan", Nickname: "zhangsan", Age: 18, FakeAge: &fakeAge, Role: "Admin", Notes: []string{"hello world", "welcome"}, flags: []byte{'x'}}
	for x := 0; x < b.N; x++ {
		employee := &Employee{
			Name:     user.Name,
			Nickname: &user.Nickname,
			Age:      int64(user.Age),
			FakeAge:  int(*user.FakeAge),
			Notes:    user.Notes,
		}
		employee.Role(user.Role)
	}
}

func BenchmarkJsonMarshalCopy(b *testing.B) {
	var fakeAge int32 = 12
	user := User{Name: "Zhangsan", Nickname: "zhangsan", Age: 18, FakeAge: &fakeAge, Role: "Admin", Notes: []string{"hello world", "welcome"}, flags: []byte{'x'}}
	for x := 0; x < b.N; x++ {
		data, _ := json.Marshal(user)
		var employee Employee
		json.Unmarshal(data, &employee)

		employee.Role(user.Role)
	}
}
