package math

import (
	"fmt"
	"math"
	"testing"
)

type Vector []float64

func NewVector(num ...float64) Vector {
	if len(num) < 2 {
		panic("参数不合法")
	}
	return num
}

func (v Vector) Add(o Vector) Vector {
	checkCalculationLegitimacy(v, o)
	ret := make([]float64, len(v))
	for i := 0; i < len(ret); i++ {
		ret[i] = v[i] + o[i]
	}
	return ret
}

func (v Vector) Subtract(o Vector) Vector {
	checkCalculationLegitimacy(v, o)
	ret := make([]float64, len(v))
	for i := 0; i < len(ret); i++ {
		ret[i] = v[i] - o[i]
	}
	return ret
}

// 数量乘法
func (v Vector) ScalarMultiply(num float64) Vector {
	ret := make([]float64, len(v))
	for i := 0; i < len(ret); i++ {
		ret[i] = v[i] * num
	}
	return ret
}

// 点乘
func (v Vector) DotMultiply(o Vector) float64 {
	checkCalculationLegitimacy(v, o)
	var ret float64 = 0
	for i := 0; i < len(v); i++ {
		ret += v[i] * o[i]
	}
	return ret
}

// 标准化
func (v Vector) Normalize() Vector {
	mo := v.Norm()
	if mo == 0 {
		panic("零向量不可以被归一化")
	}
	return v.ScalarMultiply(1 / mo)
}

// 取模
func (v Vector) Norm() float64 {
	var ret float64 = 0
	for i := 0; i < len(v); i++ {
		ret += v[i] * v[i]
	}
	return math.Sqrt(ret)
}

func checkCalculationLegitimacy(v ...Vector) {
	before := -1
	for i := 0; i < len(v); i++ {
		tmp := len(v[i])
		if before == -1 {
			before = tmp
		} else if v[i] == nil || tmp != before {
			panic("无法参与运算, 参数不合法")
		}
	}
}

func (v Vector) Print() {
	str := "("
	for i := 0; i < len(v); i++ {
		str += fmt.Sprintf("%f", v[i])
		if i != len(v)-1 {
			str += ","
		}
	}
	str += ")"
	fmt.Println(str)
}

func TestVector(t *testing.T) {
	u := NewVector(6, 2, 3)
	v := NewVector(3, 4, 5)

	u.Add(v).Print()
	u.Subtract(v).Print()
	u.ScalarMultiply(2).Print()
	fmt.Printf("%f\n", u.DotMultiply(v))
	fmt.Printf("%f\n", u.Norm())
	n := u.Normalize()
	n.Print()
	fmt.Printf("%f\n", n.Norm())

	//角度=玄度/2/PI*360
	fmt.Printf("%f\n", math.Acos(math.Sqrt(3)/float64(2))*180/math.Pi)
}
