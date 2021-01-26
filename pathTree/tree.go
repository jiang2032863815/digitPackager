package pathTree

import (
	"bytes"
	"encoding/gob"
)

type Tree struct {
	Root *Node
}
func(tr *Tree)Insert(pathAddr []string)int{
	return tr.Root.insert(pathAddr)
}
func(tr *Tree)Delete(pathAddr []string)int{
	return tr.Root.delete(pathAddr)
}
func(tr *Tree)ToBytes()[]byte{
	var bf=bytes.Buffer{}
	var encoder=gob.NewEncoder(&bf)
	encoder.Encode(tr)
	return bf.Bytes()
}
func MakeTree()*Tree{
	var tr=&Tree{newNode(nil)}
	return tr
}
func MakeTreeFromBytes(bs []byte)(*Tree,error){
	var bf=bytes.Buffer{}
	bf.Write(bs)
	var decoder=gob.NewDecoder(&bf)
	var tr = &Tree{}
	if err:=decoder.Decode(&tr);err==nil {
		return tr, nil
	}else{
		return nil,err
	}
}