package bloomfilter

import (
	"hash/fnv"
	"sync/atomic"
)

// BloomFilter represents a concurrent Bloom filter.
type BloomFilter struct {
	m         int32    // Number of bits in the bitfield
	bitfield  []int32  // Bitfield to store bloom filter values
	hashFuncs int32    // Number of hash functions
	hashSeeds []uint32 // Seeds for hash functions
}

// NewBloomFilter creates a new BloomFilter with the specified parameters.
func NewBloomFilter(m, hashFuncs int32) *BloomFilter {
	bitfield := make([]int32, m)
	hashSeeds := generateHashSeeds(hashFuncs)
	return &BloomFilter{
		m:         m,
		bitfield:  bitfield,
		hashFuncs: hashFuncs,
		hashSeeds: hashSeeds,
	}
}

// Add inserts an element into the Bloom filter.
func (bf *BloomFilter) Add(element string) {
	for _, seed := range bf.hashSeeds {
		index := bf.hash(element, seed) % uint32(bf.m)
		atomic.StoreInt32(&bf.bitfield[index], 1)
	}
}

// Contains checks if an element is possibly in the Bloom filter.
func (bf *BloomFilter) Contains(element string) bool {
	for _, seed := range bf.hashSeeds {
		index := bf.hash(element, seed) % uint32(bf.m)
		if atomic.LoadInt32(&bf.bitfield[index]) != 1 {
			return false
		}
	}

	return true
}

// hash calculates the hash value for the given element and seed.
func (bf *BloomFilter) hash(element string, seed uint32) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(element))
	hash.Write([]byte{byte(seed)})
	return hash.Sum32()
}

// generateHashSeeds generates random hash seeds for the hash functions.
func generateHashSeeds(count int32) []uint32 {
	seeds := make([]uint32, count)
	for i := int32(0); i < count; i++ {
		seeds[i] = uint32(i + 1) // You can use a more sophisticated seed generation approach
	}
	return seeds
}
