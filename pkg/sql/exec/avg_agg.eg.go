// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

//

package exec

import (
	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/types"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/pkg/errors"
)

//

func newAvgAgg(t types.T) (aggregateFunc, error) {
	switch t {
	//
	case types.Decimal:
		return &avgDecimalAgg{}, nil
	//
	case types.Float32:
		return &avgFloat32Agg{}, nil
	//
	case types.Float64:
		return &avgFloat64Agg{}, nil
	//
	default:
		return nil, errors.Errorf("unsupported avg agg type %s", t)
	}
}

//

type avgDecimalAgg struct {
	done bool

	groups  []bool
	scratch struct {
		curIdx int
		// groupSums[i] keeps track of the sum of elements belonging to the ith
		// group.
		groupSums []apd.Decimal
		// groupCounts[i] keeps track of the number of elements that we've seen
		// belonging to the ith group.
		groupCounts []int64
		// vec points to the output vector.
		vec []apd.Decimal
	}
}

var _ aggregateFunc = &avgDecimalAgg{}

func (a *avgDecimalAgg) Init(groups []bool, v ColVec) {
	a.groups = groups
	a.scratch.vec = v.Decimal()
	a.scratch.groupSums = make([]apd.Decimal, len(a.scratch.vec))
	a.scratch.groupCounts = make([]int64, len(a.scratch.vec))
	a.Reset()
}

func (a *avgDecimalAgg) Reset() {
	copy(a.scratch.groupSums, zeroDecimalBatch)
	copy(a.scratch.groupCounts, zeroInt64Batch)
	copy(a.scratch.vec, zeroDecimalBatch)
	a.scratch.curIdx = -1
}

func (a *avgDecimalAgg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *avgDecimalAgg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		copy(a.scratch.groupSums[idx+1:], zeroDecimalBatch)
		copy(a.scratch.groupCounts[idx+1:], zeroInt64Batch)
		// TODO(asubiotto): We might not have to zero a.scratch.vec since we
		// overwrite with an independent value.
		copy(a.scratch.vec[idx+1:], zeroDecimalBatch)
	}
}

func (a *avgDecimalAgg) Compute(b ColBatch, inputIdxs []uint32) {
	if a.done {
		return
	}
	inputLen := b.Length()
	if inputLen == 0 {
		// The aggregation is finished. Flush the last value.
		if a.scratch.curIdx >= 0 {
			a.scratch.vec[a.scratch.curIdx].SetInt64(a.scratch.groupCounts[a.scratch.curIdx])
			if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[a.scratch.curIdx], &a.scratch.groupSums[a.scratch.curIdx], &a.scratch.vec[a.scratch.curIdx]); err != nil {
				panic(err)
			}
		}
		a.scratch.curIdx++
		a.done = true
		return
	}
	col, sel := b.ColVec(int(inputIdxs[0])).Decimal(), b.Selection()
	if sel != nil {
		sel = sel[:inputLen]
		for _, i := range sel {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			if _, err := tree.DecimalCtx.Add(&a.scratch.groupSums[a.scratch.curIdx], &a.scratch.groupSums[a.scratch.curIdx], &col[i]); err != nil {
				panic(err)
			}
			a.scratch.groupCounts[a.scratch.curIdx]++
		}
	} else {
		col = col[:inputLen]
		for i := range col {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			if _, err := tree.DecimalCtx.Add(&a.scratch.groupSums[a.scratch.curIdx], &a.scratch.groupSums[a.scratch.curIdx], &col[i]); err != nil {
				panic(err)
			}
			a.scratch.groupCounts[a.scratch.curIdx]++
		}
	}

	for i := 0; i < a.scratch.curIdx; i++ {
		a.scratch.vec[i].SetInt64(a.scratch.groupCounts[i])
		if _, err := tree.DecimalCtx.Quo(&a.scratch.vec[i], &a.scratch.groupSums[i], &a.scratch.vec[i]); err != nil {
			panic(err)
		}
	}
}

//

type avgFloat32Agg struct {
	done bool

	groups  []bool
	scratch struct {
		curIdx int
		// groupSums[i] keeps track of the sum of elements belonging to the ith
		// group.
		groupSums []float32
		// groupCounts[i] keeps track of the number of elements that we've seen
		// belonging to the ith group.
		groupCounts []int64
		// vec points to the output vector.
		vec []float32
	}
}

var _ aggregateFunc = &avgFloat32Agg{}

func (a *avgFloat32Agg) Init(groups []bool, v ColVec) {
	a.groups = groups
	a.scratch.vec = v.Float32()
	a.scratch.groupSums = make([]float32, len(a.scratch.vec))
	a.scratch.groupCounts = make([]int64, len(a.scratch.vec))
	a.Reset()
}

func (a *avgFloat32Agg) Reset() {
	copy(a.scratch.groupSums, zeroFloat32Batch)
	copy(a.scratch.groupCounts, zeroInt64Batch)
	copy(a.scratch.vec, zeroFloat32Batch)
	a.scratch.curIdx = -1
}

