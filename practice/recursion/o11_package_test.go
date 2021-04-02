package recursion

import (
	"fmt"
	"math"
	"testing"
)

// 背包问题
//有N件物品和一个容量为N的背包。第i件物品的重量是w[i]，价值是v[i]。
//求解将哪些物品装入背包可使这些物品的重量总和不超过背包容量，且价值总和最大。

type Product struct {
	W int // 重量
	V int // 价值
}

// 拿	：	f(n,pIndex) = Product.V + f(n - Product.W, pIndex + 1)
// 不拿  ：  f(n,pIndex) = f(n, pIndex + 1)
// f(n,pIndex) = MAX {拿,不拿}
func aboutPackage(products []Product, n, pIndex int) int {
	// 没有商品可以继续拿了
	if pIndex == len(products) {
		return 0
	}
	// 只有一个并且剩余重量够  则最优策略就是"拿"
	if pIndex == len(products)-1 && n >= products[pIndex].W {
		return products[pIndex].V
	}
	// 要不起... 尝试下一个
	if n < products[pIndex].W {
		return aboutPackage(products, n, pIndex+1)
	}
	max := 0
	// 要的起 ... 选择性拾取
	if n >= products[pIndex].W {
		// 拿
		way1 := aboutPackage(products, n-products[pIndex].W, pIndex+1) + products[pIndex].V
		// 不拿
		way2 := aboutPackage(products, n, pIndex+1)

		max = int(math.Max(float64(way1), float64(way2)))
	}
	return max
}

func TestPackage(t *testing.T) {
	p1 := Product{
		5,
		10,
	}
	p2 := Product{
		3,
		8,
	}
	p3 := Product{
		2,
		5,
	}
	p4 := Product{
		1,
		3,
	}
	productList := []Product{p1, p2, p3, p4}
	max := aboutPackage(productList, 8, 0)
	fmt.Println(max)
}
