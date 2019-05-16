/*
 * LF: Global Fully Replicated Key/Value Store
 * Copyright (C) 2018-2019  ZeroTier, Inc.  https://www.zerotier.com/
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 *
 * --
 *
 * You can be released from the requirements of the license by purchasing
 * a commercial license. Buying such a license is mandatory as soon as you
 * develop commercial closed-source software that incorporates or links
 * directly against ZeroTier software without disclosing the source code
 * of your own application.
 */

package lf

import (
	"crypto/aes"
	"crypto/sha512"
	"encoding/binary"
	"encoding/json"
	"runtime"
	"sync"
)

// OrdinalSize is the size of an ordinal in bytes.
const OrdinalSize = 16

// Ordinal is the sortable/comparable part of a selector.
// It consists of a 64-bit integer masked using a simple order-preserving keyed hash.
type Ordinal [16]byte

// MarshalJSON returns this blob marshaled as a \bbase62-encoded string (like a non-UTF8 Blob).
func (b *Ordinal) MarshalJSON() ([]byte, error) {
	return []byte("\"\\b" + Base62Encode(b[:]) + "\""), nil
}

// UnmarshalJSON unmarshals this blob from a JSON array or string
func (b *Ordinal) UnmarshalJSON(j []byte) error {
	if len(j) == 0 {
		for i := range b {
			b[i] = 0
		}
		return nil
	}

	var err error
	var bb []byte

	// Default is \bbase62string
	var str string
	err = json.Unmarshal(j, &str)
	if err == nil {
		if len(str) > 0 && str[0] == '\b' {
			bb = Base62Decode(str[1:])
		} else {
			bb = []byte(str) // be consistent with Blob even though an ordinal won't satisfy this that often
		}
	}

	// Byte arrays are also accepted
	if err != nil {
		if json.Unmarshal(j, &bb) != nil {
			return err
		}
	}

	i := 0
	for i < len(bb) && i < 16 {
		b[i] = bb[i]
		i++
	}
	for i < 16 {
		b[i] = 0
		i++
	}
	return nil
}

var ordinalParallelQuicksortThreshold = func() int {
	nc := runtime.NumCPU()
	if nc <= 2 {
		return 1048576 // no sort parallelism
	}
	if nc >= 16 {
		return 2048
	}
	return 32768 / nc
}()

var ordinalAlphabetPool = sync.Pool{
	New: func() interface{} {
		return make([]uint32, 32768)
	},
}

func ordinalParallelQuicksort(a []uint32, wg *sync.WaitGroup, par bool) {
	left, right := 0, len(a)-1
	a[1], a[right] = a[right], a[1]
	for i, ai := range a {
		if ai < a[right] {
			a[left], a[i] = ai, a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	if left >= 2 {
		if left >= ordinalParallelQuicksortThreshold {
			wg.Add(1)
			go ordinalParallelQuicksort(a[:left], wg, true)
		} else {
			ordinalParallelQuicksort(a[:left], wg, false)
		}
	}

	a = a[left+1:]
	if len(a) >= 2 {
		if len(a) >= ordinalParallelQuicksortThreshold {
			wg.Add(1)
			go ordinalParallelQuicksort(a, wg, true)
		} else {
			ordinalParallelQuicksort(a, wg, false)
		}
	}

	if par {
		wg.Done()
	}
}

func ordinal16to32(wg *sync.WaitGroup, valueMaskedToColumn uint64, columnValue uint, kk int, keyHash *[64]byte, result *[4]uint32) {
	var aesTmp [16]byte
	alphabet := ordinalAlphabetPool.Get().([]uint32)
	_ = alphabet[32767]

	c, _ := aes.NewCipher(keyHash[16*kk : 16*(kk+1)])

	c.Encrypt(aesTmp[:], aesTmp[:])
	rbase := binary.LittleEndian.Uint32(aesTmp[0:4]) % (0x7fffffff + 2)

	for {
		for i := 0; i < 32768; {
			c.Encrypt(aesTmp[:], aesTmp[:])
			alphabet[i] = (binary.LittleEndian.Uint32(aesTmp[0:4]) & 0x7fffffff) + rbase
			i++
			alphabet[i] = (binary.LittleEndian.Uint32(aesTmp[4:8]) & 0x7fffffff) + rbase
			i++
			alphabet[i] = (binary.LittleEndian.Uint32(aesTmp[8:12]) & 0x7fffffff) + rbase
			i++
			alphabet[i] = (binary.LittleEndian.Uint32(aesTmp[12:16]) & 0x7fffffff) + rbase
			i++
		}

		var sortWG sync.WaitGroup
		ordinalParallelQuicksort(alphabet, &sortWG, false)
		sortWG.Wait()

		for i := 1; i < 32768; i++ {
			if (alphabet[i] - alphabet[i-1]) < 2 {
				alphabet[i] += 2
			}
		}

		if alphabet[32767] > alphabet[0] && alphabet[32767] < 0xfffffffe {
			vdiv2 := columnValue >> 1
			rv := alphabet[vdiv2]
			if (columnValue & 1) != 0 {
				binary.LittleEndian.PutUint64(aesTmp[:], valueMaskedToColumn)
				c.Encrypt(aesTmp[:], aesTmp[:])
				rn := binary.LittleEndian.Uint32(aesTmp[0:4])
				if vdiv2 == 32767 {
					rv += rn % ^rv //(0xffffffff - rv)
				} else {
					rv += rn % ((alphabet[vdiv2+1] - rv) - 1)
				}
				rv++
			}
			result[kk] = rv
			ordinalAlphabetPool.Put(alphabet)
			wg.Done()
			return
		}
	}
}

// Set sets this ordinal to a sortable masked value that hides the original value using an order preserving keyed hash.
func (b *Ordinal) Set(value uint64, key []byte) {
	var result [4]uint32

	keyHash := sha512.Sum512(key)

	var wg sync.WaitGroup
	wg.Add(4)
	go ordinal16to32(&wg, 0, uint((value>>48)&0xffff), 0, &keyHash, &result)
	go ordinal16to32(&wg, value&0xffff000000000000, uint((value>>32)&0xffff), 1, &keyHash, &result)
	go ordinal16to32(&wg, value&0xffffffff00000000, uint((value>>16)&0xffff), 2, &keyHash, &result)
	ordinal16to32(&wg, value&0xffffffffffff0000, uint(value&0xffff), 3, &keyHash, &result)
	wg.Wait()

	binary.BigEndian.PutUint32(b[0:4], result[0])
	binary.BigEndian.PutUint32(b[4:8], result[1])
	binary.BigEndian.PutUint32(b[8:12], result[2])
	binary.BigEndian.PutUint32(b[12:16], result[3])
}
