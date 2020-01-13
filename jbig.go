package jbig

// #cgo LDFLAGS: -L./lib -ljbig
// #cgo CFLAGS: -D_LINUXES_ -g -O  -pedantic  -I./include
/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <limits.h>
#include "jbig.h"
int jbig2pbm(char* jbigbuff, int len, char* pbmbuf, int* pbmlen, char *errMsg){
	int result;
	size_t cnt;
	struct jbg_dec_state s;
	jbg_dec_init(&s);
	//解析JBIG
	result = jbg_dec_in(&s, (unsigned char *)jbigbuff, len, &cnt);
    if (!(result == JBG_EAGAIN || result == JBG_EOK)){
		strcpy(errMsg,jbg_strerror(result));
		jbg_dec_free(&s);
		return result;
    }
	//解析头部
	sprintf(pbmbuf, "P4\n%10lu\n%10lu\n",jbg_dec_getwidth(&s),jbg_dec_getheight(&s));
	int headlen = strlen(pbmbuf);
	//解析body
	unsigned char* body = jbg_dec_getimage(&s, 0);
	int size =  jbg_dec_getsize(&s);
	memcpy(pbmbuf + headlen, body, size);
	*pbmlen = headlen + size;
	jbg_dec_free(&s);
	return 0;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func JBIGConvToPBM(jbigbuf []byte) ([]byte, error) {
	var pbmlen C.int

	jbiglen := len(jbigbuf)

	cjbigbuf := C.CBytes(jbigbuf)
	defer C.free(cjbigbuf)

	pbmbuf := make([]byte, 10240)
	cpbmbuf := C.CBytes(pbmbuf)
	defer C.free(cpbmbuf)

	errmsg := make([]byte, 2000)
	cerrmsg := C.CBytes(errmsg)
	defer C.free(cerrmsg)

	ret, err := C.jbig2pbm((*C.char)(cjbigbuf), C.int(jbiglen), (*C.char)(cpbmbuf), &pbmlen, (*C.char)(cerrmsg))
	if err != nil {
		return nil, fmt.Errorf("%d:%s;%v", ret, C.GoString((*C.char)(cerrmsg)), err)
	}
	return C.GoBytes(unsafe.Pointer(cpbmbuf), C.int(pbmlen)), nil
}
