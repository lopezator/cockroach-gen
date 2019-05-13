// Code generated by "stringer -type=Method"; DO NOT EDIT.

package roachpb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Get-0]
	_ = x[Put-1]
	_ = x[ConditionalPut-2]
	_ = x[Increment-3]
	_ = x[Delete-4]
	_ = x[DeleteRange-5]
	_ = x[ClearRange-6]
	_ = x[Scan-7]
	_ = x[ReverseScan-8]
	_ = x[BeginTransaction-9]
	_ = x[EndTransaction-10]
	_ = x[AdminSplit-11]
	_ = x[AdminMerge-12]
	_ = x[AdminTransferLease-13]
	_ = x[AdminChangeReplicas-14]
	_ = x[AdminRelocateRange-15]
	_ = x[HeartbeatTxn-16]
	_ = x[GC-17]
	_ = x[PushTxn-18]
	_ = x[RecoverTxn-19]
	_ = x[QueryTxn-20]
	_ = x[QueryIntent-21]
	_ = x[ResolveIntent-22]
	_ = x[ResolveIntentRange-23]
	_ = x[Merge-24]
	_ = x[TruncateLog-25]
	_ = x[RequestLease-26]
	_ = x[TransferLease-27]
	_ = x[LeaseInfo-28]
	_ = x[ComputeChecksum-29]
	_ = x[CheckConsistency-30]
	_ = x[InitPut-31]
	_ = x[WriteBatch-32]
	_ = x[Export-33]
	_ = x[Import-34]
	_ = x[AdminScatter-35]
	_ = x[AddSSTable-36]
	_ = x[RecomputeStats-37]
	_ = x[Refresh-38]
	_ = x[RefreshRange-39]
	_ = x[Subsume-40]
	_ = x[RangeStats-41]
}

const _Method_name = "GetPutConditionalPutIncrementDeleteDeleteRangeClearRangeScanReverseScanBeginTransactionEndTransactionAdminSplitAdminMergeAdminTransferLeaseAdminChangeReplicasAdminRelocateRangeHeartbeatTxnGCPushTxnRecoverTxnQueryTxnQueryIntentResolveIntentResolveIntentRangeMergeTruncateLogRequestLeaseTransferLeaseLeaseInfoComputeChecksumCheckConsistencyInitPutWriteBatchExportImportAdminScatterAddSSTableRecomputeStatsRefreshRefreshRangeSubsumeRangeStats"

var _Method_index = [...]uint16{0, 3, 6, 20, 29, 35, 46, 56, 60, 71, 87, 101, 111, 121, 139, 158, 176, 188, 190, 197, 207, 215, 226, 239, 257, 262, 273, 285, 298, 307, 322, 338, 345, 355, 361, 367, 379, 389, 403, 410, 422, 429, 439}

func (i Method) String() string {
	if i < 0 || i >= Method(len(_Method_index)-1) {
		return "Method(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Method_name[_Method_index[i]:_Method_index[i+1]]
}
