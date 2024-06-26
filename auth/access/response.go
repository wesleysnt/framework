package access

import "github.com/wesleysnt/framework/contracts/auth/access"

func NewAllowResponse() access.Response {
	return &ResponseImpl{allowed: true}
}

func NewDenyResponse(message string) access.Response {
	return &ResponseImpl{message: message}
}

type ResponseImpl struct {
	allowed bool
	message string
}

func (r *ResponseImpl) Allowed() bool {
	return r.allowed
}

func (r *ResponseImpl) Message() string {
	return r.message
}
