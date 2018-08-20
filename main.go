package main

import (
	"../../GinMall/Helper"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	RouterInit(router)
	Helper.Dbinit()
	// r1 := compareVersion("1.2.3a", "1.2.4b")
	// println(r1)
	router.Run(":8081")
}

// func compareVersion(version1 string, version2 string) int {
// 	v1 := strings.Split(version1, ".")
// 	v2 := strings.Split(version2, ".")
// 	if len(v1) >= len(v2) {
// 		for i := 0; i < len(v2); i++ {
// 			if r1 := Scompare(v1[i], v2[i]); r1 != 0 {
// 				return r1
// 			}
// 		}
// 		for i := len(v2); i < len(v1); i++ {
// 			if r1 := Scompare(v1[i], "0"); r1 != 0 {
// 				return r1
// 			}
// 		}
// 		return 0
// 	} else {
// 		for i := 0; i < len(v1); i++ {
// 			if r1 := Scompare(v1[i], v2[i]); r1 != 0 {
// 				return r1
// 			}
// 		}
// 		for i := len(v1); i < len(v2); i++ {
// 			if r1 := Scompare("0", v2[i]); r1 != 0 {
// 				return r1
// 			}
// 		}
// 		return 0
// 	}
// }
// func Scompare(s1 string, s2 string) int {
// 	var s1r, s2r []rune
// 	var ls1, ls2 int
// 	s1r = []rune(s1)
// 	s2r = []rune(s2)
// 	ls1 = len(s1r)
// 	ls2 = len(s2r)
// 	if ls1 > ls2 {
// 		t1 := make([]rune, ls1-ls2)
// 		s2r = append(t1, s2r...)
// 		for i := 0; i < ls1; i++ {
// 			if s1r[i] > s2r[i] {
// 				return 1
// 			} else if s1r[i] < s2r[i] {
// 				return -1
// 			}
// 		}
// 	} else {
// 		t1 := make([]rune, ls2-ls1)
// 		s1r = append(t1, s1r...)
// 		for i := 0; i < ls2; i++ {
// 			if s1r[i] > s2r[i] {
// 				return 1
// 			} else if s1r[i] < s2r[i] {
// 				return -1
// 			}
// 		}
// 	}
// 	return 0
// }
