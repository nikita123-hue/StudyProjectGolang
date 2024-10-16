package main

import "fmt"

type Warehouse interface {
	AddProduct(product Product) error
	UpdateQuantity(quantity int, productID int) error
}

type Product struct {
	ID       int
	Name     string
	Quantity int
}

type Storage struct {
	Products map[int]Product
}

func (st *Storage) AddProduct(product Product) error {
	if st.Products == nil {
		st.Products = make(map[int]Product)
	}
	if _, ok := st.Products[product.ID]; ok {
		return fmt.Errorf("Product with ID %d already exists", product.ID)
	}

	st.Products[product.ID] = product
	return nil
}

func (st *Storage) UpdateQuanity(quantity int, productID int) error {
	product, ok := st.Products[productID]
	if !ok {
		return fmt.Errorf("Product with ID %d does not exist", productID)
	}
	if product.Quantity+quantity < 0 {
		return fmt.Errorf("Product with ID %d has negative quantity", productID)
	}
	product.Quantity += quantity
	st.Products[productID] = product
	return nil
}

func main() {

	storage := &Storage{
		Products: make(map[int]Product),
	}

	err := storage.AddProduct(Product{1, "Knife", 13})

	if err != nil {
		fmt.Println("Error adding product:", err)
		return
	}

	fmt.Println(storage.Products[1])

	err = storage.UpdateQuanity(5, 1)
	if err != nil {
		fmt.Println("Error updating quantity:", err)
		return
	}

	fmt.Println(storage.Products[1])
}
