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

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/execerror"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/execgen"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

// Use execgen package to remove unused import warning.
var _ interface{} = execgen.UNSAFEGET

// isBufferedGroupFinished checks to see whether or not the buffered group
// corresponding to input continues in batch.
func (o *mergeJoinBase) isBufferedGroupFinished(
	input *mergeJoinInput, batch coldata.Batch, rowIdx int,
) bool {
	if batch.Length() == 0 {
		return true
	}
	bufferedGroup := o.proberState.lBufferedGroup
	if input == &o.right {
		bufferedGroup = o.proberState.rBufferedGroup
	}
	lastBufferedTupleIdx := bufferedGroup.length - 1
	tupleToLookAtIdx := uint64(rowIdx)
	sel := batch.Selection()
	if sel != nil {
		tupleToLookAtIdx = uint64(sel[rowIdx])
	}

	// Check all equality columns in the first row of batch to make sure we're in
	// the same group.
	for _, colIdx := range input.eqCols[:len(input.eqCols)] {
		colTyp := input.sourceTypes[colIdx]

		switch colTyp {
		case coltypes.Bool:
			// We perform this null check on every equality column of the last
			// buffered tuple regardless of the join type since it is done only once
			// per batch. In some cases (like INNER JOIN, or LEFT OUTER JOIN with the
			// right side being an input) this check will always return false since
			// nulls couldn't be buffered up though.
			if bufferedGroup.ColVec(int(colIdx)).Nulls().NullAt64(uint64(lastBufferedTupleIdx)) {
				return true
			}
			bufferedCol := bufferedGroup.ColVec(int(colIdx)).Bool()
			prevVal := bufferedCol[int(lastBufferedTupleIdx)]
			var curVal bool
			if batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt64(tupleToLookAtIdx) {
				return true
			}
			col := batch.ColVec(int(colIdx)).Bool()
			curVal = col[int(tupleToLookAtIdx)]
			var match bool

			{
				var cmpResult int

				if !prevVal && curVal {
					cmpResult = -1
				} else if prevVal && !curVal {
					cmpResult = 1
				} else {
					cmpResult = 0
				}

				match = cmpResult == 0
			}

			if !match {
				return true
			}
		case coltypes.Bytes:
			// We perform this null check on every equality column of the last
			// buffered tuple regardless of the join type since it is done only once
			// per batch. In some cases (like INNER JOIN, or LEFT OUTER JOIN with the
			// right side being an input) this check will always return false since
			// nulls couldn't be buffered up though.
			if bufferedGroup.ColVec(int(colIdx)).Nulls().NullAt64(uint64(lastBufferedTupleIdx)) {
				return true
			}
			bufferedCol := bufferedGroup.ColVec(int(colIdx)).Bytes()
			prevVal := bufferedCol.Get(int(lastBufferedTupleIdx))
			var curVal []byte
			if batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt64(tupleToLookAtIdx) {
				return true
			}
			col := batch.ColVec(int(colIdx)).Bytes()
			curVal = col.Get(int(tupleToLookAtIdx))
			var match bool

			{
				var cmpResult int
				cmpResult = bytes.Compare(prevVal, curVal)
				match = cmpResult == 0
			}

			if !match {
				return true
			}
		case coltypes.Decimal:
			// We perform this null check on every equality column of the last
			// buffered tuple regardless of the join type since it is done only once
			// per batch. In some cases (like INNER JOIN, or LEFT OUTER JOIN with the
			// right side being an input) this check will always return false since
			// nulls couldn't be buffered up though.
			if bufferedGroup.ColVec(int(colIdx)).Nulls().NullAt64(uint64(lastBufferedTupleIdx)) {
				return true
			}
			bufferedCol := bufferedGroup.ColVec(int(colIdx)).Decimal()
			prevVal := bufferedCol[int(lastBufferedTupleIdx)]
			var curVal apd.Decimal
			if batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt64(tupleToLookAtIdx) {
				return true
			}
			col := batch.ColVec(int(colIdx)).Decimal()
			curVal = col[int(tupleToLookAtIdx)]
			var match bool

			{
				var cmpResult int
				cmpResult = tree.CompareDecimals(&prevVal, &curVal)
				match = cmpResult == 0
			}

			if !match {
				return true
			}
		case coltypes.Int16:
			// We perform this null check on every equality column of the last
			// buffered tuple regardless of the join type since it is done only once
			// per batch. In some cases (like INNER JOIN, or LEFT OUTER JOIN with the
			// right side being an input) this check will always return false since
			// nulls couldn't be buffered up though.
			if bufferedGroup.ColVec(int(colIdx)).Nulls().NullAt64(uint64(lastBufferedTupleIdx)) {
				return true
			}
			bufferedCol := bufferedGroup.ColVec(int(colIdx)).Int16()
			prevVal := bufferedCol[int(lastBufferedTupleIdx)]
			var curVal int16
			if batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt64(tupleToLookAtIdx) {
				return true
			}
			col := batch.ColVec(int(colIdx)).Int16()
			curVal = col[int(tupleToLookAtIdx)]
			var match bool

			{
				var cmpResult int

				{
					a, b := int64(prevVal), int64(curVal)
					if a < b {
						cmpResult = -1
					} else if a > b {
						cmpResult = 1
					} else {
						cmpResult = 0
					}
				}

				match = cmpResult == 0
			}

			if !match {
				return true
			}
		case coltypes.Int32:
			// We perform this null check on every equality column of the last
			// buffered tuple regardless of the join type since it is done only once
			// per batch. In some cases (like INNER JOIN, or LEFT OUTER JOIN with the
			// right side being an input) this check will always return false since
			// nulls couldn't be buffered up though.
			if bufferedGroup.ColVec(int(colIdx)).Nulls().NullAt64(uint64(lastBufferedTupleIdx)) {
				return true
			}
			bufferedCol := bufferedGroup.ColVec(int(colIdx)).Int32()
			prevVal := bufferedCol[int(lastBufferedTupleIdx)]
			var curVal int32
			if batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt64(tupleToLookAtIdx) {
				return true
			}
			col := batch.ColVec(int(colIdx)).Int32()
			curVal = col[int(tupleToLookAtIdx)]
			var match bool

			{
				var cmpResult int

				{
					a, b := int64(prevVal), int64(curVal)
					if a < b {
						cmpResult = -1
					} else if a > b {
						cmpResult = 1
					} else {
						cmpResult = 0
					}
				}

				match = cmpResult == 0
			}

			if !match {
				return true
			}
		case coltypes.Int64:
			// We perform this null check on every equality column of the last
			// buffered tuple regardless of the join type since it is done only once
			// per batch. In some cases (like INNER JOIN, or LEFT OUTER JOIN with the
			// right side being an input) this check will always return false since
			// nulls couldn't be buffered up though.
			if bufferedGroup.ColVec(int(colIdx)).Nulls().NullAt64(uint64(lastBufferedTupleIdx)) {
				return true
			}
			bufferedCol := bufferedGroup.ColVec(int(colIdx)).Int64()
			prevVal := bufferedCol[int(lastBufferedTupleIdx)]
			var curVal int64
			if batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt64(tupleToLookAtIdx) {
				return true
			}
			col := batch.ColVec(int(colIdx)).Int64()
			curVal = col[int(tupleToLookAtIdx)]
			var match bool

			{
				var cmpResult int

				{
					a, b := int64(prevVal), int64(curVal)
					if a < b {
						cmpResult = -1
					} else if a > b {
						cmpResult = 1
					} else {
						cmpResult = 0
					}
				}

				match = cmpResult == 0
			}

			if !match {
				return true
			}
		case coltypes.Float64:
			// We perform this null check on every equality column of the last
			// buffered tuple regardless of the join type since it is done only once
			// per batch. In some cases (like INNER JOIN, or LEFT OUTER JOIN with the
			// right side being an input) this check will always return false since
			// nulls couldn't be buffered up though.
			if bufferedGroup.ColVec(int(colIdx)).Nulls().NullAt64(uint64(lastBufferedTupleIdx)) {
				return true
			}
			bufferedCol := bufferedGroup.ColVec(int(colIdx)).Float64()
			prevVal := bufferedCol[int(lastBufferedTupleIdx)]
			var curVal float64
			if batch.ColVec(int(colIdx)).MaybeHasNulls() && batch.ColVec(int(colIdx)).Nulls().NullAt64(tupleToLookAtIdx) {
				return true
			}
			col := batch.ColVec(int(colIdx)).Float64()
			curVal = col[int(tupleToLookAtIdx)]
			var match bool

			{
				var cmpResult int

				{
					a, b := float64(prevVal), float64(curVal)
					if a < b {
						cmpResult = -1
					} else if a > b {
						cmpResult = 1
					} else if a == b {
						cmpResult = 0
					} else if math.IsNaN(a) {
						if math.IsNaN(b) {
							cmpResult = 0
						} else {
							cmpResult = -1
						}
					} else {
						cmpResult = 1
					}
				}

				match = cmpResult == 0
			}

			if !match {
				return true
			}
		default:
			execerror.VectorizedInternalPanic(fmt.Sprintf("unhandled type %d", colTyp))
		}
	}
	return false
}
