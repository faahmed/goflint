package main

import "fmt"
import "goflint"

/*
#cgo CFLAGS: -I/usr/local/include/flint -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -lgmp -lflint -lmpfr
#include <gmp.h>
#include <arith.h>
*/
import "C"

func main() {
	a := *(goflint.NumberOfPartitions(6))
	fmt.Println(a)

	b := *(goflint.NumberOfPartitionsMPFR(1000))
	//fmt.Println(C.mpfr_get_d(&b, 1))
        fmt.Println(b)

	c := *(goflint.HRRExpsumFactored(15, 10))
	fmt.Println(c)

	d := *(goflint.NumberOfPartitionsVec(20))
	fmt.Println(d)
}
