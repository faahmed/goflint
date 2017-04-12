# goflint
Some Go bindings for the FLINT2 library.

Note that the following compilation/linking flags

#cgo CFLAGS: -I/usr/local/include/flint -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -lgmp -lflint -lmpfr

will differ depending on where you've installed flint and its gmp/mpfr dependencies.

More information on FLINT is available at 

http://www.flintlib.org/
