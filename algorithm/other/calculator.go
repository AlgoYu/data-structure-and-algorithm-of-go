/*
中缀计算器：输入一串无括号的中缀表达式，计算出结果。
主要思想：
	1. 创建一个计算符栈和数字栈。
	2. 扫描输入的字符串，把数字推入数字栈，把计算符推入计算符栈。
	3. 推入计算符的时候判断当前计算符与栈中计算符的优先级来处理
	4. 如果栈中的优先级等于或者大于当前的优先级，先从数字栈中取出两个值进行计算并把结果推入数字栈，否则直接把计算符推入计算符栈。
	5. 最后的数字栈中的结果便是表达式结果。
*/
package other

import "anydevelop.cn/data_structure/linear"

type Calculator struct {
	nums *linear.ArrayStack
	operators *linear.ArrayStack
}

// 初始化计算器
func (calculator *Calculator) Init() {
	calculator.nums = new(linear.ArrayStack)
	calculator.operators = new(linear.ArrayStack)
	calculator.nums.Init(20)
	calculator.operators.Init(20)
}

// 计算并返回结果
func (calculator *Calculator) CalculateExpression(expression string) int {
	// 扫描字符串存入相应的栈中
	var temp byte
	for i:= 0; i < len(expression); i++ {
		temp = expression[i]
		// 判断是数值还是操作符
		if calculator.IsOperator(int(temp)){
			// 如果栈中操作符等于或者大于当前操作符即从数值栈中取出两个值进行计算并推回数值栈
			if !calculator.operators.IsEmpty() && calculator.GetPriority(int(temp))<= calculator.GetPriority(calculator.operators.Peek()){
				num1 := calculator.nums.Pop()
				num2 := calculator.nums.Pop()
				operator := calculator.operators.Pop()
				calculator.nums.Push(calculator.Calculate(num1,num2,operator))
			}
			calculator.operators.Push(int(temp))
		}else{
			calculator.nums.Push(int(temp-48))
		}
	}
	// 反复取出操作符栈和数值栈的数据并计算
	for {
		if calculator.operators.IsEmpty(){
			break
		}
		num1 := calculator.nums.Pop()
		num2 := calculator.nums.Pop()
		operator := calculator.operators.Pop()
		calculator.nums.Push(calculator.Calculate(num1,num2,operator))
	}
	return calculator.nums.Pop()
}

// 判断是否为操作符
func (calculator *Calculator) IsOperator(value int) bool {
	return value == '+' || value == '-' || value == '*' || value == '/'
}

// 获取优先级
func (calculator *Calculator) GetPriority(operator int) int {
	switch operator {
	case '*':
		fallthrough
	case '/':
		return 1
	case '+':
		fallthrough
	case '-':
		return 0
	default:
		return -1
	}
}

// 计算数值
func (calculator *Calculator) Calculate(num1,num2,operator int) int {
	// 加法和乘法的顺序不影响计算结果，但是减法和除法的顺序会影响计算结果。
	switch operator {
	case '+':
		return num1+num2
	case '-':
		return num2-num1
	case '*':
		return num1*num2
	case '/':
		return num2/num1
	default:
		return 0
	}
}