package main

// import "fmt"

func commonOmxn(aa []int, bb []int) []int {
	bigOCnt := 0
	o := make([]int, 0, len(aa)) //
	for _, a := range aa {
		for _, b := range bb {
			bigOCnt++
			if a == b {
				o = append(o, a)
			}
		}
	}
	//fmt.Printf("bigO count for commonOmaxmn = %v\n", bigOCnt)
	return o
}
