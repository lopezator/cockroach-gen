// Code generated by execgen; DO NOT EDIT.
// Copyright 2019 The Cockroach Authors.
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"bytes"
	"fmt"
	"math"
	"time"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/execerror"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

// vecComparator is a helper for the ordered synchronizer. It stores multiple
// column vectors of a single type and facilitates comparing values between
// them.
type vecComparator interface {
	// compare compares values from two vectors. vecIdx is the index of the vector
	// and valIdx is the index of the value in that vector to compare. Returns -1,
	// 0, or 1.
	compare(vecIdx1, vecIdx2 int, valIdx1, valIdx2 uint16) int

	// set sets the value of the vector at dstVecIdx at index dstValIdx to the value
	// at the vector at srcVecIdx at index srcValIdx.
	set(srcVecIdx, dstVecIdx int, srcValIdx, dstValIdx uint16)

	// setVec updates the vector at idx.
	setVec(idx int, vec coldata.Vec)
}

type BoolVecComparator struct {
	vecs  [][]bool
	nulls []*coldata.Nulls
}

func (c *BoolVecComparator) compare(vecIdx1, vecIdx2 int, valIdx1, valIdx2 uint16) int {
	n1 := c.nulls[vecIdx1].MaybeHasNulls() && c.nulls[vecIdx1].NullAt(valIdx1)
	n2 := c.nulls[vecIdx2].MaybeHasNulls() && c.nulls[vecIdx2].NullAt(valIdx2)
	if n1 && n2 {
		return 0
	} else if n1 {
		return -1
	} else if n2 {
		return 1
	}
	left := c.vecs[vecIdx1][int(valIdx1)]
	right := c.vecs[vecIdx2][int(valIdx2)]
	var cmp int

	if !left && right {
		cmp = -1
	} else if left && !right {
		cmp = 1
	} else {
		cmp = 0
	}

	return cmp
}

func (c *BoolVecComparator) setVec(idx int, vec coldata.Vec) {
	c.vecs[idx] = vec.Bool()
	c.nulls[idx] = vec.Nulls()
}

func (c *BoolVecComparator) set(srcVecIdx, dstVecIdx int, srcIdx, dstIdx uint16) {
	if c.nulls[srcVecIdx].MaybeHasNulls() && c.nulls[srcVecIdx].NullAt(srcIdx) {
		c.nulls[dstVecIdx].SetNull(dstIdx)
	} else {
		c.nulls[dstVecIdx].UnsetNull(dstIdx)
		v := c.vecs[srcVecIdx][int(srcIdx)]
		c.vecs[dstVecIdx][int(dstIdx)] = v
	}
}

type BytesVecComparator struct {
	vecs  []*coldata.Bytes
	nulls []*coldata.Nulls
}

func (c *BytesVecComparator) compare(vecIdx1, vecIdx2 int, valIdx1, valIdx2 uint16) int {
	n1 := c.nulls[vecIdx1].MaybeHasNulls() && c.nulls[vecIdx1].NullAt(valIdx1)
	n2 := c.nulls[vecIdx2].MaybeHasNulls() && c.nulls[vecIdx2].NullAt(valIdx2)
	if n1 && n2 {
		return 0
	} else if n1 {
		return -1
	} else if n2 {
		return 1
	}
	left := c.vecs[vecIdx1].Get(int(valIdx1))
	right := c.vecs[vecIdx2].Get(int(valIdx2))
	var cmp int
	cmp = bytes.Compare(left, right)
	return cmp
}

func (c *BytesVecComparator) setVec(idx int, vec coldata.Vec) {
	c.vecs[idx] = vec.Bytes()
	c.nulls[idx] = vec.Nulls()
}

func (c *BytesVecComparator) set(srcVecIdx, dstVecIdx int, srcIdx, dstIdx uint16) {
	if c.nulls[srcVecIdx].MaybeHasNulls() && c.nulls[srcVecIdx].NullAt(srcIdx) {
		c.nulls[dstVecIdx].SetNull(dstIdx)
	} else {
		c.nulls[dstVecIdx].UnsetNull(dstIdx)
		// Since flat Bytes cannot be set at arbitrary indices (data needs to be
		// moved around), we use CopySlice to accept the performance hit.
		// Specifically, this is a performance hit because we are overwriting the
		// variable number of bytes in `dstVecIdx`, so we will have to either shift
		// the bytes after that element left or right, depending on how long the
		// source bytes slice is. Refer to the CopySlice comment for an example.
		c.vecs[dstVecIdx].CopySlice(c.vecs[srcVecIdx], int(dstIdx), int(srcIdx), int(srcIdx+1))
	}
}

