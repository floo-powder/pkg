package common

type StatusReason string

const (
	NotFound   StatusReason = "NotFound"
	BadRequest StatusReason = "BadRequest"
	Succeed    StatusReason = "Succeed"
)
