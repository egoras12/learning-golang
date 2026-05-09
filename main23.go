package main

type contact struct {
	userID       string
	sendingLimit int32
	age          int32
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

// bool: 1 byte
// int8: uint8 (byte): 1
// int16: uint16: 2
// int32: uint32 (rune): 4
// int64: uint64: 8
// int: uint, uintptr: 8
// float32: 4
// float64: 8
// complex64: 8
// complex128: 16
// string: 16 (pointer + length)
// slice: 24 (pointer + len + cap)
// map: 8 (pointer)
// chan: 8 (pointer)
// func: 8 (pointer)
// interface: 16 (type + data pointers)
// pointer *T: 8
// struct: sum of fields + padding
// array: element size * length

func main23 () {
	//  this only works when you run "go run main.go main22.go"
	main22()
}