type DecimalVecComparator struct {
	vecs  [][]apd.Decimal
	nulls []*coldata.Nulls
}

func (c *DecimalVecComparator) compare(vecIdx1, vecIdx2 int, valIdx1, valIdx2 uint16) int {
	n1 := c.nulls[vecIdx1].MaybeHasNulls() && c.nulls[vecIdx1].NullAt(valIdx1)
	n2 := c.nulls[vecIdx2].MaybeHasNulls() && c.nulls[vecIdx2].NullAt(valIdx2)
	if n1 && n2 {
		return 0
	} else if n1 {
		return -1
	} else if n2 {
		return 1
	}
	left := c.vecs[vecIdx1][int(valIdx1)]
	right := c.vecs[vecIdx2][int(valIdx2)]
	var cmp int
	cmp = tree.CompareDecimals(&left, &right)
	return cmp
}

func (c *DecimalVecComparator) setVec(idx int, vec coldata.Vec) {
	c.vecs[idx] = vec.Decimal()
	c.nulls[idx] = vec.Nulls()
}

func (c *DecimalVecComparator) set(srcVecIdx, dstVecIdx int, srcIdx, dstIdx uint16) {
	if c.nulls[srcVecIdx].MaybeHasNulls() && c.nulls[srcVecIdx].NullAt(srcIdx) {
		c.nulls[dstVecIdx].SetNull(dstIdx)
	} else {
		c.nulls[dstVecIdx].UnsetNull(dstIdx)
		v := c.vecs[srcVecIdx][int(srcIdx)]
		c.vecs[dstVecIdx][int(dstIdx)].Set(&v)
	}
}

type Int16VecComparator struct {
	vecs  [][]int16
	nulls []*coldata.Nulls
}

func (c *Int16VecComparator) compare(vecIdx1, vecIdx2 int, valIdx1, valIdx2 uint16) int {
	n1 := c.nulls[vecIdx1].MaybeHasNulls() && c.nulls[vecIdx1].NullAt(valIdx1)
	n2 := c.nulls[vecIdx2].MaybeHasNulls() && c.nulls[vecIdx2].NullAt(valIdx2)
	if n1 && n2 {
		return 0
	} else if n1 {
		return -1
	} else if n2 {
		return 1
	}
	left := c.vecs[vecIdx1][int(valIdx1)]
	right := c.vecs[vecIdx2][int(valIdx2)]
	var cmp int

	{
		a, b := int64(left), int64(right)
		if a < b {
			cmp = -1
		} else if a > b {
			cmp = 1
		} else {
			cmp = 0
		}
	}

	return cmp
}

func (c *Int16VecComparator) setVec(idx int, vec coldata.Vec) {
	c.vecs[idx] = vec.Int16()
	c.nulls[idx] = vec.Nulls()
}

func (c *Int16VecComparator) set(srcVecIdx, dstVecIdx int, srcIdx, dstIdx uint16) {
	if c.nulls[srcVecIdx].MaybeHasNulls() && c.nulls[srcVecIdx].NullAt(srcIdx) {
		c.nulls[dstVecIdx].SetNull(dstIdx)
	} else {
		c.nulls[dstVecIdx].UnsetNull(dstIdx)
		v := c.vecs[srcVecIdx][int(srcIdx)]
		c.vecs[dstVecIdx][int(dstIdx)] = v
	}
}

type Int32VecComparator struct {
	vecs  [][]int32
	nulls []*coldata.Nulls
}

