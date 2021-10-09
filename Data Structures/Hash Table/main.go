package main

import "fmt"

const tableSize = 7

type Table struct {
	buckets [tableSize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key string
	next *bucketNode
}

func InitTable() *Table {
	table := &Table{}
	for i := 0; i < tableSize; i++ {
		table.buckets[i] = &bucket{}
	}
	return table
}

func (t *Table) Insert(n string) {
	hashed := hash(n)
	t.buckets[hashed].insert(n)
}

// Search
func (t *Table) search(n string) bool {
	hashed := hash(n)
	return t.buckets[hashed].search(n)
}

// Delete

// insert
func (b *bucket) insert(n string) {
	newBucketNode := &bucketNode{key: n}
	if b == nil {
		b.head = newBucketNode
		return
	}
	newBucketNode.next = b.head
	b.head = newBucketNode
}

// search
func (b *bucket) search(n string) bool {
	if b.head == nil {
		return false
	}
	if b.head.key == n {
		return true
	}
	currentNode := b.head
	for currentNode.next != nil {
		if currentNode.next.key == n {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete

func hash(n string) int {
	sum := 0
	for _, chr := range n {
		sum += int(chr)
	}
	return sum % tableSize
}

func main() {
	myTable := InitTable()
	fmt.Println(myTable)

	myTable.Insert("duc")
	myTable.Insert("hue")
	myTable.Insert("loc")
	myTable.Insert("thang")
	myTable.Insert("ran")

	//fmt.Println("duc :", myTable.search("duc"))
	//fmt.Println("thang :", myTable.search("thang"))
	//fmt.Println("hoc :", myTable.search("hoc"))
	//fmt.Println("hue :", myTable.search("hue"))
}