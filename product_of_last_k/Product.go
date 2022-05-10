package product_of_last_k

type ProductOfNumbers struct {
	prefixProduct []int
	lastProduct   int
	isZeroPresent bool
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{prefixProduct: []int{}, lastProduct: 1, isZeroPresent: false}
}

func (product *ProductOfNumbers) Add(num int) {
	clearPrefixProduct := func() {
		for index := 0; index < len(product.prefixProduct); index++ {
			product.prefixProduct[index] = 0
		}
	}
	if num == 0 {
		clearPrefixProduct()
		product.isZeroPresent = true
		product.prefixProduct = append(product.prefixProduct, 0)
		product.lastProduct = 1
	} else {
		product.prefixProduct = append(product.prefixProduct, num*product.lastProduct)
		product.lastProduct = product.prefixProduct[len(product.prefixProduct)-1]
	}
}

func (product *ProductOfNumbers) GetProduct(k int) int {
	size := len(product.prefixProduct)
	totalProduct := product.prefixProduct[size-1]
	if k == size {
		if product.isZeroPresent {
			return 0
		}
		return totalProduct
	}
	remainingProductIndex := size - k - 1
	if product.prefixProduct[remainingProductIndex] == 0 {
		if product.prefixProduct[remainingProductIndex+1] == 0 {
			return 0
		}
		return totalProduct / 1
	}
	remainingProduct := product.prefixProduct[remainingProductIndex]
	return totalProduct / remainingProduct
}
