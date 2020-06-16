# Copier

  I am a copit, I copy everything from one to another


## Features

* Copy from field to field with same name
* Copy from field to method with same name
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

type DocModel struct {
	ID       int
	Tag      []string
	Text     string
	ModifyAt string
}

type UserModel struct {
	Name    string
	Avatar  string
	Role    string
	Birtday string
}

type Doc struct {
	Owner  string `copit:"Name"`
	Avatar string
	Zodiac string
	DocID  int `copit:"ID"`
	Tag    []string
	Text   string
}

func (d *Doc) Birtday(birth string) {
	d.Zodiac = "天平座"
}

func main() {
	var (
		user     = UserModel{Name: "Zhangsan", Avatar: "http://a.b.c/a.png", Role: "Admin", Birtday: "2001-10-11"}
		docModel = DocModel{ID: 2501, Text: "this is all doc text", ModifyAt: "2020-1-12", Tag: []string{"a", "b", "c"}}
		doc      = Doc{}
	)

	copit.Copy(&doc, &user)
	copit.Copy(&doc, &docModel)

	fmt.Printf("%#v \n", doc)
}
```

## Contributing

You can help to make the project better, check out [http://gorm.io/contribute.html](http://gorm.io/contribute.html) for things you can do.

# Author

**pkgng**

* <http://github.com/pkgng>


## License

Released under the [MIT License](https://github.com/pkgng/copit/blob/master/License).
