package common

import "flag"

type RunMode string

const (
	READ RunMode = "read"
	WRITE RunMode = "write"
)

func FlagInit(mode *string) {
	flag.StringVar(mode, "run-mode", string(READ), "how the program works")
}
