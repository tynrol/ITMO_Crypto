package main

import (
	"encoding/binary"
	"errors"
	"fmt"
)

const blocksize = 8
const lenght = 6*blocksize + 4

type ideaCipher struct {
	encryptKey [lenght]uint16
	decryptKey [lenght]uint16
}

func cipherInit(idea ideaCipher, key []byte) error {

	if l := len(key); l != 16 {
		return errors.New(fmt.Sprintf("Wrong key len %d", l))
	}

	for i := 0; i < blocksize; i++ {
		idea.encryptKey[i] = (uint16(key[0]) << 8) + uint16(key[1])
		key = key[2:]
	}

	//generate encryptrion key
	for i := blocksize; i < lenght; i++ {
		if (i % 8) == 6 {
			idea.encryptKey[i] = (idea.encryptKey[i-7] << 9) | (idea.encryptKey[i-14] >> 7)
		} else if (i % 8) == 7 {
			idea.encryptKey[i] = (idea.encryptKey[i-15] << 9) | (idea.encryptKey[i-14] >> 7)
		} else {
			idea.encryptKey[i] = (idea.encryptKey[i-7] << 9) | (idea.encryptKey[i-6] >> 7)
		}
	}

	//generate decryption key
	for i := 0; i < lenght; i += 6 {
		idea.decryptKey[i] = mulInv(idea.encryptKey[48-i])

		if i == 0 || i == 48 {
			idea.decryptKey[i+1] = -idea.encryptKey[49-i]
			idea.decryptKey[i+2] = -idea.encryptKey[50-i]
		} else {
			idea.decryptKey[i+1] = -idea.encryptKey[50-i]
			idea.decryptKey[i+2] = -idea.encryptKey[49-i]
		}

		idea.decryptKey[i+3] = mulInv(idea.encryptKey[51-i])

		if i < 48 {
			idea.decryptKey[i+4] = idea.decryptKey[46-i]
			idea.decryptKey[i+5] = idea.decryptKey[47-i]
		}
	}

	return nil
}

func cryptBlock(key []uint16, input, output []byte) {
	var x1, x2, x3, x4, s2, s3 uint16

	x1 = binary.BigEndian.Uint16(input[0:])
	x2 = binary.BigEndian.Uint16(input[2:])
	x3 = binary.BigEndian.Uint16(input[4:])
	x4 = binary.BigEndian.Uint16(input[6:])

	for r := blocksize; r > 0; r-- {

		x1 = mmul(x1, key[0])
		key = key[1:]
		x2 += key[0]
		key = key[1:]
		x3 += key[0]
		key = key[1:]

		x4 = mmul(x4, key[0])
		key = key[1:]

		s3 = x3
		x3 ^= x1
		x3 = mmul(x3, key[0])
		key = key[1:]
		s2 = x2

		x2 ^= x4
		x2 += x3
		x2 = mmul(x2, key[0])
		key = key[1:]
		x3 += x2

		x1 ^= x2
		x4 ^= x3

		x2 ^= s3
		x3 ^= s2

	}
	x1 = mmul(x1, key[0])
	key = key[1:]

	x3 += key[0]
	key = key[1:]
	x2 += key[0]
	key = key[1:]
	x4 = mmul(x4, key[0])

	binary.BigEndian.PutUint16(output[0:], x1)
	binary.BigEndian.PutUint16(output[2:], x3)
	binary.BigEndian.PutUint16(output[4:], x2)
	binary.BigEndian.PutUint16(output[6:], x4)
}

//x*y mod 2^16+1
func mmul(a, b uint16) uint16 {
	var c uint32

	c = uint32(a) * uint32(b)
	if c == 0 {
		return 1 - a - b
	}
	a = uint16(c)
	b = uint16(c >> 16)
	if a < b {
		return a - b + 1
	}
	return a - b
}

//inverse of x mod 2^16+1
func mulInv(a uint16) uint16 {
	var b, q, r uint32
	var t, u, v int32

	b = 0x10001
	u = 0
	v = 1

	for a > 0 {
		q = b / uint32(a)
		r = b % uint32(a)

		b = uint32(a)
		a = uint16(r)

		t = v
		v = u - int32(q)*v
		u = t
	}

	if u < 0 {
		u += 0x10001
	}

	return uint16(u)
}
