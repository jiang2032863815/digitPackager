package pathProcessor

import "runtime"

func filterBlock(addr string)string{
	var ret string
	var L=len(addr)
	for i:=0;i<L;i++{
		if addr[i]==':'||addr[i]=='?'||addr[i]=='*'||addr[i]=='|'||addr[i]=='<'||addr[i]=='>'{
			continue
		}
		ret+=string(addr[i])
	}
	return ret
}
func AddrFilter(addr string)string{
	var p=Process(addr)
	var idx int
	if runtime.GOOS=="windows"{
		idx=1
	}
	var L=len(p)
	for idx<L{
		p[idx]=filterBlock(p[idx])
		idx++
	}
	return ToAddr(p)
}