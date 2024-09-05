package main

import (
	"fmt"
	"strings"
)

// https://en.wikipedia.org/wiki/SHA-2
// https://github.com/B-Con/crypto-algorithms/blob/master/sha256.c

type Sha256Data struct {
	original []byte
	k        [64]uint32
	m        [64]uint32
	data     [64]byte
	datalen  uint32
	bitlen   uint64
	state    [8]uint32
	hash     [32]byte
}

func NewSha256Data() *Sha256Data {
	var d Sha256Data
	d.Initialize()
	return &d
}

func (data *Sha256Data) Initialize() {

	data.k[0] = 0x428a2f98
	data.k[1] = 0x71374491
	data.k[2] = 0xb5c0fbcf
	data.k[3] = 0xe9b5dba5
	data.k[4] = 0x3956c25b
	data.k[5] = 0x59f111f1
	data.k[6] = 0x923f82a4
	data.k[7] = 0xab1c5ed5

	data.k[8] = 0xd807aa98
	data.k[9] = 0x12835b01
	data.k[10] = 0x243185be
	data.k[11] = 0x550c7dc3
	data.k[12] = 0x72be5d74
	data.k[13] = 0x80deb1fe
	data.k[14] = 0x9bdc06a7
	data.k[15] = 0xc19bf174

	data.k[16] = 0xe49b69c1
	data.k[17] = 0xefbe4786
	data.k[18] = 0x0fc19dc6
	data.k[19] = 0x240ca1cc
	data.k[20] = 0x2de92c6f
	data.k[21] = 0x4a7484aa
	data.k[22] = 0x5cb0a9dc
	data.k[23] = 0x76f988da

	data.k[24] = 0x983e5152
	data.k[25] = 0xa831c66d
	data.k[26] = 0xb00327c8
	data.k[27] = 0xbf597fc7
	data.k[28] = 0xc6e00bf3
	data.k[29] = 0xd5a79147
	data.k[30] = 0x06ca6351
	data.k[31] = 0x14292967

	data.k[32] = 0x27b70a85
	data.k[33] = 0x2e1b2138
	data.k[34] = 0x4d2c6dfc
	data.k[35] = 0x53380d13
	data.k[36] = 0x650a7354
	data.k[37] = 0x766a0abb
	data.k[38] = 0x81c2c92e
	data.k[39] = 0x92722c85

	data.k[40] = 0xa2bfe8a1
	data.k[41] = 0xa81a664b
	data.k[42] = 0xc24b8b70
	data.k[43] = 0xc76c51a3
	data.k[44] = 0xd192e819
	data.k[45] = 0xd6990624
	data.k[46] = 0xf40e3585
	data.k[47] = 0x106aa070

	data.k[48] = 0x19a4c116
	data.k[49] = 0x1e376c08
	data.k[50] = 0x2748774c
	data.k[51] = 0x34b0bcb5
	data.k[52] = 0x391c0cb3
	data.k[53] = 0x4ed8aa4a
	data.k[54] = 0x5b9cca4f
	data.k[55] = 0x682e6ff3

	data.k[56] = 0x748f82ee
	data.k[57] = 0x78a5636f
	data.k[58] = 0x84c87814
	data.k[59] = 0x8cc70208
	data.k[60] = 0x90befffa
	data.k[61] = 0xa4506ceb
	data.k[62] = 0xbef9a3f7
	data.k[63] = 0xc67178f2

}

// Operations
func ROTLEFT(a uint32, b int) uint32 {
	return (a << b) | (a >> (32 - b))
}
func ROTRIGHT(a uint32, b int) uint32 {
	return (a >> b) | (a << (32 - b))
}
func CH(x, y, z uint32) uint32 {
	notx := ^x
	return (x & y) ^ (notx & z)
}
func MAJ(x, y, z uint32) uint32 {
	return (x & y) ^ (x & z) ^ (y & z)
}
func EP0(x uint32) uint32 {
	return ROTRIGHT(x, 2) ^ ROTRIGHT(x, 13) ^ ROTRIGHT(x, 22)
}
func EP1(x uint32) uint32 {
	return ROTRIGHT(x, 6) ^ ROTRIGHT(x, 11) ^ ROTRIGHT(x, 25)
}
func SIG0(x uint32) uint32 {
	return ROTRIGHT(x, 7) ^ ROTRIGHT(x, 18) ^ (x >> 3)
}
func SIG1(x uint32) uint32 {
	return ROTRIGHT(x, 17) ^ ROTRIGHT(x, 19) ^ (x >> 10)
}

