package pathProcessor

import "path"

func Process(addr string)[]string{
	var resAddr string
	var L=len(addr)
	for i:=0;i<L;i++{
		if addr[i]=='\\'{
			resAddr+=string('/')
		}else{
			resAddr+=string(addr[i])
		}
	}
	resAddr=path.Clean(resAddr)
	var ret []string
	L=len(resAddr)
	var now string
	for i:=0;i<L;i++{
		if resAddr[i]=='/'{
			if len(now)>0 {
				ret = append(ret, now)
				now = ""
			}
		}else{
			now+=string(resAddr[i])
		}
	}
	if len(now)>0{
		ret=append(ret,now)
	}
	return ret
}