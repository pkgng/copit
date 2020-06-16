package copit

import (
	"reflect"
	"testing"
	"time"

	"github.com/pkgng/copit"
)

type Human struct {
	Name     string
	Birthday *time.Time
	Nickname string
	Role     string
	Age      int32
	FakeAge  *int32
	Notes    []string
	flags    []byte
}

func (man Human) DoubleAge() int32 {
	return 2 * man.Age
}

type Farmer struct {
	Name      string `copit:"Name"`
	Birthday  *time.Time
	Nickname  *string `copit:"Nickname"`
	Age       int64
	FakeAge   int
	EmployeID int64
	DoubleAge int32
	SuperRule string
	Notes     []string
	flags     []byte
}

func (farmer *Farmer) Role(role string) {
	farmer.SuperRule = "Super " + role
}

func checkFarmer(farmer Farmer, man Human, t *testing.T, testCase string) {
	if farmer.Name != man.Name {
		t.Errorf("%v: Name haven't been copied correctly.", testCase)
	}
	if farmer.Nickname == nil || *farmer.Nickname != man.Nickname {
		t.Errorf("%v: NickName haven't been copied correctly.", testCase)
	}
	if farmer.Birthday == nil && man.Birthday != nil {
		t.Errorf("%v: Birthday haven't been copied correctly.", testCase)
	}
	if farmer.Birthday != nil && man.Birthday == nil {
		t.Errorf("%v: Birthday haven't been copied correctly.", testCase)
	}
	if farmer.Birthday != nil && man.Birthday != nil &&
		!farmer.Birthday.Equal(*(man.Birthday)) {
		t.Errorf("%v: Birthday haven't been copied correctly.", testCase)
	}
	if farmer.Age != int64(man.Age) {
		t.Errorf("%v: Age haven't been copied correctly.", testCase)
	}
	if man.FakeAge != nil && farmer.FakeAge != int(*man.FakeAge) {
		t.Errorf("%v: FakeAge haven't been copied correctly.", testCase)
	}
	if !reflect.DeepEqual(farmer.Notes, man.Notes) {
		t.Errorf("%v: Copy from slice doen't work", testCase)
	}
}

func TestCopySameStructWithPointerField2(t *testing.T) {
	var fakeAge int32 = 12
	var currentTime time.Time = time.Now()
	man := &Human{Birthday: &currentTime, Name: "Zhangsan", Nickname: "zhangsan", Age: 18, FakeAge: &fakeAge, Role: "Admin", Notes: []string{"hello world", "welcome"}, flags: []byte{'x'}}
	newHuman := &Human{}
	copit.Copy(newHuman, man)
	if man.Birthday == newHuman.Birthday {
		t.Errorf("TestCopySameStructWithPointerField: copy Birthday failed since they need to have different address")
	}

	if man.FakeAge == newHuman.FakeAge {
		t.Errorf("TestCopySameStructWithPointerField: copy FakeAge failed since they need to have different address")
	}
}

func checkFarmer2(farmer Farmer, man *Human, t *testing.T, testCase string) {
	if man == nil {
		if farmer.Name != "" || farmer.Nickname != nil || farmer.Birthday != nil || farmer.Age != 0 ||
			farmer.DoubleAge != 0 || farmer.FakeAge != 0 || farmer.SuperRule != "" || farmer.Notes != nil {
			t.Errorf("%v : farmer should be empty", testCase)
		}
		return
	}

	checkFarmer(farmer, *man, t, testCase)
}

func TestCopyStruct2(t *testing.T) {
	var fakeAge int32 = 12
	man := Human{Name: "Zhangsan", Nickname: "zhangsan", Age: 18, FakeAge: &fakeAge, Role: "Admin", Notes: []string{"hello world", "welcome"}, flags: []byte{'x'}}
	farmer := Farmer{}

	if err := copit.Copy(farmer, &man); err == nil {
		t.Errorf("Copy to unaddressable value should get error")
	}

	copit.Copy(&farmer, &man)
	checkFarmer(farmer, man, t, "Copy From Ptr To Ptr")

	farmer2 := Farmer{}
	copit.Copy(&farmer2, man)
	checkFarmer(farmer2, man, t, "Copy From Struct To Ptr")

	farmer3 := Farmer{}
	ptrToHuman := &man
	copit.Copy(&farmer3, &ptrToHuman)
	checkFarmer(farmer3, man, t, "Copy From Double Ptr To Ptr")

	farmer4 := &Farmer{}
	copit.Copy(&farmer4, man)
	checkFarmer(*farmer4, man, t, "Copy From Ptr To Double Ptr")
}
