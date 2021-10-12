package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data string
	next	*Node
	previous *Node
}

type LinkList struct {
	root *Node
	length int
}

// Insert
// take a command and return a current Node
func (l *LinkList) Insert(currentNode *Node, cmd string) *Node {
	// for each dir
	arguments := cmd[3:]
	dirs := strings.Split(arguments, "/")
	for _, dir := range dirs {
		switch dir {
		case "":
			currentNode = l.root
			l.length = 1
		case "..":
			currentNode = currentNode.previous
			l.length--
		default:
			currentNode.next = &Node{data: dir}
			currentNode.next.previous = currentNode

			currentNode = currentNode.next
			l.length++
		}
	}
	return currentNode
}

func (l *LinkList) Pwd()  {
	fmt.Printf("/")
	currentNode := l.root
	for i := 1; i < l.length; i++ {
		fmt.Printf("%v/", currentNode.next.data)
		currentNode = currentNode.next
	}
	fmt.Printf(" ")
}

func main() {
	root := &Node{data: "/"}
	myList := &LinkList{
		root:   root,
		length: 1,
	}

	currentNode := root
	cmds := []string{"cd Desktop", "pwd", "cd /etc", "pwd", "cd apt", "pwd", "cd ../../home", "pwd"}
	//cmds := []string{"pwd", "cd /home/vasya", "pwd", "cd ..", "pwd", "cd vasya/../pesya", "pwd"}
	//cmds := []string{"cd /a/b", "pwd", "cd ../a/b", "pwd"}

	for _, cmd := range cmds {
		if cmd[0] == 99 {
			// insert
			currentNode = myList.Insert(currentNode, cmd)
		} else {
			// print
			myList.Pwd()
		}
	}
}

// output: /Desktop/ /etc/ /etc/apt/ /home/
// output: / /home/vasya/ /home/ /home/pesya/
// output: /a/b/ /a/a/b/