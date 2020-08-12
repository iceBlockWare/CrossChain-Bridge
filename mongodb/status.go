package mongodb

// -----------------------------------------------
// swap status change graph
// symbol '--->' mean transfer only under checked condition (eg. manual process)
//
// -----------------------------------------------
// 1. swap register status change graph
//
// TxNotStable -> |- TxVerifyFailed -> manaul
//                |- TxCanRecall -> TxToBeRecall -> |- TxRecallFailed -> manual
//                                                  |- TxProcessed (->MatchTxNotStable)
//                |- TxWithBigValue        ---> TxNotSwapped
//                |- TxSenderNotRegistered ---> TxNotStable
//                |- TxNotSwapped -> |- TxSwapFailed -> manual
//                                   |- TxProcessed (->MatchTxNotStable)
// -----------------------------------------------
// 2. swap result status change graph
//
// TxWithWrongMemo -> manual
// TxWithBigValue        ---> MatchTxEmpty
// TxSenderNotRegistered ---> MatchTxEmpty
// MatchTxEmpty          -> | MatchTxNotStable -> |- MatchTxStable
//                                                |- MatchTxFailed -> manual
// -----------------------------------------------

// SwapStatus swap status
type SwapStatus uint16

// swap status values
const (
	TxNotStable           SwapStatus = iota // 0
	TxVerifyFailed                          // 1
	TxCanRecall                             // 2
	TxToBeRecall                            // 3
	TxRecallFailed                          // 4
	TxNotSwapped                            // 5
	TxSwapFailed                            // 6
	TxProcessed                             // 7
	MatchTxEmpty                            // 8
	MatchTxNotStable                        // 9
	MatchTxStable                           // 10
	TxWithWrongMemo                         // 11
	TxWithBigValue                          // 12
	TxSenderNotRegistered                   // 13
	MatchTxFailed                           // 14
	SwapInBlacklist                         // 15
)

// CanRetry can retry
func (status SwapStatus) CanRetry() bool {
	switch status {
	case TxSenderNotRegistered:
		return true
	default:
		return false
	}
}

// nolint:gocyclo // allow large switch
func (status SwapStatus) String() string {
	switch status {
	case TxNotStable:
		return "TxNotStable"
	case TxVerifyFailed:
		return "TxVerifyFailed"
	case TxCanRecall:
		return "TxCanRecall"
	case TxToBeRecall:
		return "TxToBeRecall"
	case TxRecallFailed:
		return "TxRecallFailed"
	case TxNotSwapped:
		return "TxNotSwapped"
	case TxSwapFailed:
		return "TxSwapFailed"
	case TxProcessed:
		return "TxProcessed"
	case MatchTxEmpty:
		return "MatchTxEmpty"
	case MatchTxNotStable:
		return "MatchTxNotStable"
	case MatchTxStable:
		return "MatchTxStable"
	case TxWithWrongMemo:
		return "TxWithWrongMemo"
	case TxWithBigValue:
		return "TxWithBigValue"
	case TxSenderNotRegistered:
		return "TxSenderNotRegistered"
	case MatchTxFailed:
		return "MatchTxFailed"
	case SwapInBlacklist:
		return "SwapInBlacklist"
	default:
		return "unknown swap status"
	}
}
