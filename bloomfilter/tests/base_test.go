package tests

import (
	"context"
	"fmt"
	"go-leetcode/bloomfilter/base"
	"gotest.tools/assert"
	"math"
	"sync"
	"testing"
)

func TestPow(t *testing.T) {
	fmt.Println(math.Pow(2, 3))
}

func TestEstimateParams(t *testing.T) {
	//https://hur.st/bloomfilter/
	tests := []struct {
		n uint64
		p float64
		m uint64
		k uint64
	}{
		{n: 4000, p: 0.0000001, m: 134191, k: 23},
		{n: 100, p: 0.02, m: 815, k: 6},
		{n: 100000000, p: 0.0001, m: 1917011676, k: 13},
	}
	for row, test := range tests {
		inf := base.BFInfo{
			N:    test.n,
			P:    test.p,
			Name: "fuck",
		}
		inf.EstimateParams()
		assert.Equal(t, inf.M, test.m, "should equal m, row:%d", row)
		assert.Equal(t, inf.K, test.k, "should equal m, row:%d", row)
	}
	//	fmt.Println("123")
	//	var b bool
	//	if b {
	//		goto prinfs
	//		fmt.Println("exist")
	//	}
	//prinfs:
	//	fmt.Println("goto 1234")
}

func TestCalculateParts(t *testing.T) {
	//https://hur.st/bloomfilter/
	tests := []struct {
		name        string
		m           uint64
		parts       int
		lastPart    string
		lastPartMax uint32
	}{
		{name: "bf1", m: 1024, parts: 1, lastPart: "bf1:0", lastPartMax: 1024},
		{name: "bf2", m: base.PartBitCount*5 + 1024, parts: 6, lastPart: "bf2:5", lastPartMax: 1024},
	}

	for row, test := range tests {
		bloom := &base.BFInfo{Name: test.name, M: test.m}
		bloom.CalculateParts()
		assert.Equal(t, test.parts, len(bloom.Parts), "part size should be valid, row: %d", row)
		assert.Equal(t, test.lastPart, bloom.Parts[test.parts-1].Name, "last part name should be valid, row: %d", row)
		assert.Equal(t, test.lastPartMax, bloom.Parts[test.parts-1].Max, "last part max should be valid, row: %d", row)
	}
}

func truePositiveTest(bf base.BloomFilter, sampleSize int) bool {

	var added = sampleSize
	group := sync.WaitGroup{}
	i := 0
	for i != added {
		uid := fmt.Sprintf("%d", i)
		group.Add(1)
		go func() {
			bf.Add(context.Background(), []byte(uid))
			group.Done()
		}()
		i++
	}
	group.Wait()
	//check
	i = 0
	result := make(chan bool, 100000)
	checkGroup := sync.WaitGroup{}
	var flag bool = true

	go func() {
		for f := range result {
			if !f {
				print(f)
			}
			flag = f && flag
		}
	}()

	for i != added {
		uid := fmt.Sprintf("%d", i)
		checkGroup.Add(1)
		go func() {
			ex, _ := bf.Exists(context.Background(), []byte(uid))
			result <- ex
			checkGroup.Done()
		}()
		i++
	}

	checkGroup.Wait()
	return flag
}

func falseNegativeTest(bf base.BloomFilter, sampleSize int) bool {
	var added = sampleSize
	i := 0
	group := sync.WaitGroup{}
	for i != added {
		uid := fmt.Sprintf("%d", i)
		group.Add(1)
		go func() {
			bf.Add(context.Background(), []byte(uid))
			group.Done()
		}()
		i++
	}
	group.Wait()
	//check
	i = added
	result := make(chan bool, 1000)
	checkGroup := sync.WaitGroup{}
	var flag bool = false

	go func() {
		for f := range result {
			flag = f || flag
		}
	}()
	for i != 2*added {
		uid := fmt.Sprintf("%d", i)
		checkGroup.Add(1)
		go func() {
			ex, _ := bf.Exists(context.Background(), []byte(uid))
			result <- ex
			checkGroup.Done()
		}()
		i++
	}
	checkGroup.Wait()
	return flag
}

func testSerialTruePostive(bf base.BloomFilter, sampleSize int) bool {
	var added = sampleSize
	i := 0
	for i != added {
		uid := fmt.Sprintf("%d", i)
		bf.Add(context.Background(), []byte(uid))
		i++
	}

	//check
	i = 0
	for i != added {
		uid := fmt.Sprintf("%d", i)
		ex, _ := bf.Exists(context.Background(), []byte(uid))
		if !ex {
			return false
		}
		i++

	}
	return true
}

func testSerialFalseNegative(bf base.BloomFilter, sampleSize int) bool {
	var added = sampleSize
	i := 0
	for i != added {
		uid := fmt.Sprintf("%d", i)
		bf.Add(context.Background(), []byte(uid))
		i++
	}

	//check
	i = added
	for i != 2*added {
		uid := fmt.Sprintf("%d", i)
		ex, _ := bf.Exists(context.Background(), []byte(uid))
		if ex {
			return true
		}
		i++
	}
	return false
}

func testSerialBatchAdds(uids [][]byte, bf base.BloomFilter, batchsize int) error {

	added := len(uids)
	j := 0
	bs := batchsize
	for j < added {
		end := j + bs
		if j+bs > added {
			end = added
		}
		err := bf.Adds(context.Background(), uids[j:end])
		if err != nil {
			return err
		}
		j += bs
	}
	return nil
}

func testSerialBatchGets(uids [][]byte, bf base.BloomFilter, batchsize int) ([]bool, error) {
	added := len(uids)
	k := 0
	kbs := batchsize
	res := make([]bool, 0)
	for k < added {
		end := k + kbs
		if k+kbs > added {
			end = added
		}
		bools, err := bf.BatchExists(context.Background(), uids[k:end])
		if err != nil {
			return nil, err
		}
		k += kbs
		res = append(res, bools...)
	}
	return res, nil
}
