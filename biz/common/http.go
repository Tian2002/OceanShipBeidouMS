package common

func init() {
	ErrMap[ParamError] = "param error"
}

var ErrMap map[int]string

const ParamError = 1001
