package main

type user struct {
	name string
	surname string
	email string
	product []product
}

type product struct {
	name string
	price float64
	quantity int
}

// func (p *product) newProduct(pName string, pPrice float64) product{
// 	pName := product{name: pName, price: pPrice}
// }

// func addProduct(user user, product product, quantity int)  {
// 	user.product = append(user.product, product)
	
// }

func main()  {
	
}