func (a *avgFloat32Agg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *avgFloat32Agg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		copy(a.scratch.groupSums[idx+1:], zeroFloat32Batch)
		copy(a.scratch.groupCounts[idx+1:], zeroInt64Batch)
		// TODO(asubiotto): We might not have to zero a.scratch.vec since we
		// overwrite with an independent value.
		copy(a.scratch.vec[idx+1:], zeroFloat32Batch)
	}
}

func (a *avgFloat32Agg) Compute(b ColBatch, inputIdxs []uint32) {
	if a.done {
		return
	}
	inputLen := b.Length()
	if inputLen == 0 {
		// The aggregation is finished. Flush the last value.
		if a.scratch.curIdx >= 0 {
			a.scratch.vec[a.scratch.curIdx] = a.scratch.groupSums[a.scratch.curIdx] / float32(a.scratch.groupCounts[a.scratch.curIdx])
		}
		a.scratch.curIdx++
		a.done = true
		return
	}
	col, sel := b.ColVec(int(inputIdxs[0])).Float32(), b.Selection()
	if sel != nil {
		sel = sel[:inputLen]
		for _, i := range sel {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.groupSums[a.scratch.curIdx] = a.scratch.groupSums[a.scratch.curIdx] + col[i]
			a.scratch.groupCounts[a.scratch.curIdx]++
		}
	} else {
		col = col[:inputLen]
		for i := range col {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.groupSums[a.scratch.curIdx] = a.scratch.groupSums[a.scratch.curIdx] + col[i]
			a.scratch.groupCounts[a.scratch.curIdx]++
		}
	}

	for i := 0; i < a.scratch.curIdx; i++ {
		a.scratch.vec[i] = a.scratch.groupSums[i] / float32(a.scratch.groupCounts[i])
	}
}

//

type avgFloat64Agg struct {
	done bool

	groups  []bool
	scratch struct {
		curIdx int
		// groupSums[i] keeps track of the sum of elements belonging to the ith
		// group.
		groupSums []float64
		// groupCounts[i] keeps track of the number of elements that we've seen
		// belonging to the ith group.
		groupCounts []int64
		// vec points to the output vector.
		vec []float64
	}
}

var _ aggregateFunc = &avgFloat64Agg{}

func (a *avgFloat64Agg) Init(groups []bool, v ColVec) {
	a.groups = groups
	a.scratch.vec = v.Float64()
	a.scratch.groupSums = make([]float64, len(a.scratch.vec))
	a.scratch.groupCounts = make([]int64, len(a.scratch.vec))
	a.Reset()
}

func (a *avgFloat64Agg) Reset() {
	copy(a.scratch.groupSums, zeroFloat64Batch)
	copy(a.scratch.groupCounts, zeroInt64Batch)
	copy(a.scratch.vec, zeroFloat64Batch)
	a.scratch.curIdx = -1
}

func (a *avgFloat64Agg) CurrentOutputIndex() int {
	return a.scratch.curIdx
}

func (a *avgFloat64Agg) SetOutputIndex(idx int) {
	if a.scratch.curIdx != -1 {
		a.scratch.curIdx = idx
		copy(a.scratch.groupSums[idx+1:], zeroFloat64Batch)
		copy(a.scratch.groupCounts[idx+1:], zeroInt64Batch)
		// TODO(asubiotto): We might not have to zero a.scratch.vec since we
		// overwrite with an independent value.
		copy(a.scratch.vec[idx+1:], zeroFloat64Batch)
	}
}

func (a *avgFloat64Agg) Compute(b ColBatch, inputIdxs []uint32) {
	if a.done {
		return
	}
	inputLen := b.Length()
	if inputLen == 0 {
		// The aggregation is finished. Flush the last value.
		if a.scratch.curIdx >= 0 {
			a.scratch.vec[a.scratch.curIdx] = a.scratch.groupSums[a.scratch.curIdx] / float64(a.scratch.groupCounts[a.scratch.curIdx])
		}
		a.scratch.curIdx++
		a.done = true
		return
	}
	col, sel := b.ColVec(int(inputIdxs[0])).Float64(), b.Selection()
	if sel != nil {
		sel = sel[:inputLen]
		for _, i := range sel {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.groupSums[a.scratch.curIdx] = a.scratch.groupSums[a.scratch.curIdx] + col[i]
			a.scratch.groupCounts[a.scratch.curIdx]++
		}
	} else {
		col = col[:inputLen]
		for i := range col {
			x := 0
			if a.groups[i] {
				x = 1
			}
			a.scratch.curIdx += x
			a.scratch.groupSums[a.scratch.curIdx] = a.scratch.groupSums[a.scratch.curIdx] + col[i]
			a.scratch.groupCounts[a.scratch.curIdx]++
		}
	}

	for i := 0; i < a.scratch.curIdx; i++ {
		a.scratch.vec[i] = a.scratch.groupSums[i] / float64(a.scratch.groupCounts[i])
	}
}

//