package endian

import "encoding/binary"

var Native = binary.LittleEndian

const (
	IsBigEndian     = false
	NativeEndian    = LittleEndian
	NonNativeEndian = BigEndian
)
