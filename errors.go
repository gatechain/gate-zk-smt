package zk_smt

import "errors"

var (
	ErrEmptyRoot = errors.New("empty root")

	ErrVersionTooOld = errors.New("the version is lower than the rollback version")

	ErrVersionTooHigh = errors.New("the version is higher than the latest version")

	ErrNodeNotFound = errors.New("tree node not found")

	ErrVersionMismatched = errors.New("the version is mismatched with the database")

	ErrUnexpected = errors.New("unexpected error")

	ErrInvalidKey = errors.New("invalid key")

	ErrInvalidDepth = errors.New("depth must be a multiple of 4")

	ErrExtendNode = errors.New("extending node error")
)
