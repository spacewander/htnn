package ir

import (
	"context"

	"github.com/go-logr/logr"
)

type retryableError struct {
	err error
}

func (r *retryableError) Error() string {
	return r.err.Error()
}

func markAsRetryable(err error) error {
	return &retryableError{err: err}
}

type Ctx struct {
	context.Context

	logger *logr.Logger
}

func (c *Ctx) Logger() *logr.Logger {
	return c.logger
}