func (data *Sha256Data) CalculateSha256() {
	// Initialize hash values:
	// (first 32 bits of the fractional parts of the square roots of the first 8 primes 2..19):
	data.state[0] = 0x6a09e667
	data.state[1] = 0xbb67ae85
	data.state[2] = 0x3c6ef372
	data.state[3] = 0xa54ff53a
	data.state[4] = 0x510e527f
	data.state[5] = 0x9b05688c
	data.state[6] = 0x1f83d9ab
	data.state[7] = 0x5be0cd19

	data.datalen = 0
	data.bitlen = 0

	data.Update()
	data.Final()
}

func (data *Sha256Data) Update() {
	length := len(data.original)
	for i := 0; i < length; i++ {
		data.data[data.datalen] = data.original[i]
		data.datalen++
		if data.datalen == 64 {
			data.Transform()
			data.bitlen += 512
			data.datalen = 0
		}
	}
}

func (data *Sha256Data) Transform() {
	var a, b, c, d, e, f, g, h, i, j, t1, t2 uint32

	for i, j = 0, 0; i < 16; i, j = i+1, j+4 {
		data.m[i] = (uint32(data.data[j]) << 24) | (uint32(data.data[j+1]) << 16) | (uint32(data.data[j+2]) << 8) | (uint32(data.data[j+3]))
	}
	for ; i < 64; i++ {
		data.m[i] = SIG1(data.m[i-2]) + data.m[i-7] + SIG0(data.m[i-15]) + data.m[i-16]
	}

	a = data.state[0]
	b = data.state[1]
	c = data.state[2]
	d = data.state[3]
	e = data.state[4]
	f = data.state[5]
	g = data.state[6]
	h = data.state[7]

	for i = 0; i < 64; i++ {
		t1 = h + EP1(e) + CH(e, f, g) + data.k[i] + data.m[i]
		t2 = EP0(a) + MAJ(a, b, c)
		h = g
		g = f
		f = e
		e = d + t1
		d = c
		c = b
		b = a
		a = t1 + t2
	}

	data.state[0] += a
	data.state[1] += b
	data.state[2] += c
	data.state[3] += d
	data.state[4] += e
	data.state[5] += f
	data.state[6] += g
	data.state[7] += h
}

func (data *Sha256Data) Final() {
	var i uint32 = data.datalen

	// Pad whatever data is left in the buffer.
	if data.datalen < 56 {
		// binary 10000000
		data.data[i] = 0x80
		i++
		for i < 56 {
			data.data[i] = 0
			i++
		}
	} else {
		data.data[i] = 0x80
		i++
		for i < 64 {
			data.data[i] = 0
			i++
		}
		data.Transform()
		for j := 0; j < 56; j++ {
			data.data[j] = 0
		}
	}

	// Append to the padding the total message's length in bits and transform.
	data.bitlen += uint64(data.datalen) * 8
	data.data[63] = byte(data.bitlen & 0xFF)
	data.data[62] = byte((data.bitlen >> 8) & 0xFF)
	data.data[61] = byte((data.bitlen >> 16) & 0xFF)
	data.data[60] = byte((data.bitlen >> 24) & 0xFF)
	data.data[59] = byte((data.bitlen >> 32) & 0xFF)
	data.data[58] = byte((data.bitlen >> 40) & 0xFF)
	data.data[57] = byte((data.bitlen >> 48) & 0xFF)
	data.data[56] = byte((data.bitlen >> 56) & 0xFF)
	data.Transform()

	// Since this implementation uses little endian byte ordering and SHA uses big endian,
	// reverse all the bytes when copying the final state to the output hash.
	for index := 0; index < 8; index++ {
		for part := 0; part < 4; part++ {
			data.hash[part+index*4] = byte((data.state[index] >> (24 - part*8)) & 0xFF)
		}
	}
}

func (data *Sha256Data) GetHashString() string {
	var result strings.Builder
	for _, v := range data.hash {
		si := fmt.Sprintf("%02x", v)
		result.WriteString(si)
	}
	return result.String()
}
