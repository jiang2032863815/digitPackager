package main

import (
	"digitPackager/pathProcessor"
	"digitPackager/pathTree"
	"fmt"
)
func printSpace(cnt int){
	for cnt>0{
		fmt.Print(" ")
		cnt--
	}
}
func dfs(nd *pathTree.Node,ex int){
	if nd==nil{
		return
	}
	printSpace(ex)
	fmt.Println(nd.PathAddr,nd.Count)
	for _,v:=range nd.Son{
		dfs(v,ex+3)
	}
}
func main(){
	var tr=pathTree.MakeTree()
	tr.Insert(pathProcessor.Process("root/a.txt"))
	tr.Insert(pathProcessor.Process("root/b.txt"))
	tr.Insert(pathProcessor.Process("root/c/a.txt"))
	tr.Insert(pathProcessor.Process("root/c/b.txt"))

}