package main

func main() {
	//a := make([]int, 10, 100)
	//
	//fmt.Printf("%T\n%v\n%v\n%v", a, a, len(a), cap(a))
	//
	//var runes []rune
	//for _, r := range "Hello, 世界" {
	//	runes = append(runes, r)
	//}
	//fmt.Printf("%\n", runes)

}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow . Extend the slice. 扩展切片
		z = x[:zlen]
	} else {
		// 空间不足，分配一个新的数组
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
