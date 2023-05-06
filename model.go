package deepcopy

import "github.com/bitly/go-simplejson"

// Interface for delegating copy process to type
type Interface interface {
	DeepCopy() interface{}
}

// Iface is an alias to Copy; this exists for backwards compatibility reasons.
func Iface(iface interface{}) interface{} {
	return Copy(iface)
}

type SJ struct {
	Sjp *simplejson.Json
	Sj  simplejson.Json

	I *I
}

type I struct {
	A string
}

type PackData struct {
	a        int
	b        string
	C        []int
	D        map[string]int
	LogExtra *simplejson.Json
}

type SubData struct {
	A int
	B *int
}
