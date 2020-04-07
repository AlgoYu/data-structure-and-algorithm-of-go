/*
表达式转换器：输入带括号中缀表达式，转换为后缀表达式。
主要思想：
	1. 创建两个栈。
	2. 把字符推入操作符栈，把数字推入数值栈。
	3. 左括号直接入操作符栈，右括号取出操作符栈左括号前所有操作符压入数值栈，若是操作符比较操作符栈中的优先级，优先级大于直接压入操作符栈，优先级小于操作符栈中的优先级则先取出操作符栈的操作符推入数值栈再把当前操作符推入操作符栈。
	4. 把所有操作符栈取出并压入数值栈。
	5. 头拼取出所有数值栈。
*/
package other

import (
	"anydevelop.cn/data_structure/linear"
	"strconv"
	"strings"
)

type ExpressionConverter struct {
	st1 *linear.ArrayStack
	st2 *linear.ArrayStack
}

func (expressionConverter *ExpressionConverter) Init() {
	expressionConverter.st1 = new(linear.ArrayStack)
	expressionConverter.st2 = new(linear.ArrayStack)
	expressionConverter.st1.Init(30)
	expressionConverter.st2.Init(30)
}

// 转换表达式
func (expressionConverter *ExpressionConverter) ConversionExpression(expression string) string {
	suffixExpression := ""
	values := strings.Split(expression, " ")
	// 遍历数组
	for i:=0; i< len(values); i++ {
		// 判断是否为操作符
		if expressionConverter.IsOperator(values[i]){
			// 左括号直接推入st1
			if values[i] == "("{
				bytes := []byte(values[i])
				expressionConverter.st1.Push(int(bytes[0]))
				// 右括号则一直取出st1的直到左括号为止的所有字符推入st2
			}else if values[i] == ")"{
				for string(byte(expressionConverter.st1.Peek()))!="("{
					expressionConverter.st2.Push(expressionConverter.st1.Pop())
				}
				expressionConverter.st1.Pop()
				// 操作符小于st1中的优先级则先取出st1的操作符推入st2再把当前操作符推入st1
			}else if !expressionConverter.st1.IsEmpty() && expressionConverter.GetPriority(values[i]) <= expressionConverter.GetPriority(string(byte(expressionConverter.st1.Peek()))){
				expressionConverter.st2.Push(expressionConverter.st1.Pop())
				bytes := []byte(values[i])
				expressionConverter.st1.Push(int(bytes[0]))
			}else{
				bytes := []byte(values[i])
				expressionConverter.st1.Push(int(bytes[0]))
			}
		}else{
			atoi, err := strconv.Atoi(values[i])
			if err!=nil{
				panic(err)
			}
			expressionConverter.st2.Push(atoi)
		}
	}
	for !expressionConverter.st1.IsEmpty(){
		expressionConverter.st2.Push(expressionConverter.st1.Pop())
	}
	for !expressionConverter.st2.IsEmpty(){
		value := expressionConverter.st2.Pop()
		if expressionConverter.IsOperator(string(byte(value))){
			suffixExpression = string(byte(value))+" "+suffixExpression
		}else{
			suffixExpression = strconv.Itoa(value)+" "+suffixExpression
		}
	}
	return suffixExpression
}

// 判断是否为操作符
func (expressionConverter *ExpressionConverter) IsOperator(value string) bool {
	return value == "+" || value == "-" || value == "*" || value == "/" || value == "(" || value == ")"
}

// 获取优先级
func (expressionConverter *ExpressionConverter) GetPriority(operator string) int {
	switch operator {
	case "*":
		fallthrough
	case "/":
		return 1
	case "+":
		fallthrough
	case "-":
		return 0
	default:
		return -1
	}
}