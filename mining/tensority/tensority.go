package tensority

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L./lib/ -l:cSimdTs.o -lstdc++
// #include "./lib/cSimdTs.h"
import "C"

import(
    "unsafe"

    "github.com/bytom/protocol/bc"
)

func Hash(blockHeader, seed *bc.Hash) *bc.Hash {
    bhBytes := blockHeader.Bytes()
    sdBytes := seed.Bytes()

    bhPtr := (*C.uchar)(unsafe.Pointer(&bhBytes[0]))
    seedPtr := (*C.uchar)(unsafe.Pointer(&sdBytes[0]))

    resPtr := C.SimdTs(bhPtr, seedPtr)
    
    res := bc.NewHash(*(*[32]byte)(unsafe.Pointer(resPtr)))
    return &res
}