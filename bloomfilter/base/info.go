package base

import (
	"fmt"
	"github.com/twmb/murmur3"
	"math"
)

const PartBitCount uint64 = 1 << 32

type PartInfo struct {
	Name  string `json:"name"`
	Max   uint32 `json:"max"`
	Index int    `json:"index"`
}

type Location struct {
	Name   string
	Offset uint64
	Index  int
}

type BFInfo struct {
	N     uint64      `json:"n"`     // 总数
	P     float64     `json:"p"`     // 假阳率
	Name  string      `json:"name"`  // 名字
	M     uint64      `json:"m"`     // bitmap 大小
	K     uint64      `json:"k"`     // hash函数
	Parts []*PartInfo `json:"parts"` // 分段，redis string 最大512MB 2^32 -1
}

func NewBFInfo(name string, n uint64, p float64) *BFInfo {
	inf := &BFInfo{
		N:     n,
		P:     p,
		Name:  name,
		Parts: nil,
	}
	inf.EstimateParams()
	inf.CalculateParts()
	return inf
}

func (i *BFInfo) EstimateParams() {
	// https://en.wikipedia.org/wiki/Bloom_filter
	// https://hur.st/bloomfilter/
	m := math.Ceil(float64(i.N) * math.Log(i.P) / math.Log(1.0/math.Pow(2.0, math.Ln2)))
	k := math.Ln2*m/float64(i.N) + 0.5
	i.M = uint64(m)
	i.K = uint64(k)
}

func (i *BFInfo) CalculateParts() {
	cnt := i.M/PartBitCount + 1
	i.Parts = make([]*PartInfo, cnt)
	for p := range i.Parts {
		p := p
		i.Parts[p] = &PartInfo{
			Name:  fmt.Sprintf("%s:%d", i.Name, p),
			Max:   math.MaxUint32,
			Index: p,
		}
	}
	i.Parts[cnt-1].Max = uint32(i.M % PartBitCount)
}

// rejectionSample normalize the hash value or throw it away
func (b *BFInfo) rejectionSample(random uint64, m uint64) (uint64, bool) {
	if random > (math.MaxUint64-math.MaxUint64%m) || random == 0 {
		// if random fall in range above the hash result can not be evenly distributed
		// so skip this result
		return 0, false
	}
	return random % m, true
}

// hashes get k hash results of value
func (b *BFInfo) Hashes(value []byte) (hashes []uint64) {
	hashes = make([]uint64, b.K)
	var seed uint64

	var i uint64
	for i < b.K {
		seed = murmur3.SeedSum64(seed, value)
		hash, valid := b.rejectionSample(seed, b.M)
		if valid {
			hashes[i] = hash
			i++
		}
	}
	return hashes
}

func (b *BFInfo) Locations(value []byte) []*Location {
	hashes := b.Hashes(value)
	locs := make([]*Location, len(hashes))
	for i := range locs {
		p := hashes[i] / PartBitCount
		locs[i] = &Location{
			Name:   b.Parts[p].Name,
			Offset: hashes[i] % (PartBitCount - 1),
			Index:  int(p),
		}

	}
	return locs
}
