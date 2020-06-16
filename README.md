# Copier

  I am a copit, I copy everything from one to another


## Features

* Copy from field to field with same name
* Copy from field where tag assigned
* Copy from slice to slice
* Copy from struct to slice

## Usage

```go
package main

import (
	"fmt"
	"github.com/pkgng/copit"
)

type Human struct {
	Name string
	Role string
	Age  int32
}

type Farmer struct {
	Name      string
	Age       int32
	DoubleAge int32
	EmployeId int64
	SuperRule string
}

func main() {
	var (
		man      = Human{Name: "Zhangsan", Age: 18, Role: "Admin"}
		mans     = []Human{{Name: "Zhangsan", Age: 18, Role: "Admin"}, {Name: "zhangsan 2", Age: 30, Role: "Dev"}}
		farmer  = Farmer{}
		farmers = []Farmer{}
	)

	copit.Copy(&farmer, &man)

	fmt.Printf("%#v \n", farmer)
	// Farmer{
	//    Name: "Zhangsan",           // Copy from field
	//    Age: 18,                  // Copy from field
	//    DoubleAge: 36,            // Copy from method
	//    FarmerId: 0,            // Ignored
	//    SuperRule: "Super Admin", // Copy to method
	// }

	// Copy struct to slice
	copit.Copy(&farmers, &man)

	fmt.Printf("%#v \n", farmers)
	// []Farmer{
	//   {Name: "Zhangsan", Age: 18, DoubleAge: 36, EmployeId: 0, SuperRule: "Super Admin"}
	// }

	// Copy slice to slice
	farmers = []Farmer{}
	copit.Copy(&farmers, &mans)

	fmt.Printf("%#v \n", farmers)
	// []Farmer{
	//   {Name: "Zhangsan", Age: 18, DoubleAge: 36, EmployeId: 0, SuperRule: "Super Admin"},
	//   {Name: "zhangsan 2", Age: 30, DoubleAge: 60, EmployeId: 0, SuperRule: "Super Dev"},
	// }
}
```

## Contributing

You can help to make the project better, check out [http://gorm.io/contribute.html](http://gorm.io/contribute.html) for things you can do.

# Author

**pkgng**

* <http://github.com/pkgng>


## License

Released under the [MIT License](https://github.com/pkgng/copit/blob/master/License).
