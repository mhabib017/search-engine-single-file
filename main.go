package main

import (
	"fmt"
	"os"
	"strings"
)

//------------------------//
//------------------------//
//---------Trie starts----//
//------------------------//
//------------------------//

type Node struct {
	children [26]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func InitTie() *Trie {
	trie:= Trie{root: &Node{}}
	return &trie
}

func (t *Trie) insert(w string) {
	w= strings.ToLower(w)
	wordLength:= len(w)
	currentNode:= t.root
	for i:=0; i< wordLength; i++{
		currentIndex:= w[i]- 'a'
		if currentNode.children[currentIndex] == nil {
			currentNode.children[currentIndex] = &Node{}
		}
		currentNode = currentNode.children[currentIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) search(w string) bool {
	w= strings.ToLower(w)
	wordLength:= len(w)
	currentNode:= t.root
	for i:=0; i< wordLength; i++{
		currentIndex:= w[i]- 'a'
		if currentNode.children[currentIndex] == nil {
			return false
		}
		currentNode = currentNode.children[currentIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

//------------------------//
//------------------------//
//---------Trie ends------//
//------------------------//
//------------------------//

func getTextFromFile(file string) []string{
	data, err := os.ReadFile(file)
	if err != nil {
        return make([]string, 0)
    }
	content:= string(data)
	// fmt.Println(content)
	lines := strings.Split(content, "\n")
	return lines
	
}

func getLineNoTokenMap(lines []string) map[int][]string{
	m := make(map[int][]string)
	for i,v:=range  lines {
		m[i] = strings.Fields(v)
	}
	return m
}

func getTokenMap(mp map[int][]string) map[string][]int {
	m := make(map[string][]int)
	for i:=0; i < len(mp); i++ {
		for _,s:=range  mp[i] {
			m[s]= append(m[s], i)
		}
	}
	return m
}

func getTokenArray(m map[string][]int) []string {
	ar:= make([]string,0)
	for key,_:=range m {
		ar= append(ar, key)
	}
	return ar
}

func main(){ 
	txt:= ""
	if len(os.Args) <2 {
		fmt.Println("String is required.")
		return
	} 
	txt = os.Args[1]

	lines:=getTextFromFile("input.txt")
	// for _,v:= range lines {
	// fmt.Println(v)
	// }

	lineNoTokenMap := 	getLineNoTokenMap(lines)
	// for i:=0; i < len(lineNoTokenMap); i++ {
	// 	fmt.Println("Token No: ", i)
	// 	for _,s:=range  lineNoTokenMap[i] {
	// 		fmt.Println(s)
	// 	}
	// }

	tokenMap := getTokenMap(lineNoTokenMap)
	// for key,value:= range tokenMap {
	// 	fmt.Println(key," => ",value)
	// }
	tokenArray := getTokenArray(tokenMap)
	// for _,v:= range tokenArray {
	// 	fmt.Println(v)
	// }

	trie:= InitTie()
	for _,v:= range tokenArray {
		trie.insert(v)
	}
	
	if !trie.search(txt) {
		fmt.Println("Match not found for ==> ", txt)
		return
	}
	fmt.Println("Match found for ==> ", txt)
	fmt.Println("Following are the found matches:")
	lineNumbers:= tokenMap[txt]
	for _,v:=range lineNumbers {
		fmt.Printf("[%d] %s\n",v+1, strings.Join(lineNoTokenMap[v], " ") )
	}
}