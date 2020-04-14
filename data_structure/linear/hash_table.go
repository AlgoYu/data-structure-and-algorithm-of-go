/*
散列表：散列表是一种线性结构，是根据关键码值(Key value)而直接进行访问的数据结构，把数据转换为存储位置的映射函数叫做散列函数，存放记录的数组叫做散列表。
主要思想：
	1. 创建一个数组作为散列表。
	2. 使用取模的方法作为散列函数。
	3. 使用散列散列函数转换存储地址。
	4. 数组中存储链表结构完整增删改查操作。
*/
package linear

import "fmt"

// 散列表节点
type HashTableNode struct {
	Id int
	Data int
	Next *HashTableNode
}

type HashTable struct {
	Size int
	HashTable []*HashTableNode
}

// 初始化散列表
func (hashTable *HashTable) Init(size int) {
	hashTable.Size = size
	hashTable.HashTable = make([]*HashTableNode,size)
}

// 散列函数
func (hashTable *HashTable) HashFunction(key int) int {
	return key % hashTable.Size
}

// 增加节点
func (hashTable *HashTable) AddNode(hashTableNode *HashTableNode) {
	if hashTable.GetNode(hashTableNode.Id) != nil{
		fmt.Println("The node already exist!")
		return
	}
	index := hashTable.HashFunction(hashTableNode.Id)
	if hashTable.HashTable[index] == nil{
		hashTable.HashTable[index] = hashTableNode
	}else{
		temp := hashTable.HashTable[index]
		for temp.Next!=nil {
			temp = temp.Next
		}
		temp.Next = hashTableNode
	}
}

// 增加顺序节点
func (hashTable *HashTable) AddOrderNode(hashTableNode *HashTableNode) {
	if hashTable.GetNode(hashTableNode.Id) != nil{
		fmt.Println("The node already exist!")
		return
	}
	index := hashTable.HashFunction(hashTableNode.Id)
	if hashTable.HashTable[index] == nil{
		hashTable.HashTable[index] = hashTableNode
	}else{
		temp := hashTable.HashTable[index]
		if temp.Id > hashTableNode.Id{
			hashTableNode.Next = temp
			hashTable.HashTable[index] = hashTableNode
		}else{
			for temp.Next!=nil && temp.Next.Id < hashTableNode.Id {
				temp = temp.Next
			}
			hashTableNode.Next = temp.Next
			temp.Next = hashTableNode
		}
	}
}

// 删除节点
func (hashTable *HashTable) DeleteNode(id int) {
	index := hashTable.HashFunction(id)
	temp := hashTable.HashTable[index]
	if temp!=nil{
		if temp.Id == id{
			hashTable.HashTable[index] = temp.Next
		}else{
			for temp.Next!=nil && temp.Next.Id!=id {
				temp = temp.Next
			}
			if temp.Next !=nil {
				temp.Next = temp.Next.Next
			}
		}
	}
}

// 修改节点
func (hashTable *HashTable) ModifyNode(hashTableNode *HashTableNode) {
	node := hashTable.GetNode(hashTableNode.Id)
	node.Data = hashTableNode.Data
}

// 获取节点
func (hashTable *HashTable) GetNode(id int) *HashTableNode {
	index := hashTable.HashFunction(id)
	temp := hashTable.HashTable[index]
	for temp!=nil && temp.Id!=id {
		temp = temp.Next
	}
	return temp
}

// 打印节点
func (hashTable *HashTable) PrintHashTable(){
	for i:=0; i< len(hashTable.HashTable);i++ {
		fmt.Println("The ",i," linked list:")
		temp := hashTable.HashTable[i]
		for temp!=nil {
			fmt.Print("id=",temp.Id,"data=",temp.Data,"=>")
			temp = temp.Next
		}
		fmt.Println()
	}
}