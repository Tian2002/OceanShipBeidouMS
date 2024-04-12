package common

func init() {
	ErrMap[ParamError] = "param error"
}

var ErrMap map[int]string

const (
	Error      = 10000
	ParamError = 10001
)
