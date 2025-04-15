package storage

import "errors"

var (
	ErrUniqueBarcode = errors.New("barcode already in use")
	ErrBeginTxn      = errors.New("begin transaction error")
	ErrNotFound      = errors.New("cartridge not found")
	ErrCommitTxn     = errors.New("commit transaction error")
)
