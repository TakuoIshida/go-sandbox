package apperror

import "github.com/morikuni/failure"

const (
	ClientError       failure.StringCode = "XXX0001"
	DBConnectionError failure.StringCode = "XXX0002"
	XXAPIRequestError failure.StringCode = "XXX0003"
)
