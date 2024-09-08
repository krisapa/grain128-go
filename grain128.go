package grain128

import "fmt"

type Grain128 struct {
	lfsr []byte
	nfsr []byte
	key  []byte
}

func NewGrain128(key []byte) (*Grain128, error) {
	if len(key) != 16 {
		return nil, fmt.Errorf("invalid key size: %d", len(key))
	}
	return &Grain128{
		lfsr: make([]byte, 128),
		nfsr: make([]byte, 128),
		key:  key,
	}, nil
}

func (c *Grain128) IVSetup(iv []byte) {
	ivBitSize := len(iv) * 8

	// Load nfsr with bits from key
	for i := 0; i < ivBitSize/8; i++ {
		for j := 0; j < 8; j++ {
			c.nfsr[i*8+j] = (c.key[i] >> j) & 1
			c.lfsr[i*8+j] = (iv[i] >> j) & 1
		}
	}

	// Load the remaining nfsr with bits from key and lfsr with 1
	for i := ivBitSize / 8; i < len(c.key); i++ {
		for j := 0; j < 8; j++ {
			c.nfsr[i*8+j] = (c.key[i] >> j) & 1
			c.lfsr[i*8+j] = 1
		}
	}

	// Initial clockings
	for i := 0; i < 256; i++ {
		outbit := c.keystream()
		c.lfsr[127] ^= outbit
		c.nfsr[127] ^= outbit
	}
}

func (c *Grain128) keystream() byte {
	// Calculate feedback and output bits
	outbit := c.nfsr[2] ^
		c.nfsr[15] ^
		c.nfsr[36] ^
		c.nfsr[45] ^
		c.nfsr[64] ^
		c.nfsr[73] ^
		c.nfsr[89] ^
		c.lfsr[93] ^
		(c.nfsr[12] & c.lfsr[8]) ^
		(c.lfsr[13] & c.lfsr[20]) ^
		(c.nfsr[95] & c.lfsr[42]) ^
		(c.lfsr[60] & c.lfsr[79]) ^
		(c.nfsr[12] & c.nfsr[95] & c.lfsr[95])

	nBit := c.lfsr[0] ^
		c.nfsr[0] ^
		c.nfsr[26] ^
		c.nfsr[56] ^
		c.nfsr[91] ^
		c.nfsr[96] ^
		(c.nfsr[3] & c.nfsr[67]) ^
		(c.nfsr[11] & c.nfsr[13]) ^
		(c.nfsr[17] & c.nfsr[18]) ^
		(c.nfsr[27] & c.nfsr[59]) ^
		(c.nfsr[40] & c.nfsr[48]) ^
		(c.nfsr[61] & c.nfsr[65]) ^
		(c.nfsr[68] & c.nfsr[84])

	lBit := c.lfsr[0] ^
		c.lfsr[7] ^
		c.lfsr[38] ^
		c.lfsr[70] ^
		c.lfsr[81] ^
		c.lfsr[96]

	keyBitSize := len(c.key) * 8

	// Update registers
	for i := 1; i < keyBitSize; i++ {
		c.nfsr[i-1] = c.nfsr[i]
		c.lfsr[i-1] = c.lfsr[i]
	}

	c.nfsr[keyBitSize-1] = nBit
	c.lfsr[keyBitSize-1] = lBit
	return outbit
}

func (c *Grain128) KeystreamBytes(keystream []byte) {
	for i := range keystream {
		keystream[i] = 0
		for j := 0; j < 8; j++ {
			keystream[i] |= c.keystream() << j
		}
	}
}

func (c *Grain128) XORKeyStream(dst, src []byte) {
	if len(dst) < len(src) {
		panic("dst is smaller than src")
	}
	for i := range src {
		var k byte
		for j := 0; j < 8; j++ {
			k |= c.keystream() << j
		}
		dst[i] = src[i] ^ k
	}
}
