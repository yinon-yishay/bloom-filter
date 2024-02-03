package bitset

import (
	"fmt"
	"strings"
)

type BitSet struct {
	numOfBits int
	bitSet    []uint64
}

func New(bitNum int) *BitSet {
	return &BitSet{
		numOfBits: bitNum,
		bitSet:    make([]uint64, (bitNum/64)+1),
	}
}

func (bs *BitSet) SetBitOn(bitIndex uint64) {
	idx := getIndex(bitIndex)

	bs.bitSet[idx] = bs.bitSet[idx] | getBit(bitIndex)
}

func (bs *BitSet) SetBitOff(bitIndex uint64) {
	idx := getIndex(bitIndex)

	bs.bitSet[idx] = bs.bitSet[idx] &^ getBit(bitIndex)
}

func (bs *BitSet) IsBitOn(bitIndex uint64) bool {
	return (bs.bitSet[getIndex(bitIndex)] & getBit(bitIndex)) != 0
}

func (bs *BitSet) IsBitOff(bitIndex uint64) bool {
	return !bs.IsBitOn(bitIndex)
}

func (bs *BitSet) ClearAll() {
	for i := range bs.bitSet {
		bs.bitSet[i] = 0
	}
}

func (bs *BitSet) SetAll() {
	for i := range bs.bitSet {
		bs.bitSet[i] |= 0xffffffffffffffff
	}
}

func (bs *BitSet) CountBitsOn() int {
	count := 0
	for _, value := range bs.bitSet {
		count += countBits(value)
	}

	return count
}

func (bs *BitSet) CountBitsOff() int {
	return bs.numOfBits - bs.CountBitsOn()
}

func (bs *BitSet) Copy() *BitSet {
	newBitSet := make([]uint64, len(bs.bitSet))
	copy(newBitSet, bs.bitSet)

	return &BitSet{bitSet: newBitSet}
}

func (bs *BitSet) String() string {
	var result strings.Builder
	for _, value := range bs.bitSet {
		result.WriteString(fmt.Sprintf("%064b ", value))
	}

	return result.String()
}

func (bs *BitSet) Equals(other *BitSet) bool {
	if len(bs.bitSet) != len(other.bitSet) {
		return false
	}

	for i, value := range bs.bitSet {
		if value != other.bitSet[i] {
			return false
		}
	}

	return true
}

func getIndex(idx uint64) uint64 {
	return idx / 64
}

func getBit(idx uint64) uint64 {
	return 1 << (idx % 64)
}

func countBits(value uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int((value >> uint(i)) & 1)
	}

	return count
}
