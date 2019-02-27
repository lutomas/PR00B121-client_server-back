package handler

type ConsultError int64

const (
	NotImplemented ConsultError = iota + 1
	DbAccessError
	DbInsertError
	DbReadError
)
