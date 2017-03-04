package goflint

/*
#cgo CFLAGS: -I/usr/local/include/flint -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -lgmp -lflint -lmpfr
#include <gmp.h>
#include <arith.h>
*/
import "C"

//NumberOfPartitions returns a pointer to a C.fmpz type containing p(n), the number of ways that
//n can be written as a sum of positive integers without regard to order.
func NumberOfPartitions(n uint) *C.fmpz {
	var x C.fmpz
	C.arith_number_of_partitions(&x, C.ulong(n))
	return &x
}

//NumberOfPartitionsMPFR returns a pointer to a C.mpfr type containing p(n), the number of ways that
//n can be written as a sum of positive integers without regard to order.
func NumberOfPartitionsMPFR(n uint) *C.mpfr {
	var x C.mpfr
	C.mpfr_init(&x)
	C.arith_number_of_partitions_mpfr(&x, C.ulong(n))
	return &x
}

//HRRExpsumFactored functions similarly to the core FLINT function, except that
//it only requires inputs k and n and returns a pointer to the result as
//*C.trig_prod_struct.
func HRRExpsumFactored(k uint, n uint) *C.trig_prod_struct {
	var prod C.trig_prod_struct
	C.trig_prod_init(&prod)
	C.arith_hrr_expsum_factored(&prod, C.ulong(k), C.ulong(n))
	return &prod
}

//NumberOfPartitionsNmodVec computes the first len values of the partition function p(n) starting
//with p(0), modulo the modulus defined by mod.
func NumberOfPartitionsNmodVec(len int, mod uint) {
	var res C.mp_limb_t
	var nmod C.nmod_t
	C.nmod_init(&nmod, C.ulong(mod))
	C.arith_number_of_partitions_nmod_vec(&res, C.slong(len), nmod)
}

//NumberOfPartitionsVec computes first len values of the partition function p(n), starting with
//p(0)
func NumberOfPartitionsVec(len uint) *C.fmpz {
	var res C.fmpz
	C.arith_number_of_partitions(&res, C.ulong(len))
	return &res
}
