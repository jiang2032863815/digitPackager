package pathProcessor

import "path"

func ToAddr(pathAddr []string)string{
	return path.Join(pathAddr...)
}