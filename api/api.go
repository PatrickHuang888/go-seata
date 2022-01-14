package api

import "context"

const (
	TccKey         = "TccKey"
	TccKeyBranchId = "BranchId"
	TccKeyXid      = "XId"
)

type TCC interface {
	Action(context.Context) error
	Commit(context.Context) error
	Rollback(context.Context) error
}
