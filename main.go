package main

import (
	"bufio"
	"digitPackager/pathProcessor"
	"digitPackager/pathTree"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path"
)

func createTree(bs []byte) *pathTree.Tree {
	if tr, err := pathTree.MakeTreeFromBytes(bs); err == nil {
		return tr
	} else {
		return pathTree.MakeTree()
	}
}
func fileExist(addr string) bool {
	_, err := os.Stat(addr)
	if err == nil {
		return true
	}
	return false
}
func copyFile(source, target string) {
	if sourceFile, err := os.Open(source); err == nil {
		defer sourceFile.Close()
		if targetFile, err := os.Create(target); err == nil {
			defer targetFile.Close()
			if _, err := io.Copy(targetFile, sourceFile); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}
func copyDir(source, target string) {
	if info, err := os.Stat(source); err == nil {
		if info.IsDir() {
			if err := os.Mkdir(target, os.ModeDir); err == nil {
				if sonInfo, err := ioutil.ReadDir(source); err == nil {
					for _, v := range sonInfo {
						copyDir(path.Join(source, v.Name()), path.Join(target, v.Name()))
					}
				} else {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			copyFile(source, target)
		}
	} else {
		fmt.Println(err)
	}
}
func dfsStore(nd *pathTree.Node, rootAddr string) {
	if nd == nil {
		return
	}
	if len(nd.PathAddr) != 0 {
		var addr = pathProcessor.ToAddr(nd.PathAddr)
		if info, err := os.Stat(addr); err == nil {
			if info.IsDir(){
				if nd.Count == 1 {
					copyDir(addr, pathProcessor.AddrFilter(path.Join(rootAddr, addr)))
				} else {
					if err := os.Mkdir(pathProcessor.AddrFilter(path.Join(rootAddr, addr)), os.ModeDir); err == nil {
						for _, v := range nd.Son {
							dfsStore(v, rootAddr)
						}
					}else{
						fmt.Println(err)
					}
				}
			} else {
				copyFile(addr, pathProcessor.AddrFilter(path.Join(rootAddr, addr)))
			}
		}else{
			fmt.Println(err)
		}
	} else {
		for _, v := range nd.Son {
			dfsStore(v, rootAddr)
		}
	}
}
func readFile(addr string) []byte {
	bs, _ := ioutil.ReadFile(addr)
	return bs
}
func saveFile(addr string, data []byte) {
	os.Remove(addr)
	ioutil.WriteFile(addr, data, os.ModeAppend)
}
func dfsShow(nd *pathTree.Node, ex int) {
	if nd == nil {
		return
	}
	for i := 1; i <= ex; i++ {
		fmt.Print(" ")
	}
	if len(nd.PathAddr) == 0 {
		fmt.Println("/")
	} else {
		fmt.Println(pathProcessor.ToAddr(nd.PathAddr))
	}
	for _, v := range nd.Son {
		dfsShow(v, ex+3)
	}
}
func parseCommandLine(data []byte) (string, []string) {
	var idx, L = 0, len(data)
	for idx < L && data[idx] == ' ' {
		idx++
	}
	var cmd string
	for idx < L && data[idx] != ' ' {
		cmd += string(data[idx])
		idx++
	}
	var params []string
	for idx < L {
		if data[idx] == '"' {
			idx++
			var t string
			for idx < L && data[idx] != '"' {
				t += string(data[idx])
				idx++
			}
			idx++
			params = append(params, t)
		} else {
			for idx < L && data[idx] != '"' && data[idx] == ' ' {
				idx++
			}
			var t string
			for idx < L && data[idx] != '"' && data[idx] != ' ' {
				t += string(data[idx])
				idx++
			}
			params = append(params, t)
		}
	}
	return cmd, params
}
func main() {
	var src = flag.String("src", "", "source file")
	flag.Parse()
	if *src == "" {
		log.Fatal("未指定数据包文件路径")
	}
	if path.Ext(*src) != ".package" {
		log.Fatal("not package file")
	}
	var tr = createTree(readFile(*src))
	defer func() {
		saveFile(*src, tr.ToBytes())
	}()
	go func() {
		var ch = make(chan os.Signal)
		signal.Notify(ch, os.Interrupt, os.Kill)
		for {
			<-ch
			fmt.Println()
			fmt.Println("please input quit")
			fmt.Print("digitPackager>")
		}
	}()
	var inputReader = bufio.NewReader(os.Stdin)
	for {
		fmt.Print("digitPackager>")
		var cmd string
		var params []string
		if line, _, err := inputReader.ReadLine(); err == nil {
			cmd, params = parseCommandLine(line)
		}
		switch cmd {
		case "quit":
			{
				fmt.Println("bye")
				return
			}
		case "add":
			{
				if len(params) < 1 {
					fmt.Println("error:need 1 param")
					break
				}
				var addr = params[0]
				if !fileExist(addr) {
					fmt.Println("error:not found")
				} else {
					fmt.Println("insert:", tr.Insert(pathProcessor.Process(addr)))
				}
			}
		case "del":
			{
				if len(params) < 1 {
					fmt.Println("error:need 1 param")
					break
				}
				var addr = params[0]
				fmt.Println("delete:", tr.Delete(pathProcessor.Process(addr)))
			}
		case "show":
			{
				dfsShow(tr.Root, 0)
			}
		case "store":
			{
				if len(params) < 1 {
					fmt.Println("error:need 1 param")
					break
				}
				var target = params[0]
				if fileExist(target) {
					fmt.Println("error:existed")
				} else {
					if err := os.Mkdir(target, os.ModeDir); err == nil {
						dfsStore(tr.Root, target)
					} else {
						fmt.Println(err)
					}
				}
			}
		default:
			{
				fmt.Println("command not found")
			}
		}
	}
}
