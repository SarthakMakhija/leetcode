package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	return productOf(nums)
}

func productOf(nums []int) []int {
	var productOfInner func(index int)

	var products = make([]int, len(nums))
	products[0] = 1
	var pendingProduct = 1

	productOfInner = func(index int) {
		if index >= len(nums) {
			return
		}
		products[index] = products[index-1] * nums[index-1]
		productOfInner(index + 1)
		if index != len(nums)-1 {
			products[index] = products[index] * pendingProduct
		}
		pendingProduct = pendingProduct * nums[index]
	}
	productOfInner(1)
	products[0] = products[0] * pendingProduct
	return products
}
