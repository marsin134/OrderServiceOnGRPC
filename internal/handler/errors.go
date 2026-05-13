package handler

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func handleServiceError(err error) error {
	if err == nil {
		return nil
	}

	errMsg := err.Error()

	// Order not found -> NotFound
	if strings.Contains(errMsg, "not found") {
		return status.Error(codes.NotFound, errMsg)
	}

	// Validation errors -> InvalidArgument
	if strings.Contains(errMsg, "cannot transition") ||
		strings.Contains(errMsg, "validation") ||
		strings.Contains(errMsg, "required") {
		return status.Error(codes.FailedPrecondition, errMsg)
	}

	// Unknown status -> InvalidArgument
	if strings.Contains(errMsg, "unknown order status") {
		return status.Error(codes.InvalidArgument, errMsg)
	}

	// Default -> Internal
	return status.Error(codes.Internal, "internal server error")
}
