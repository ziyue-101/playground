package checker

import (
	"context"
	"time"

	"github.com/ziyue-101/playground/healthcheckPoc/pkg/cluster"
)

// Check is a health check func that the framework will initialize at runtime for a HealthChecker.
type HealthCheckerInterface interface {
	Check(ctx context.Context, c cluster.Cluster) error
}

type ErrorType string
type Severity string

// healthchecker definition/metadata
// Required for emitting health check CRs.
type CheckerMeta struct {
	// - name: Must be of DNS1123Subdomain format.
	//   To avoid conflict, the name should be a composite name in the format         of controllerName.targetResourceNamespace.checkName.
	// The targetResourceNamespace is the namespace under which you perform healthcheck and this field is optional. E.g.
	// configconnector.config-control.crd
	// configconnector.config-control.operator
	Name string
	// The kubernetes namespace in which the health check CR is published.
	Namespace string
	// - errortype: Type of the error when a health check fails.
	//   Enum("ServiceError", "UserError").
	ErrorType ErrorType
	// - severity: Severity of the error when a health check fails.
	//   Enum("FatalError", "NonFatalError").
	Severity Severity
	// Timeout for the checker.
	// Optional
	Timeout time.Duration
}
