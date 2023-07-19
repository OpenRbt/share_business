package vo

type OperationKind int64

const (
	BalanceAdd OperationKind = iota
	BalanceRemove
)
