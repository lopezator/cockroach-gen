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
	"math"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/col/coltypes"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/execgen"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/pkg/errors"
)

// Use execgen package to remove unused import warning.
var _ interface{} = execgen.UNSAFEGET

// tuplesDiffer takes in two ColVecs as well as tuple indices to check whether
// the tuples differ.
func tuplesDiffer(
	t coltypes.T,
	aColVec coldata.Vec,
	aTupleIdx int,
	bColVec coldata.Vec,
	bTupleIdx int,
	differ *bool,
) error {
	switch t {
	case coltypes.Bool:
		aCol := aColVec.Bool()
		bCol := bColVec.Bool()
		var unique bool
		arg1 := aCol[aTupleIdx]
		arg2 := bCol[bTupleIdx]

		{
			var cmpResult int

			if !arg1 && arg2 {
				cmpResult = -1
			} else if arg1 && !arg2 {
				cmpResult = 1
			} else {
				cmpResult = 0
			}

			unique = cmpResult != 0
		}

		*differ = *differ || unique
		return nil
	case coltypes.Bytes:
		aCol := aColVec.Bytes()
		bCol := bColVec.Bytes()
		var unique bool
		arg1 := aCol.Get(aTupleIdx)
		arg2 := bCol.Get(bTupleIdx)

		{
			var cmpResult int
			cmpResult = bytes.Compare(arg1, arg2)
			unique = cmpResult != 0
		}

		*differ = *differ || unique
		return nil
	case coltypes.Decimal:
		aCol := aColVec.Decimal()
		bCol := bColVec.Decimal()
		var unique bool
		arg1 := aCol[aTupleIdx]
		arg2 := bCol[bTupleIdx]

		{
			var cmpResult int
			cmpResult = tree.CompareDecimals(&arg1, &arg2)
			unique = cmpResult != 0
		}

		*differ = *differ || unique
		return nil
	case coltypes.Int16:
		aCol := aColVec.Int16()
		bCol := bColVec.Int16()
		var unique bool
		arg1 := aCol[aTupleIdx]
		arg2 := bCol[bTupleIdx]

		{
			var cmpResult int

			{
				a, b := int64(arg1), int64(arg2)
				if a < b {
					cmpResult = -1
				} else if a > b {
					cmpResult = 1
				} else {
					cmpResult = 0
				}
			}

			unique = cmpResult != 0
		}

		*differ = *differ || unique
		return nil
	case coltypes.Int32:
		aCol := aColVec.Int32()
		bCol := bColVec.Int32()
		var unique bool
		arg1 := aCol[aTupleIdx]
		arg2 := bCol[bTupleIdx]

		{
			var cmpResult int

			{
				a, b := int64(arg1), int64(arg2)
				if a < b {
					cmpResult = -1
				} else if a > b {
					cmpResult = 1
				} else {
					cmpResult = 0
				}
			}

			unique = cmpResult != 0
		}

		*differ = *differ || unique
		return nil
	case coltypes.Int64:
		aCol := aColVec.Int64()
		bCol := bColVec.Int64()
		var unique bool
		arg1 := aCol[aTupleIdx]
		arg2 := bCol[bTupleIdx]

		{
			var cmpResult int

			{
				a, b := int64(arg1), int64(arg2)
				if a < b {
					cmpResult = -1
				} else if a > b {
					cmpResult = 1
				} else {
					cmpResult = 0
				}
			}

			unique = cmpResult != 0
		}

		*differ = *differ || unique
		return nil
	case coltypes.Float64:
		aCol := aColVec.Float64()
		bCol := bColVec.Float64()
		var unique bool
		arg1 := aCol[aTupleIdx]
		arg2 := bCol[bTupleIdx]

		{
			var cmpResult int

			{
				a, b := float64(arg1), float64(arg2)
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

			unique = cmpResult != 0
		}

		*differ = *differ || unique
		return nil
	default:
		return errors.Errorf("unsupported tuplesDiffer type %s", t)
	}
}
