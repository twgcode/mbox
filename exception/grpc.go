/**
@Author: twgcode
@Email: 17600113577@163.com
@Date: 2023/7/10 00:2411
@Description:
*/

package exception

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	grpcCodeToCode = map[codes.Code]Code{
		codes.InvalidArgument:  BadRequest,
		codes.NotFound:         NotFound,
		codes.AlreadyExists:    Conflict,
		codes.PermissionDenied: Forbidden,
		codes.Internal:         InternalServerError,
		codes.Unauthenticated:  Unauthorized,
	}
)

func ConvertGRPCStatusErrorToAPIExceptionError(sourceErr error) (err error) {
	return ConvertGRPCStatusErrorToAPIException(sourceErr)

}

// ConvertGRPCStatusErrorToAPIException 逻辑类似 status.Convert
func ConvertGRPCStatusErrorToAPIException(err error) (a APIException) {
	a, _ = FromGRPCStatusErrorToAPIException(err)
	return a
}

// FromGRPCStatusErrorToAPIException 逻辑类似 status.FromError
func FromGRPCStatusErrorToAPIException(err error) (a APIException, ok bool) {
	var (
		st *status.Status
	)
	if err == nil {
		return nil, true
	}

	if a, ok = FromError(err); ok {
		return
	}

	st, ok = status.FromError(err)
	if ok {
		return FromGRPCStatusToAPIException(st)
	}
	return NewAPIException(UnKnownException, err.Error()), false
}

func FromGRPCStatusToAPIException(st *status.Status) (a APIException, ok bool) {
	var (
		code Code
	)
	code, ok = grpcCodeToCode[st.Code()]
	if ok {
		return NewAPIException(code, st.Message()), ok
	}

	return NewAPIException(UnKnownException, st.Message()), false
}
