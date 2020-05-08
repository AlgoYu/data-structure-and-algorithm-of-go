/*
哈夫曼编码：哈夫曼编码是一种可变字长编码(VLC)的一种算法，利用数据出现的概率构建哈夫曼树以获取异字头码字，也称为最佳编码。
主要思想：
	1. 使用HashMap统计字节出现的数量。
	2. 使用统计的数字构建哈弗曼树。
	3. 遍历哈弗曼树获取每个叶子节点存储字节的哈夫曼编码，往左为0，往右为1。
	4. 拼接哈夫曼编码，并以每8位存放为一个字节。
	5. 解压为逆向操作。
*/
package compression

import (
	"fmt"
	"strconv"
)

type HuffmanCodingNode struct {
	Data byte
	Weight int
	Left *HuffmanCodingNode
	Right *HuffmanCodingNode
}

type HuffmanCoding struct {
	HuffmanCode map[byte]string
}

func NewHuffmanCoding() *HuffmanCoding {
	return &HuffmanCoding{HuffmanCode: make(map[byte]string)}
}

// 哈夫曼压缩
func (huffmanCoding *HuffmanCoding) HuffmanCompression(bytes []byte)[]byte{
	huffmanCodingNodes := huffmanCoding.createNode(bytes)
	huffmanTree := huffmanCoding.createHuffmanTree(huffmanCodingNodes)
	huffmanCoding.createHuffmanCoding(huffmanTree,"","")
	huffmanStr := ""
	for i:=0; i < len(bytes); i++ {
		huffmanStr += huffmanCoding.HuffmanCode[bytes[i]]
	}
	fmt.Println(huffmanStr)
	huffmanCode := make([]byte,(len(huffmanStr)+15)/8)
	index := 0
	for i:=0; i < len(huffmanStr); i+=8 {
		if i+8 > len(huffmanStr){
			str := huffmanStr[i:]
			parseUint, err := strconv.ParseUint(str, 2, 8)
			if err!=nil{
				panic("Error to convert byte from string")
			}
			huffmanCode[index] = byte(parseUint)
			huffmanCode[index+1] = byte(len(str))
		}else{
			parseUint, err := strconv.ParseUint(huffmanStr[i:i+8], 2, 8)
			if err!=nil{
				panic("Error to convert byte from string")
			}
			huffmanCode[index] = byte(parseUint)
		}
		index++
	}
	return huffmanCode
}

// 获取字节节点
func (huffmanCoding *HuffmanCoding) createNode(bytes []byte) []*HuffmanCodingNode{
	var huffmanCodingNodes []*HuffmanCodingNode
	count := make(map[byte]int)
	for i:=0; i < len(bytes); i++ {
		count[bytes[i]] = count[bytes[i]]+1
	}
	for key,value := range count {
		huffmanCodingNodes = append(huffmanCodingNodes,&HuffmanCodingNode{Data: key,Weight: value})
	}
	return huffmanCodingNodes
}

// 创建哈夫曼树
func (huffmanCoding *HuffmanCoding) createHuffmanTree(huffmanCodingNoes []*HuffmanCodingNode) *HuffmanCodingNode {
	for len(huffmanCodingNoes) > 1 {
		huffmanCoding.bubbleSortHuffmanCodingNodes(huffmanCodingNoes)
		left := huffmanCodingNoes[0]
		right := huffmanCodingNoes[1]
		parent := new(HuffmanCodingNode)
		parent.Weight = left.Weight+right.Weight
		parent.Left = left
		parent.Right = right
		huffmanCodingNoes = append(huffmanCodingNoes,parent)
		huffmanCodingNoes = huffmanCodingNoes[2:]
	}
	return huffmanCodingNoes[0]
}

// 冒牌排序哈夫曼树
func (huffmanCoding *HuffmanCoding) bubbleSortHuffmanCodingNodes(huffmanCodingNoes []*HuffmanCodingNode) {
	var index int
	var temp *HuffmanCodingNode
	for i:=1; i < len(huffmanCodingNoes); i++ {
		index = i-1
		temp = huffmanCodingNoes[i]
		for index>=0 && temp.Weight < huffmanCodingNoes[index].Weight {
			huffmanCodingNoes[index+1] = huffmanCodingNoes[index]
			index--
		}
		huffmanCodingNoes[index+1] = temp
	}
}

// 创建哈夫曼编码
func (huffmanCoding *HuffmanCoding) createHuffmanCoding(node *HuffmanCodingNode,code,str string) {
	if node != nil{
		temp := str
		temp += code
		if node.Data!=0{
			huffmanCoding.HuffmanCode[node.Data] = temp
		}else{
			huffmanCoding.createHuffmanCoding(node.Left,"0",temp)
			huffmanCoding.createHuffmanCoding(node.Right,"1",temp)
		}
	}
}

// 哈夫曼解压
func (huffmanCoding *HuffmanCoding) HuffmanDecompression(bytes []byte) []byte {
	str := ""
	for i:=0; i< len(bytes);i++ {
		if i == len(bytes)-2{
			str+=byteToBinaryString(bytes[i], int(bytes[i+1]))
			break
		}
		str+=byteToBinaryString(bytes[i],0)
	}
	fmt.Println(str)
	deHuffmanCode := make(map[string]byte)
	for key,value := range huffmanCoding.HuffmanCode{
		deHuffmanCode[value]=key
	}
	source := make([]byte, 10)
	for i:=0;i< len(str); {
		index := 1
		for true {
			value := deHuffmanCode[str[i:i+index]]
			if value!=0{
				source = append(source, value)
				break
			}
			index++
		}
		i = i+index
	}
	return source
}

// 字节转字符串
func byteToBinaryString(data byte,count int) (str string) {
	var value byte
	for i:=0; i < 8; i++ {
		value = data
		data <<= 1
		data >>= 1
		switch (value) {
		case data: str += "0"
		default: str += "1"
		}
		data <<= 1
	}
	if count > 0{
		return str[len(str)-count:]
	}
	return str
}