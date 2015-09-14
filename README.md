# tmlib
Code used to debug Swig typemaps

To build:

gcc -c -Wall -Werror -fpic tmlib.c
gcc -shared -o libtmlib.so tmlib.o

The unit tests are run using ginkgo, so you will need ginkgo and gomega:

go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
