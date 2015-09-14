#include <string.h>

#define NAME "John"

int copyToBuffer(unsigned char *buf, const char *name) {
	memcpy(buf, name, sizeof(*name));
	return(1);
}