/**
 * @Author raven
 * @Description
 * @Date 2023/1/10
 **/
package bfbyte

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RoaringBitmap/roaring/roaring64"
	"go-leetcode/bloomfilter/base"
	"sync"
)

type Config struct {
	base.BFInfo
}

func NewConfig(info base.BFInfo) *Config {
	return &Config{info}
}

func NewDefaultConfig() *Config {
	return NewConfig(base.BFInfo{
		N:    1000000,
		P:    0.0001,
		Name: "bf_byte",
	})
}

// GetOrBuild 通过反序列化 创建一个布隆过滤器
func (c *Config) GetOrBuild(bmsByte []byte) (*BFByte, error) {
	// 新建一个
	if len(bmsByte) == 0 {
		b := &BFByte{c, make(map[int]*roaring64.Bitmap), sync.RWMutex{}}
		b.BFInfo.EstimateParams()
		b.BFInfo.CalculateParts()
		b.restoreBloom()
		return b, nil
	}
	// 反序列化
	var bmsData map[int]string

	err := json.Unmarshal(bmsByte, &bmsData)
	// 反序列化失败
	if err != nil {
		return nil, err
	}
	bms := make(map[int]*roaring64.Bitmap, len(bmsData))

	for k, v := range bmsData {
		fmt.Println([]byte(v))
		fmt.Println(v)
		bitMap := roaring64.New()
		err = bitMap.UnmarshalBinary([]byte(v))
		if err != nil {
			return nil, err
		}
		bms[k] = bitMap
	}
	b := &BFByte{c, bms, sync.RWMutex{}}
	b.BFInfo.EstimateParams()
	b.BFInfo.CalculateParts()
	return b, nil
}

type BFByte struct {
	*Config
	bms        map[int]*roaring64.Bitmap
	storeMutex sync.RWMutex
}

func (b *BFByte) restoreBloom() {
	b.storeMutex.Lock()
	defer b.storeMutex.Unlock()
	for _, p := range b.Parts {
		b.bms[p.Index] = roaring64.New()
	}
	return
}

// Marshal 序列化
func (b *BFByte) Marshal() ([]byte, error) {
	b.storeMutex.Lock()
	defer b.storeMutex.Unlock()
	bmsData := make(map[int][]byte, len(b.bms))
	for k, v := range b.bms {
		bmsByte, err := v.MarshalBinary()
		if err != nil {
			return nil, err
		}
		bmsData[k] = bmsByte
	}
	return json.Marshal(bmsData)
}

func (b *BFByte) Add(ctx context.Context, value []byte) error {
	locs := b.BFInfo.Locations(value)
	b.storeMutex.Lock()
	defer b.storeMutex.Unlock()
	for _, loc := range locs {
		b.bms[loc.Index].Add(loc.Offset)
	}
	return nil
}

func (b *BFByte) Adds(ctx context.Context, keys [][]byte) error {
	for _, v := range keys {
		b.Add(ctx, v)
	}
	return nil
}

func (b *BFByte) Exists(ctx context.Context, value []byte) (bool, error) {
	locs := b.BFInfo.Locations(value)
	b.storeMutex.RLock()
	defer b.storeMutex.RUnlock()
	for _, loc := range locs {
		if !b.bms[loc.Index].Contains(loc.Offset) {
			return false, nil
		}
	}
	return true, nil
}

func (b *BFByte) BatchExists(ctx context.Context, key [][]byte) ([]bool, error) {
	res := make([]bool, len(key))
	for i, v := range key {
		ok, err := b.Exists(ctx, v)
		if err != nil {
			return nil, err
		}
		res[i] = ok
	}
	return res, nil
}

func (b *BFByte) Clear() {
	b.storeMutex.Lock()
	defer b.storeMutex.Unlock()
	b.bms = make(map[int]*roaring64.Bitmap)
	for _, p := range b.Parts {
		b.bms[p.Index] = roaring64.New()
	}
}