func (c *Int32VecComparator) compare(vecIdx1, vecIdx2 int, valIdx1, valIdx2 uint16) int {
	n1 := c.nulls[vecIdx1].MaybeHasNulls() && c.nulls[vecIdx1].NullAt(valIdx1)
	n2 := c.nulls[vecIdx2].MaybeHasNulls() && c.nulls[vecIdx2].NullAt(valIdx2)
	if n1 && n2 {
		return 0
	} else if n1 {
		return -1
	} else if n2 {
		return 1
	}
	left := c.vecs[vecIdx1][int(valIdx1)]
	right := c.vecs[vecIdx2][int(valIdx2)]
	var cmp int

	{
		a, b := int64(left), int64(right)
		if a < b {
			cmp = -1
		} else if a > b {
			cmp = 1
		} else {
			cmp = 0
		}
	}

	return cmp
}

func (c *Int32VecComparator) setVec(idx int, vec coldata.Vec) {
	c.vecs[idx] = vec.Int32()
	c.nulls[idx] = vec.Nulls()
}

func (c *Int32VecComparator) set(srcVecIdx, dstVecIdx int, srcIdx, dstIdx uint16) {
	if c.nulls[srcVecIdx].MaybeHasNulls() && c.nulls[srcVecIdx].NullAt(srcIdx) {
		c.nulls[dstVecIdx].SetNull(dstIdx)
	} else {
		c.nulls[dstVecIdx].UnsetNull(dstIdx)
		v := c.vecs[srcVecIdx][int(srcIdx)]
		c.vecs[dstVecIdx][int(dstIdx)] = v
	}
}

type Int64VecComparator struct {
	vecs  [][]int64
	nulls []*coldata.Nulls
}

func (c *Int64VecComparator) compare(vecIdx1, vecIdx2 int, valIdx1, valIdx2 uint16) int {
	n1 := c.nulls[vecIdx1].MaybeHasNulls() && c.nulls[vecIdx1].NullAt(valIdx1)
	n2 := c.nulls[vecIdx2].MaybeHasNulls() && c.nulls[vecIdx2].NullAt(valIdx2)
	if n1 && n2 {
		return 0
	} else if n1 {
		return -1
	} else if n2 {
		return 1
	}
	left := c.vecs[vecIdx1][int(valIdx1)]
	right := c.vecs[vecIdx2][int(valIdx2)]
	var cmp int

	{
		a, b := int64(left), int64(right)
		if a < b {
			cmp = -1
		} else if a > b {
			cmp = 1
		} else {
			cmp = 0
		}
	}

	return cmp
}

func (c *Int64VecComparator) setVec(idx int, vec coldata.Vec) {
	c.vecs[idx] = vec.Int64()
	c.nulls[idx] = vec.Nulls()
}

func (c *Int64VecComparator) set(srcVecIdx, dstVecIdx int, srcIdx, dstIdx uint16) {
	if c.nulls[srcVecIdx].MaybeHasNulls() && c.nulls[srcVecIdx].NullAt(srcIdx) {
		c.nulls[dstVecIdx].SetNull(dstIdx)
	} else {
		c.nulls[dstVecIdx].UnsetNull(dstIdx)
		v := c.vecs[srcVecIdx][int(srcIdx)]
		c.vecs[dstVecIdx][int(dstIdx)] = v
	}
}

type Float64VecComparator struct {
	vecs  [][]float64
	nulls []*coldata.Nulls
}

func (c *Float64VecComparator) compare(vecIdx1, vecIdx2 int, valIdx1, valIdx2 uint16) int {
	n1 := c.nulls[vecIdx1].MaybeHasNulls() && c.nulls[vecIdx1].NullAt(valIdx1)
	n2 := c.nulls[vecIdx2].MaybeHasNulls() && c.nulls[vecIdx2].NullAt(valIdx2)
	if n1 && n2 {
		return 0
	} else if n1 {
		return -1
	} else if n2 {
		return 1
	}
	left := c.vecs[vecIdx1][int(valIdx1)]
	right := c.vecs[vecIdx2][int(valIdx2)]
	var cmp int

	{
		a, b := float64(left), float64(right)
		if a < b {
			cmp = -1
		} else if a > b {
			cmp = 1
		} else if a == b {
			cmp = 0
		} else if math.IsNaN(a) {
			if math.IsNaN(b) {
				cmp = 0
			} else {
				cmp = -1
			}
		} else {
			cmp = 1
		}
	}

	return cmp
}

