package rpcerrors

import (
	"backend/common/errors"
	"backend/common/errors/rpcerrors/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewRpcError(grpcCode codes.Code, e *errors.Error) error {
	errDetail := &pb.RpcError{
		Code:    e.Code,
		Message: e.Msg,
	}

	st := status.New(grpcCode, e.Msg)
	st, _ = st.WithDetails(errDetail)
	return st.Err()
}

func FromError(err error) (*pb.RpcError, bool) {
	s, ok := status.FromError(err)
	if !ok {
		return nil, false
	}

	for _, detail := range s.Details() {
		if bizErr, ok := detail.(*pb.RpcError); ok {
			return bizErr, true
		}
	}

	return nil, false
}

var (
	InvalidParams   = NewRpcError(codes.InvalidArgument, errors.InvalidParams)
	UserNoFound     = NewRpcError(codes.NotFound, errors.UserNoFound)
	InvalidPassword = NewRpcError(codes.Unauthenticated, errors.InvalidPassword)

	EmailAlreadyRegister    = NewRpcError(codes.AlreadyExists, errors.EmailAlreadyRegister)
	UserNameAlreadyRegister = NewRpcError(codes.AlreadyExists, errors.UserNameAlreadyRegister)

	ServerError      = NewRpcError(codes.Internal, errors.ServerError)
	DBError          = NewRpcError(codes.Internal, errors.DBError)
	GenerateJWTError = NewRpcError(codes.Internal, errors.GenerateJWTError)
)
