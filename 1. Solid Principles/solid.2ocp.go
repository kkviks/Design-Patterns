package solid

/* Open-Closed Principle
Types should be open for extension (grab the interface and implement this somewhere in your code)
but closed for modication.

We don't want to intefere with what has already been written and tested.
Open-Closed is all about being open to extension
We want to extend a scenario but by adding additional types/func without modifying already existing code written or test
*/

/* Enterprise Pattern - Specification
Say, you own an online store, you want your uses to sort products based on some criteron
And you have a specification: Our website has to allow the user to be able to filter by some criteron

Initial criteron: Filter by color,
later comes: Filter by size? Same code copy karoge?
.......
More specifications kicks in overtime


*/

import "fmt"

// combination of OCP and Repository demo

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	// Filter on GPU or CPU for example
}

func (f *Filter) filterByColor(
	products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// Violation of Open-Close principle
func (f *Filter) filterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) filterBySizeAndColor(
	products []Product, size Size,
	color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// filterBySize, filterBySizeAndColor

// Specification Interface type is open for extension, meaning you can implement this
// but close to modification, you're very unlikely to modify this Specification interface
// Also you are unlikely to modify the BetterFilter function as well.
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

// Composite
type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) &&
		spec.second.IsSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main_ocp() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	fmt.Print("Green products (old):\n")
	f := Filter{}
	for _, v := range f.filterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}
	// ^^^ BEFORE

	// vvv AFTER
	fmt.Print("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}

	largeGreenSpec := AndSpecification{largeSpec, greenSpec}
	fmt.Print("Large blue items:\n")
	for _, v := range bf.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
