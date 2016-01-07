package inventory

type Inventory struct {
	products []Product
}

func (i *Inventory) AddProduct(p Product) {
	i.products = append(i.products, p)
}

func (i *Inventory) Get() []Product {
	return i.products
}