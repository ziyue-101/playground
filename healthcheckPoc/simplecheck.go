package simplecheck

import (
	"context"

	"github.com/ziyue-101/playground/healthcheckPoc/checker"
	"github.com/ziyue-101/playground/healthcheckPoc/manager"

	"k8s.io/client-go/rest"
)

type simpleHealthChecker struct {
	logic SimpleHealthCheck
	meta  checker.CheckerMeta
}

// NewSimpleHealthChecker returns a HealthChecker
func NewSimpleHealthChecker(cm checker.CheckerMeta, shc SimpleHealthCheck) checker.HealthCheckerInterface {
	return simpleHealthChecker{
		logic: shc,
		meta:  cm,
	}

}

// SimpleHealthCheck is runned as a goroutine. It will only be scheduled once.
// User defined health check logic.
type SimpleHealthCheck func(ctx context.Context, config *rest.Config) error

func (shc simpleHealthChecker) Check(ctx context.Context, mgr manager.Manager) error {
	shc.logic(ctx, mgr.GetRestConfig(ctx))
	return nil
}
