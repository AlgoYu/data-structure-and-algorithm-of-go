package other

import (
	"anydevelop.cn/data_structure/linear"
	"regexp"
	"strconv"
	"strings"
)

type ReversePolishCalculator struct {
	nums *linear.ArrayStack
}

// 初始化计算器
func (reversePolishCalculator *ReversePolishCalculator) Init() {
	reversePolishCalculator.nums = new(linear.ArrayStack)
	reversePolishCalculator.nums.Init(20)
}

// 计算并返回结果
func (reversePolishCalculator *ReversePolishCalculator) CalculateReversePolishExpression(expression string) int {
	// 以空格分割
	values := strings.Split(expression, " ")
	reg, err := regexp.Compile("\\d+")
	// 遍历数组判断是操作符还是数值
	if err!=nil{
		panic(err)
	}
	for i:=0; i< len(values); i++ {
		// 数值入栈
		if reg.MatchString(values[i]){
			atoi, err := strconv.Atoi(values[i])
			if err!=nil{
				panic(err)
			}
			reversePolishCalculator.nums.Push(atoi)
		}else{
			// 如果是操作符即从数值栈中取出两个值进行计算并推回数值栈
			num1 := reversePolishCalculator.nums.Pop()
			num2 := reversePolishCalculator.nums.Pop()
			reversePolishCalculator.nums.Push(reversePolishCalculator.Calculate(num1,num2,values[i]))
		}
	}
	// 返回结果
	return reversePolishCalculator.nums.Pop()
}

// 计算数值
func (reversePolishCalculator *ReversePolishCalculator) Calculate(num1,num2 int,operator string) int {
	// 加法和乘法的顺序不影响计算结果，但是减法和除法的顺序会影响计算结果。
	switch operator {
	case "+":
		return num1+num2
	case "-":
		return num2-num1
	case "*":
		return num1*num2
	case "/":
		return num2/num1
	default:
		panic("Operation mismatch.")
	}
}