func (c *Float64VecComparator) setVec(idx int, vec coldata.Vec) {
	c.vecs[idx] = vec.Float64()
	c.nulls[idx] = vec.Nulls()
}

func (c *Float64VecComparator) set(srcVecIdx, dstVecIdx int, srcIdx, dstIdx uint16) {
	if c.nulls[srcVecIdx].MaybeHasNulls() && c.nulls[srcVecIdx].NullAt(srcIdx) {
		c.nulls[dstVecIdx].SetNull(dstIdx)
	} else {
		c.nulls[dstVecIdx].UnsetNull(dstIdx)
		v := c.vecs[srcVecIdx][int(srcIdx)]
		c.vecs[dstVecIdx][int(dstIdx)] = v
	}
}

type TimestampVecComparator struct {
	vecs  [][]time.Time
	nulls []*coldata.Nulls
}

func (c *TimestampVecComparator) compare(vecIdx1, vecIdx2 int, valIdx1, valIdx2 uint16) int {
	n1 := c.nulls[vecIdx1].MaybeHasNulls() && c.nulls[vecIdx1].NullAt(valIdx1)
	n2 := c.nulls[vecIdx2].MaybeHasNulls() && c.nulls[vecIdx2].NullAt(valIdx2)
	if n1 && n2 {
		return 0
	} else if n1 {
		return -1
	} else if n2 {
		return 1
	}
	left := c.vecs[vecIdx1][int(valIdx1)]
	right := c.vecs[vecIdx2][int(valIdx2)]
	var cmp int

	if left.Before(right) {
		cmp = -1
	} else if right.Before(left) {
		cmp = 1
	} else {
		cmp = 0
	}
	return cmp
}

func (c *TimestampVecComparator) setVec(idx int, vec coldata.Vec) {
	c.vecs[idx] = vec.Timestamp()
	c.nulls[idx] = vec.Nulls()
}

func (c *TimestampVecComparator) set(srcVecIdx, dstVecIdx int, srcIdx, dstIdx uint16) {
	if c.nulls[srcVecIdx].MaybeHasNulls() && c.nulls[srcVecIdx].NullAt(srcIdx) {
		c.nulls[dstVecIdx].SetNull(dstIdx)
	} else {
		c.nulls[dstVecIdx].UnsetNull(dstIdx)
		v := c.vecs[srcVecIdx][int(srcIdx)]
		c.vecs[dstVecIdx][int(dstIdx)] = v
	}
}

func GetVecComparator(t coltypes.T, numVecs int) vecComparator {
	switch t {
	case coltypes.Bool:
		return &BoolVecComparator{
			vecs:  make([][]bool, numVecs),
			nulls: make([]*coldata.Nulls, numVecs),
		}
	case coltypes.Bytes:
		return &BytesVecComparator{
			vecs:  make([]*coldata.Bytes, numVecs),
			nulls: make([]*coldata.Nulls, numVecs),
		}
	case coltypes.Decimal:
		return &DecimalVecComparator{
			vecs:  make([][]apd.Decimal, numVecs),
			nulls: make([]*coldata.Nulls, numVecs),
		}
	case coltypes.Int16:
		return &Int16VecComparator{
			vecs:  make([][]int16, numVecs),
			nulls: make([]*coldata.Nulls, numVecs),
		}
	case coltypes.Int32:
		return &Int32VecComparator{
			vecs:  make([][]int32, numVecs),
			nulls: make([]*coldata.Nulls, numVecs),
		}
	case coltypes.Int64:
		return &Int64VecComparator{
			vecs:  make([][]int64, numVecs),
			nulls: make([]*coldata.Nulls, numVecs),
		}
	case coltypes.Float64:
		return &Float64VecComparator{
			vecs:  make([][]float64, numVecs),
			nulls: make([]*coldata.Nulls, numVecs),
		}
	case coltypes.Timestamp:
		return &TimestampVecComparator{
			vecs:  make([][]time.Time, numVecs),
			nulls: make([]*coldata.Nulls, numVecs),
		}
	}
	execerror.VectorizedInternalPanic(fmt.Sprintf("unhandled type %v", t))
	// This code is unreachable, but the compiler cannot infer that.
	return nil
}
