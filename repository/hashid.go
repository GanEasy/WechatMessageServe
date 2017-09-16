package repository

import "github.com/speps/go-hashids"

// Hid ...
var Hid *hashids.HashID

func init() {
	hd := hashids.NewData()
	hd.Salt = "thisismysaltpleaseletgo"
	hd.MinLength = 32
	Hid, _ = hashids.NewWithData(hd)
	// e, _ := Hid.Encode([]int{45, 434, 1313, 99, 2017091416})
	// fmt.Println(e)
	// d, _ := Hid.DecodeWithError(e)
	// fmt.Println(d)
}

//Encode 加密
func Encode(ids []int) (s string) {
	s, _ = Hid.Encode(ids)
	return s
}

//Decode 解密
func Decode(s string) (ids []int) {
	ids, _ = Hid.DecodeWithError(s)
	return ids
}
