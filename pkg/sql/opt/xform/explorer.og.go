// Code generated by optgen; DO NOT EDIT.

package xform

import (
	"github.com/cockroachdb/cockroach/pkg/sql/opt"
	"github.com/cockroachdb/cockroach/pkg/sql/opt/memo"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

func (_e *explorer) exploreExpr(_state *exploreState, _eid memo.ExprID) (_fullyExplored bool) {
	_expr := _e.mem.Expr(_eid)
	switch _expr.Operator() {
	case opt.ScanOp:
		return _e.exploreScan(_state, _eid)
	case opt.SelectOp:
		return _e.exploreSelect(_state, _eid)
	case opt.InnerJoinOp:
		return _e.exploreInnerJoin(_state, _eid)
	case opt.LeftJoinOp:
		return _e.exploreLeftJoin(_state, _eid)
	case opt.RightJoinOp:
		return _e.exploreRightJoin(_state, _eid)
	case opt.FullJoinOp:
		return _e.exploreFullJoin(_state, _eid)
	case opt.SemiJoinOp:
		return _e.exploreSemiJoin(_state, _eid)
	case opt.AntiJoinOp:
		return _e.exploreAntiJoin(_state, _eid)
	case opt.ScalarGroupByOp:
		return _e.exploreScalarGroupBy(_state, _eid)
	case opt.LimitOp:
		return _e.exploreLimit(_state, _eid)
	}

	// No rules for other operator types.
	return true
}

func (_e *explorer) exploreScan(_rootState *exploreState, _root memo.ExprID) (_fullyExplored bool) {
	_rootExpr := _e.mem.Expr(_root).AsScan()
	_fullyExplored = true

	// [GenerateIndexScans]
	{
		if _root.Expr >= _rootState.start {
			def := _rootExpr.Def()
			if _e.funcs.CanGenerateIndexScans(def) {
				if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateIndexScans) {
					_exprs := _e.funcs.GenerateIndexScans(def)
					_before := _e.mem.ExprCount(_root.Group)
					for i := range _exprs {
						_e.mem.MemoizeDenormExpr(_root.Group, _exprs[i])
					}
					if _e.o.appliedRule != nil {
						_after := _e.mem.ExprCount(_root.Group)
						_e.o.appliedRule(opt.GenerateIndexScans, _root.Group, _root.Expr, _after-_before)
					}
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreSelect(_rootState *exploreState, _root memo.ExprID) (_fullyExplored bool) {
	_rootExpr := _e.mem.Expr(_root).AsSelect()
	_fullyExplored = true

	// [ConstrainScan]
	{
		_partlyExplored := _root.Expr < _rootState.start
		_state := _e.exploreGroup(_rootExpr.Input())
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		start := memo.ExprOrdinal(0)
		if _partlyExplored {
			start = _state.start
		}
		for _ord := start; _ord < _state.end; _ord++ {
			_eid := memo.ExprID{Group: _rootExpr.Input(), Expr: _ord}
			_scanExpr := _e.mem.Expr(_eid).AsScan()
			if _scanExpr != nil {
				def := _scanExpr.Def()
				if _e.funcs.CanConstrainScan(def) {
					filter := _rootExpr.Filter()
					if _e.o.matchedRule == nil || _e.o.matchedRule(opt.ConstrainScan) {
						_exprs := _e.funcs.ConstrainScan(filter, def)
						_before := _e.mem.ExprCount(_root.Group)
						for i := range _exprs {
							_e.mem.MemoizeDenormExpr(_root.Group, _exprs[i])
						}
						if _e.o.appliedRule != nil {
							_after := _e.mem.ExprCount(_root.Group)
							_e.o.appliedRule(opt.ConstrainScan, _root.Group, _root.Expr, _after-_before)
						}
					}
				}
			}
		}
	}

	// [PushFilterIntoIndexJoinNoRemainder]
	{
		_partlyExplored := _root.Expr < _rootState.start
		_state := _e.exploreGroup(_rootExpr.Input())
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		start := memo.ExprOrdinal(0)
		if _partlyExplored {
			start = _state.start
		}
		for _ord := start; _ord < _state.end; _ord++ {
			_eid := memo.ExprID{Group: _rootExpr.Input(), Expr: _ord}
			_indexJoinExpr := _e.mem.Expr(_eid).AsIndexJoin()
			if _indexJoinExpr != nil {
				input := _indexJoinExpr.Input()
				def := _indexJoinExpr.Def()
				filter := _rootExpr.Filter()
				if _e.funcs.IsBoundBy(filter, _e.funcs.OutputCols(input)) {
					if _e.o.matchedRule == nil || _e.o.matchedRule(opt.PushFilterIntoIndexJoinNoRemainder) {
						_expr := memo.MakeIndexJoinExpr(
							_e.f.ConstructSelect(
								input,
								filter,
							),
							def,
						)
						_before := _e.mem.ExprCount(_root.Group)
						_e.mem.MemoizeDenormExpr(_root.Group, memo.Expr(_expr))
						if _e.o.appliedRule != nil {
							_after := _e.mem.ExprCount(_root.Group)
							_e.o.appliedRule(opt.PushFilterIntoIndexJoinNoRemainder, _root.Group, _root.Expr, _after-_before)
						}
					}
				}
			}
		}
	}

	// [PushFilterIntoIndexJoin]
	{
		_partlyExplored := _root.Expr < _rootState.start
		_state := _e.exploreGroup(_rootExpr.Input())
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		start := memo.ExprOrdinal(0)
		if _partlyExplored {
			start = _state.start
		}
		for _ord := start; _ord < _state.end; _ord++ {
			_eid := memo.ExprID{Group: _rootExpr.Input(), Expr: _ord}
			_indexJoinExpr := _e.mem.Expr(_eid).AsIndexJoin()
			if _indexJoinExpr != nil {
				input := _indexJoinExpr.Input()
				def := _indexJoinExpr.Def()
				filter := _rootExpr.Filter()
				_eid := memo.MakeNormExprID(_rootExpr.Filter())
				_filtersExpr := _e.mem.Expr(_eid).AsFilters()
				if _filtersExpr != nil {
					list := _filtersExpr.Conditions()
					for _, _item := range _e.mem.LookupList(_filtersExpr.Conditions()) {
						condition := _item
						inputCols := _e.funcs.OutputCols(input)
						if _e.funcs.IsBoundBy(condition, inputCols) {
							if !_e.funcs.IsBoundBy(filter, inputCols) {
								if _e.o.matchedRule == nil || _e.o.matchedRule(opt.PushFilterIntoIndexJoin) {
									_expr := memo.MakeSelectExpr(
										_e.f.ConstructIndexJoin(
											_e.f.ConstructSelect(
												input,
												_e.f.ConstructFilters(
													_e.funcs.ExtractBoundConditions(list, inputCols),
												),
											),
											def,
										),
										_e.f.ConstructFilters(
											_e.funcs.ExtractUnboundConditions(list, inputCols),
										),
									)
									_before := _e.mem.ExprCount(_root.Group)
									_e.mem.MemoizeDenormExpr(_root.Group, memo.Expr(_expr))
									if _e.o.appliedRule != nil {
										_after := _e.mem.ExprCount(_root.Group)
										_e.o.appliedRule(opt.PushFilterIntoIndexJoin, _root.Group, _root.Expr, _after-_before)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	// [ConstrainIndexJoinScan]
	{
		_partlyExplored := _root.Expr < _rootState.start
		_state := _e.exploreGroup(_rootExpr.Input())
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		for _ord := memo.ExprOrdinal(0); _ord < _state.end; _ord++ {
			_partlyExplored := _partlyExplored && _ord < _state.start
			_eid := memo.ExprID{Group: _rootExpr.Input(), Expr: _ord}
			_indexJoinExpr := _e.mem.Expr(_eid).AsIndexJoin()
			if _indexJoinExpr != nil {
				_state := _e.exploreGroup(_indexJoinExpr.Input())
				if !_state.fullyExplored {
					_fullyExplored = false
				}
				start := memo.ExprOrdinal(0)
				if _partlyExplored {
					start = _state.start
				}
				for _ord := start; _ord < _state.end; _ord++ {
					_eid := memo.ExprID{Group: _indexJoinExpr.Input(), Expr: _ord}
					_scanExpr := _e.mem.Expr(_eid).AsScan()
					if _scanExpr != nil {
						scanDef := _scanExpr.Def()
						if _e.funcs.CanConstrainScan(scanDef) {
							indexJoinDef := _indexJoinExpr.Def()
							filter := _rootExpr.Filter()
							if _e.o.matchedRule == nil || _e.o.matchedRule(opt.ConstrainIndexJoinScan) {
								_exprs := _e.funcs.ConstrainIndexJoinScan(filter, scanDef, indexJoinDef)
								_before := _e.mem.ExprCount(_root.Group)
								for i := range _exprs {
									_e.mem.MemoizeDenormExpr(_root.Group, _exprs[i])
								}
								if _e.o.appliedRule != nil {
									_after := _e.mem.ExprCount(_root.Group)
									_e.o.appliedRule(opt.ConstrainIndexJoinScan, _root.Group, _root.Expr, _after-_before)
								}
							}
						}
					}
				}
			}
		}
	}

	// [GenerateInvertedIndexScans]
	{
		_partlyExplored := _root.Expr < _rootState.start
		_state := _e.exploreGroup(_rootExpr.Input())
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		start := memo.ExprOrdinal(0)
		if _partlyExplored {
			start = _state.start
		}
		for _ord := start; _ord < _state.end; _ord++ {
			_eid := memo.ExprID{Group: _rootExpr.Input(), Expr: _ord}
			_scanExpr := _e.mem.Expr(_eid).AsScan()
			if _scanExpr != nil {
				def := _scanExpr.Def()
				if _e.funcs.CanGenerateInvertedIndexScans(def) {
					filter := _rootExpr.Filter()
					if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateInvertedIndexScans) {
						_exprs := _e.funcs.GenerateInvertedIndexScans(def, filter)
						_before := _e.mem.ExprCount(_root.Group)
						for i := range _exprs {
							_e.mem.MemoizeDenormExpr(_root.Group, _exprs[i])
						}
						if _e.o.appliedRule != nil {
							_after := _e.mem.ExprCount(_root.Group)
							_e.o.appliedRule(opt.GenerateInvertedIndexScans, _root.Group, _root.Expr, _after-_before)
						}
					}
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreInnerJoin(_rootState *exploreState, _root memo.ExprID) (_fullyExplored bool) {
	_rootExpr := _e.mem.Expr(_root).AsInnerJoin()
	_fullyExplored = true

	// [CommuteJoin]
	{
		if _root.Expr >= _rootState.start {
			left := _rootExpr.Left()
			right := _rootExpr.Right()
			on := _rootExpr.On()
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.CommuteJoin) {
				_expr := memo.MakeInnerJoinExpr(
					right,
					left,
					on,
				)
				_before := _e.mem.ExprCount(_root.Group)
				_e.mem.MemoizeDenormExpr(_root.Group, memo.Expr(_expr))
				if _e.o.appliedRule != nil {
					_after := _e.mem.ExprCount(_root.Group)
					_e.o.appliedRule(opt.CommuteJoin, _root.Group, _root.Expr, _after-_before)
				}
			}
		}
	}

	// [GenerateMergeJoins]
	{
		if _root.Expr >= _rootState.start {
			left := _rootExpr.Left()
			right := _rootExpr.Right()
			on := _rootExpr.On()
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateMergeJoins) {
				_exprs := _e.funcs.ConstructMergeJoins(opt.InnerJoinOp, left, right, on)
				_before := _e.mem.ExprCount(_root.Group)
				for i := range _exprs {
					_e.mem.MemoizeDenormExpr(_root.Group, _exprs[i])
				}
				if _e.o.appliedRule != nil {
					_after := _e.mem.ExprCount(_root.Group)
					_e.o.appliedRule(opt.GenerateMergeJoins, _root.Group, _root.Expr, _after-_before)
				}
			}
		}
	}

	// [GenerateLookupJoin]
	{
		_partlyExplored := _root.Expr < _rootState.start
		left := _rootExpr.Left()
		_state := _e.exploreGroup(_rootExpr.Right())
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		start := memo.ExprOrdinal(0)
		if _partlyExplored {
			start = _state.start
		}
		for _ord := start; _ord < _state.end; _ord++ {
			_eid := memo.ExprID{Group: _rootExpr.Right(), Expr: _ord}
			_scanExpr := _e.mem.Expr(_eid).AsScan()
			if _scanExpr != nil {
				scanDef := _scanExpr.Def()
				on := _rootExpr.On()
				if _e.funcs.CanUseLookupJoin(left, scanDef, on) {
					if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateLookupJoin) {
						_exprs := _e.funcs.ConstructLookupJoin(opt.InnerJoinOp, left, scanDef, on)
						_before := _e.mem.ExprCount(_root.Group)
						for i := range _exprs {
							_e.mem.MemoizeDenormExpr(_root.Group, _exprs[i])
						}
						if _e.o.appliedRule != nil {
							_after := _e.mem.ExprCount(_root.Group)
							_e.o.appliedRule(opt.GenerateLookupJoin, _root.Group, _root.Expr, _after-_before)
						}
					}
				}
			}
		}
	}

	// [GenerateLookupJoinWithFilter]
	{
		_partlyExplored := _root.Expr < _rootState.start
		left := _rootExpr.Left()
		_state := _e.exploreGroup(_rootExpr.Right())
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		for _ord := memo.ExprOrdinal(0); _ord < _state.end; _ord++ {
			_partlyExplored := _partlyExplored && _ord < _state.start
			_eid := memo.ExprID{Group: _rootExpr.Right(), Expr: _ord}
			_selectExpr := _e.mem.Expr(_eid).AsSelect()
			if _selectExpr != nil {
				_state := _e.exploreGroup(_selectExpr.Input())
				if !_state.fullyExplored {
					_fullyExplored = false
				}
				start := memo.ExprOrdinal(0)
				if _partlyExplored {
					start = _state.start
				}
				for _ord := start; _ord < _state.end; _ord++ {
					_eid := memo.ExprID{Group: _selectExpr.Input(), Expr: _ord}
					_scanExpr := _e.mem.Expr(_eid).AsScan()
					if _scanExpr != nil {
						scanDef := _scanExpr.Def()
						filter := _selectExpr.Filter()
						on := _rootExpr.On()
						if _e.funcs.CanUseLookupJoin(left, scanDef, on) {
							if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateLookupJoinWithFilter) {
								_exprs := _e.funcs.ConstructLookupJoin(opt.InnerJoinOp, left, scanDef, _e.funcs.ConcatFilters(on, filter))
								_before := _e.mem.ExprCount(_root.Group)
								for i := range _exprs {
									_e.mem.MemoizeDenormExpr(_root.Group, _exprs[i])
								}
								if _e.o.appliedRule != nil {
									_after := _e.mem.ExprCount(_root.Group)
									_e.o.appliedRule(opt.GenerateLookupJoinWithFilter, _root.Group, _root.Expr, _after-_before)
								}
							}
						}
					}
				}
			}
		}
	}

	// [PushJoinThroughIndexJoin]
	{
		_partlyExplored := _root.Expr < _rootState.start
		left := _rootExpr.Left()
		_state := _e.exploreGroup(_rootExpr.Right())
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		start := memo.ExprOrdinal(0)
		if _partlyExplored {
			start = _state.start
		}
		for _ord := start; _ord < _state.end; _ord++ {
			_eid := memo.ExprID{Group: _rootExpr.Right(), Expr: _ord}
			_indexJoinExpr := _e.mem.Expr(_eid).AsIndexJoin()
			if _indexJoinExpr != nil {
				right := _indexJoinExpr.Input()
				indexJoinDef := _indexJoinExpr.Def()
				on := _rootExpr.On()
				if _e.funcs.IsBoundBy(on, _e.funcs.OutputCols2(left, right)) {
					if _e.o.matchedRule == nil || _e.o.matchedRule(opt.PushJoinThroughIndexJoin) {
						newJoin := _e.f.ConstructInnerJoin(
							left,
							right,
							on,
						)
						_expr := memo.MakeLookupJoinExpr(
							newJoin,
							_e.f.ConstructTrue(),
							_e.funcs.HoistIndexJoinDef(indexJoinDef, newJoin, opt.InnerJoinOp),
						)
						_before := _e.mem.ExprCount(_root.Group)
						_e.mem.MemoizeDenormExpr(_root.Group, memo.Expr(_expr))
						if _e.o.appliedRule != nil {
							_after := _e.mem.ExprCount(_root.Group)
							_e.o.appliedRule(opt.PushJoinThroughIndexJoin, _root.Group, _root.Expr, _after-_before)
						}
					}
				}
			}
		}
	}

	// [PushJoinThroughIndexJoinWithExtraFilter]
	{
		_partlyExplored := _root.Expr < _rootState.start
		left := _rootExpr.Left()
		_state := _e.exploreGroup(_rootExpr.Right())
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		start := memo.ExprOrdinal(0)
		if _partlyExplored {
			start = _state.start
		}
		for _ord := start; _ord < _state.end; _ord++ {
			_eid := memo.ExprID{Group: _rootExpr.Right(), Expr: _ord}
			_indexJoinExpr := _e.mem.Expr(_eid).AsIndexJoin()
			if _indexJoinExpr != nil {
				right := _indexJoinExpr.Input()
				indexJoinDef := _indexJoinExpr.Def()
				on := _rootExpr.On()
				innerCols := _e.funcs.OutputCols2(left, right)
				if !_e.funcs.IsBoundBy(on, innerCols) {
					_eid := memo.MakeNormExprID(_rootExpr.On())
					_filtersExpr := _e.mem.Expr(_eid).AsFilters()
					if _filtersExpr != nil {
						onList := _filtersExpr.Conditions()
						for _, _item := range _e.mem.LookupList(_filtersExpr.Conditions()) {
							condition := _item
							if _e.funcs.IsBoundBy(condition, innerCols) {
								if _e.o.matchedRule == nil || _e.o.matchedRule(opt.PushJoinThroughIndexJoinWithExtraFilter) {
									newJoin := _e.f.ConstructInnerJoin(
										left,
										right,
										_e.f.ConstructFilters(
											_e.funcs.ExtractBoundConditions(onList, innerCols),
										),
									)
									_expr := memo.MakeLookupJoinExpr(
										newJoin,
										_e.f.ConstructFilters(
											_e.funcs.ExtractUnboundConditions(onList, innerCols),
										),
										_e.funcs.HoistIndexJoinDef(indexJoinDef, newJoin, opt.InnerJoinOp),
									)
									_before := _e.mem.ExprCount(_root.Group)
									_e.mem.MemoizeDenormExpr(_root.Group, memo.Expr(_expr))
									if _e.o.appliedRule != nil {
										_after := _e.mem.ExprCount(_root.Group)
										_e.o.appliedRule(opt.PushJoinThroughIndexJoinWithExtraFilter, _root.Group, _root.Expr, _after-_before)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreLeftJoin(_rootState *exploreState, _root memo.ExprID) (_fullyExplored bool) {
	_rootExpr := _e.mem.Expr(_root).AsLeftJoin()
	_fullyExplored = true

	// [CommuteLeftJoin]
	{
		if _root.Expr >= _rootState.start {
			left := _rootExpr.Left()
			right := _rootExpr.Right()
			on := _rootExpr.On()
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.CommuteLeftJoin) {
				_expr := memo.MakeRightJoinExpr(
					right,
					left,
					on,
				)
				_before := _e.mem.ExprCount(_root.Group)
				_e.mem.MemoizeDenormExpr(_root.Group, memo.Expr(_expr))
				if _e.o.appliedRule != nil {
					_after := _e.mem.ExprCount(_root.Group)
					_e.o.appliedRule(opt.CommuteLeftJoin, _root.Group, _root.Expr, _after-_before)
				}
			}
		}
	}

	// [GenerateMergeJoins]
	{
		if _root.Expr >= _rootState.start {
			left := _rootExpr.Left()
			right := _rootExpr.Right()
			on := _rootExpr.On()
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateMergeJoins) {
				_exprs := _e.funcs.ConstructMergeJoins(opt.LeftJoinOp, left, right, on)
				_before := _e.mem.ExprCount(_root.Group)
				for i := range _exprs {
					_e.mem.MemoizeDenormExpr(_root.Group, _exprs[i])
				}
				if _e.o.appliedRule != nil {
					_after := _e.mem.ExprCount(_root.Group)
					_e.o.appliedRule(opt.GenerateMergeJoins, _root.Group, _root.Expr, _after-_before)
				}
			}
		}
	}

	// [GenerateLookupJoin]
	{
		_partlyExplored := _root.Expr < _rootState.start
		left := _rootExpr.Left()
		_state := _e.exploreGroup(_rootExpr.Right())
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		start := memo.ExprOrdinal(0)
		if _partlyExplored {
			start = _state.start
		}
		for _ord := start; _ord < _state.end; _ord++ {
			_eid := memo.ExprID{Group: _rootExpr.Right(), Expr: _ord}
			_scanExpr := _e.mem.Expr(_eid).AsScan()
			if _scanExpr != nil {
				scanDef := _scanExpr.Def()
				on := _rootExpr.On()
				if _e.funcs.CanUseLookupJoin(left, scanDef, on) {
					if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateLookupJoin) {
						_exprs := _e.funcs.ConstructLookupJoin(opt.LeftJoinOp, left, scanDef, on)
						_before := _e.mem.ExprCount(_root.Group)
						for i := range _exprs {
							_e.mem.MemoizeDenormExpr(_root.Group, _exprs[i])
						}
						if _e.o.appliedRule != nil {
							_after := _e.mem.ExprCount(_root.Group)
							_e.o.appliedRule(opt.GenerateLookupJoin, _root.Group, _root.Expr, _after-_before)
						}
					}
				}
			}
		}
	}

	// [GenerateLookupJoinWithFilter]
	{
		_partlyExplored := _root.Expr < _rootState.start
		left := _rootExpr.Left()
		_state := _e.exploreGroup(_rootExpr.Right())
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		for _ord := memo.ExprOrdinal(0); _ord < _state.end; _ord++ {
			_partlyExplored := _partlyExplored && _ord < _state.start
			_eid := memo.ExprID{Group: _rootExpr.Right(), Expr: _ord}
			_selectExpr := _e.mem.Expr(_eid).AsSelect()
			if _selectExpr != nil {
				_state := _e.exploreGroup(_selectExpr.Input())
				if !_state.fullyExplored {
					_fullyExplored = false
				}
				start := memo.ExprOrdinal(0)
				if _partlyExplored {
					start = _state.start
				}
				for _ord := start; _ord < _state.end; _ord++ {
					_eid := memo.ExprID{Group: _selectExpr.Input(), Expr: _ord}
					_scanExpr := _e.mem.Expr(_eid).AsScan()
					if _scanExpr != nil {
						scanDef := _scanExpr.Def()
						filter := _selectExpr.Filter()
						on := _rootExpr.On()
						if _e.funcs.CanUseLookupJoin(left, scanDef, on) {
							if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateLookupJoinWithFilter) {
								_exprs := _e.funcs.ConstructLookupJoin(opt.LeftJoinOp, left, scanDef, _e.funcs.ConcatFilters(on, filter))
								_before := _e.mem.ExprCount(_root.Group)
								for i := range _exprs {
									_e.mem.MemoizeDenormExpr(_root.Group, _exprs[i])
								}
								if _e.o.appliedRule != nil {
									_after := _e.mem.ExprCount(_root.Group)
									_e.o.appliedRule(opt.GenerateLookupJoinWithFilter, _root.Group, _root.Expr, _after-_before)
								}
							}
						}
					}
				}
			}
		}
	}

	// [PushJoinThroughIndexJoin]
	{
		_partlyExplored := _root.Expr < _rootState.start
		left := _rootExpr.Left()
		_state := _e.exploreGroup(_rootExpr.Right())
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		start := memo.ExprOrdinal(0)
		if _partlyExplored {
			start = _state.start
		}
		for _ord := start; _ord < _state.end; _ord++ {
			_eid := memo.ExprID{Group: _rootExpr.Right(), Expr: _ord}
			_indexJoinExpr := _e.mem.Expr(_eid).AsIndexJoin()
			if _indexJoinExpr != nil {
				right := _indexJoinExpr.Input()
				indexJoinDef := _indexJoinExpr.Def()
				on := _rootExpr.On()
				if _e.funcs.IsBoundBy(on, _e.funcs.OutputCols2(left, right)) {
					if _e.o.matchedRule == nil || _e.o.matchedRule(opt.PushJoinThroughIndexJoin) {
						newJoin := _e.f.ConstructLeftJoin(
							left,
							right,
							on,
						)
						_expr := memo.MakeLookupJoinExpr(
							newJoin,
							_e.f.ConstructTrue(),
							_e.funcs.HoistIndexJoinDef(indexJoinDef, newJoin, opt.LeftJoinOp),
						)
						_before := _e.mem.ExprCount(_root.Group)
						_e.mem.MemoizeDenormExpr(_root.Group, memo.Expr(_expr))
						if _e.o.appliedRule != nil {
							_after := _e.mem.ExprCount(_root.Group)
							_e.o.appliedRule(opt.PushJoinThroughIndexJoin, _root.Group, _root.Expr, _after-_before)
						}
					}
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreRightJoin(_rootState *exploreState, _root memo.ExprID) (_fullyExplored bool) {
	_rootExpr := _e.mem.Expr(_root).AsRightJoin()
	_fullyExplored = true

	// [CommuteRightJoin]
	{
		if _root.Expr >= _rootState.start {
			left := _rootExpr.Left()
			right := _rootExpr.Right()
			on := _rootExpr.On()
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.CommuteRightJoin) {
				_expr := memo.MakeLeftJoinExpr(
					right,
					left,
					on,
				)
				_before := _e.mem.ExprCount(_root.Group)
				_e.mem.MemoizeDenormExpr(_root.Group, memo.Expr(_expr))
				if _e.o.appliedRule != nil {
					_after := _e.mem.ExprCount(_root.Group)
					_e.o.appliedRule(opt.CommuteRightJoin, _root.Group, _root.Expr, _after-_before)
				}
			}
		}
	}

	// [GenerateMergeJoins]
	{
		if _root.Expr >= _rootState.start {
			left := _rootExpr.Left()
			right := _rootExpr.Right()
			on := _rootExpr.On()
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateMergeJoins) {
				_exprs := _e.funcs.ConstructMergeJoins(opt.RightJoinOp, left, right, on)
				_before := _e.mem.ExprCount(_root.Group)
				for i := range _exprs {
					_e.mem.MemoizeDenormExpr(_root.Group, _exprs[i])
				}
				if _e.o.appliedRule != nil {
					_after := _e.mem.ExprCount(_root.Group)
					_e.o.appliedRule(opt.GenerateMergeJoins, _root.Group, _root.Expr, _after-_before)
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreFullJoin(_rootState *exploreState, _root memo.ExprID) (_fullyExplored bool) {
	_rootExpr := _e.mem.Expr(_root).AsFullJoin()
	_fullyExplored = true

	// [CommuteJoin]
	{
		if _root.Expr >= _rootState.start {
			left := _rootExpr.Left()
			right := _rootExpr.Right()
			on := _rootExpr.On()
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.CommuteJoin) {
				_expr := memo.MakeFullJoinExpr(
					right,
					left,
					on,
				)
				_before := _e.mem.ExprCount(_root.Group)
				_e.mem.MemoizeDenormExpr(_root.Group, memo.Expr(_expr))
				if _e.o.appliedRule != nil {
					_after := _e.mem.ExprCount(_root.Group)
					_e.o.appliedRule(opt.CommuteJoin, _root.Group, _root.Expr, _after-_before)
				}
			}
		}
	}

	// [GenerateMergeJoins]
	{
		if _root.Expr >= _rootState.start {
			left := _rootExpr.Left()
			right := _rootExpr.Right()
			on := _rootExpr.On()
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateMergeJoins) {
				_exprs := _e.funcs.ConstructMergeJoins(opt.FullJoinOp, left, right, on)
				_before := _e.mem.ExprCount(_root.Group)
				for i := range _exprs {
					_e.mem.MemoizeDenormExpr(_root.Group, _exprs[i])
				}
				if _e.o.appliedRule != nil {
					_after := _e.mem.ExprCount(_root.Group)
					_e.o.appliedRule(opt.GenerateMergeJoins, _root.Group, _root.Expr, _after-_before)
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreSemiJoin(_rootState *exploreState, _root memo.ExprID) (_fullyExplored bool) {
	_rootExpr := _e.mem.Expr(_root).AsSemiJoin()
	_fullyExplored = true

	// [GenerateMergeJoins]
	{
		if _root.Expr >= _rootState.start {
			left := _rootExpr.Left()
			right := _rootExpr.Right()
			on := _rootExpr.On()
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateMergeJoins) {
				_exprs := _e.funcs.ConstructMergeJoins(opt.SemiJoinOp, left, right, on)
				_before := _e.mem.ExprCount(_root.Group)
				for i := range _exprs {
					_e.mem.MemoizeDenormExpr(_root.Group, _exprs[i])
				}
				if _e.o.appliedRule != nil {
					_after := _e.mem.ExprCount(_root.Group)
					_e.o.appliedRule(opt.GenerateMergeJoins, _root.Group, _root.Expr, _after-_before)
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreAntiJoin(_rootState *exploreState, _root memo.ExprID) (_fullyExplored bool) {
	_rootExpr := _e.mem.Expr(_root).AsAntiJoin()
	_fullyExplored = true

	// [GenerateMergeJoins]
	{
		if _root.Expr >= _rootState.start {
			left := _rootExpr.Left()
			right := _rootExpr.Right()
			on := _rootExpr.On()
			if _e.o.matchedRule == nil || _e.o.matchedRule(opt.GenerateMergeJoins) {
				_exprs := _e.funcs.ConstructMergeJoins(opt.AntiJoinOp, left, right, on)
				_before := _e.mem.ExprCount(_root.Group)
				for i := range _exprs {
					_e.mem.MemoizeDenormExpr(_root.Group, _exprs[i])
				}
				if _e.o.appliedRule != nil {
					_after := _e.mem.ExprCount(_root.Group)
					_e.o.appliedRule(opt.GenerateMergeJoins, _root.Group, _root.Expr, _after-_before)
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreScalarGroupBy(_rootState *exploreState, _root memo.ExprID) (_fullyExplored bool) {
	_rootExpr := _e.mem.Expr(_root).AsScalarGroupBy()
	_fullyExplored = true

	// [ReplaceMinWithLimit]
	{
		if _root.Expr >= _rootState.start {
			input := _rootExpr.Input()
			_eid := memo.MakeNormExprID(_rootExpr.Aggregations())
			_aggregationsExpr := _e.mem.Expr(_eid).AsAggregations()
			if _aggregationsExpr != nil {
				if _aggregationsExpr.Aggs().Length == 1 {
					_item := _e.mem.LookupList(_aggregationsExpr.Aggs())[0]
					_eid := memo.MakeNormExprID(_item)
					_minExpr := _e.mem.Expr(_eid).AsMin()
					if _minExpr != nil {
						variable := _minExpr.Input()
						_eid := memo.MakeNormExprID(_minExpr.Input())
						_variableExpr := _e.mem.Expr(_eid).AsVariable()
						if _variableExpr != nil {
							col := _variableExpr.Col()
							cols := _aggregationsExpr.Cols()
							def := _rootExpr.Def()
							if _e.o.matchedRule == nil || _e.o.matchedRule(opt.ReplaceMinWithLimit) {
								_expr := memo.MakeScalarGroupByExpr(
									_e.f.ConstructLimit(
										_e.f.ConstructSelect(
											input,
											_e.f.ConstructFilters(
												_e.mem.InternList([]memo.GroupID{_e.f.ConstructIsNot(
													variable,
													_e.f.ConstructNull(
														_e.funcs.AnyType(),
													),
												)}),
											),
										),
										_e.f.ConstructConst(
											_e.mem.InternDatum(tree.NewDInt(1)),
										),
										_e.funcs.MakeOrderingChoiceFromColumn(opt.MinOp, col),
									),
									_e.f.ConstructAggregations(
										_e.mem.InternList([]memo.GroupID{_e.f.ConstructConstAgg(
											variable,
										)}),
										cols,
									),
									def,
								)
								_before := _e.mem.ExprCount(_root.Group)
								_e.mem.MemoizeDenormExpr(_root.Group, memo.Expr(_expr))
								if _e.o.appliedRule != nil {
									_after := _e.mem.ExprCount(_root.Group)
									_e.o.appliedRule(opt.ReplaceMinWithLimit, _root.Group, _root.Expr, _after-_before)
								}
							}
						}
					}
				}
			}
		}
	}

	// [ReplaceMaxWithLimit]
	{
		if _root.Expr >= _rootState.start {
			input := _rootExpr.Input()
			_eid := memo.MakeNormExprID(_rootExpr.Aggregations())
			_aggregationsExpr := _e.mem.Expr(_eid).AsAggregations()
			if _aggregationsExpr != nil {
				if _aggregationsExpr.Aggs().Length == 1 {
					_item := _e.mem.LookupList(_aggregationsExpr.Aggs())[0]
					_eid := memo.MakeNormExprID(_item)
					_maxExpr := _e.mem.Expr(_eid).AsMax()
					if _maxExpr != nil {
						variable := _maxExpr.Input()
						_eid := memo.MakeNormExprID(_maxExpr.Input())
						_variableExpr := _e.mem.Expr(_eid).AsVariable()
						if _variableExpr != nil {
							col := _variableExpr.Col()
							cols := _aggregationsExpr.Cols()
							def := _rootExpr.Def()
							if _e.o.matchedRule == nil || _e.o.matchedRule(opt.ReplaceMaxWithLimit) {
								_expr := memo.MakeScalarGroupByExpr(
									_e.f.ConstructLimit(
										_e.f.ConstructSelect(
											input,
											_e.f.ConstructFilters(
												_e.mem.InternList([]memo.GroupID{_e.f.ConstructIsNot(
													variable,
													_e.f.ConstructNull(
														_e.funcs.AnyType(),
													),
												)}),
											),
										),
										_e.f.ConstructConst(
											_e.mem.InternDatum(tree.NewDInt(1)),
										),
										_e.funcs.MakeOrderingChoiceFromColumn(opt.MaxOp, col),
									),
									_e.f.ConstructAggregations(
										_e.mem.InternList([]memo.GroupID{_e.f.ConstructConstAgg(
											variable,
										)}),
										cols,
									),
									def,
								)
								_before := _e.mem.ExprCount(_root.Group)
								_e.mem.MemoizeDenormExpr(_root.Group, memo.Expr(_expr))
								if _e.o.appliedRule != nil {
									_after := _e.mem.ExprCount(_root.Group)
									_e.o.appliedRule(opt.ReplaceMaxWithLimit, _root.Group, _root.Expr, _after-_before)
								}
							}
						}
					}
				}
			}
		}
	}

	return _fullyExplored
}

func (_e *explorer) exploreLimit(_rootState *exploreState, _root memo.ExprID) (_fullyExplored bool) {
	_rootExpr := _e.mem.Expr(_root).AsLimit()
	_fullyExplored = true

	// [PushLimitIntoScan]
	{
		_partlyExplored := _root.Expr < _rootState.start
		_state := _e.exploreGroup(_rootExpr.Input())
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		start := memo.ExprOrdinal(0)
		if _partlyExplored {
			start = _state.start
		}
		for _ord := start; _ord < _state.end; _ord++ {
			_eid := memo.ExprID{Group: _rootExpr.Input(), Expr: _ord}
			_scanExpr := _e.mem.Expr(_eid).AsScan()
			if _scanExpr != nil {
				def := _scanExpr.Def()
				_eid := memo.MakeNormExprID(_rootExpr.Limit())
				_constExpr := _e.mem.Expr(_eid).AsConst()
				if _constExpr != nil {
					limit := _constExpr.Value()
					ordering := _rootExpr.Ordering()
					if _e.funcs.CanLimitScan(def, limit, ordering) {
						if _e.o.matchedRule == nil || _e.o.matchedRule(opt.PushLimitIntoScan) {
							_expr := memo.MakeScanExpr(
								_e.funcs.LimitScanDef(def, limit, ordering),
							)
							_before := _e.mem.ExprCount(_root.Group)
							_e.mem.MemoizeDenormExpr(_root.Group, memo.Expr(_expr))
							if _e.o.appliedRule != nil {
								_after := _e.mem.ExprCount(_root.Group)
								_e.o.appliedRule(opt.PushLimitIntoScan, _root.Group, _root.Expr, _after-_before)
							}
						}
					}
				}
			}
		}
	}

	// [PushLimitIntoIndexJoin]
	{
		_partlyExplored := _root.Expr < _rootState.start
		_state := _e.exploreGroup(_rootExpr.Input())
		if !_state.fullyExplored {
			_fullyExplored = false
		}
		start := memo.ExprOrdinal(0)
		if _partlyExplored {
			start = _state.start
		}
		for _ord := start; _ord < _state.end; _ord++ {
			_eid := memo.ExprID{Group: _rootExpr.Input(), Expr: _ord}
			_indexJoinExpr := _e.mem.Expr(_eid).AsIndexJoin()
			if _indexJoinExpr != nil {
				input := _indexJoinExpr.Input()
				def := _indexJoinExpr.Def()
				limit := _rootExpr.Limit()
				ordering := _rootExpr.Ordering()
				if _e.funcs.HasColsInOrdering(input, ordering) {
					if _e.o.matchedRule == nil || _e.o.matchedRule(opt.PushLimitIntoIndexJoin) {
						_expr := memo.MakeIndexJoinExpr(
							_e.f.ConstructLimit(
								input,
								limit,
								ordering,
							),
							def,
						)
						_before := _e.mem.ExprCount(_root.Group)
						_e.mem.MemoizeDenormExpr(_root.Group, memo.Expr(_expr))
						if _e.o.appliedRule != nil {
							_after := _e.mem.ExprCount(_root.Group)
							_e.o.appliedRule(opt.PushLimitIntoIndexJoin, _root.Group, _root.Expr, _after-_before)
						}
					}
				}
			}
		}
	}

	return _fullyExplored
}
