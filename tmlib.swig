%module tmlib

%{
#include <string.h>
#include <tmlib.h>
%}

%include "typemaps.i"

%apply unsigned char *OUTPUT {unsigned char *buf};

int copyToBuffer(unsigned char *buf, const char *name);