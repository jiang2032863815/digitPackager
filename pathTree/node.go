package pathTree

type Node struct{
	PathAddr []string
	Son map[string]*Node
	Count int
}
func newNode(pathAddr []string)*Node{
	var nd=&Node{}
	nd.PathAddr=pathAddr
	nd.Son=map[string]*Node{}
	nd.Count=1
	return nd
}
func(nd *Node)insert(pathAddr []string)int{
	if len(pathAddr)==0{
		return 0
	}
	var tCnt=0
	if nd.Son[pathAddr[0]]==nil{
		nd.Son[pathAddr[0]]=newNode(append(nd.PathAddr, pathAddr[0]))
		nd.Count++
		tCnt++
	}
	if len(pathAddr)>1 {
		var t=nd.Son[pathAddr[0]].insert(pathAddr[1:])
		tCnt+=t
		nd.Count+=t
	}
	return tCnt
}
func(nd *Node)delete(pathAddr []string)int{
	if len(pathAddr)==0{
		return 0
	}
	if len(pathAddr)==1{
		if nd.Son[pathAddr[0]]!=nil{
			var ret=nd.Son[pathAddr[0]].Count
			delete(nd.Son,pathAddr[0])
			nd.Count-=ret
			return ret
		}
		return 0
	}
	if nd.Son[pathAddr[0]]!=nil{
		var t=nd.Son[pathAddr[0]].delete(pathAddr[1:])
		nd.Count-=t
		return t
	}
	return 0
}