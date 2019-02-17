package main

import "fmt"

type Animal struct {
	category string // 动物学基本分类。
}

func (self Animal) Category() string {
	return self.category
}

type Cat struct {
	name           string // 名字。
	scientificName string // 学名。
	*Animal
}

// 注意如果写成`func (self *Cat) String() string`，值变量就无法被 print 正常打印
func (self Cat) String() string {
	return fmt.Sprintf("%s, %s, %s", self.name, self.scientificName, self.Animal.Category())
}

func NewPointer(name, scientificName, category string) *Cat {
	return &Cat{
		name:           name,
		scientificName: scientificName,
		Animal: &Animal{
			category: category,
		},
	}
}

func New(name, scientificName, category string) Cat {
	return Cat{
		name:           name,
		scientificName: scientificName,
		Animal: &Animal{
			category: category,
		},
	}
}

func main() {
	cat := New("little pig", "American Shorthair", "cat")
	pointerCat := NewPointer("little pig", "American Shorthair", "cat")
	fmt.Println(cat)
	fmt.Println(pointerCat)
}
