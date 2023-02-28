package manager

import (
	"context"

	"github.com/ziyue-101/playground/healthcheckPoc/pkg/checker"
	"github.com/ziyue-101/playground/healthcheckPoc/pkg/cluster"
	"k8s.io/client-go/rest"
)

func NewManager(name string, cfg *rest.Config) (Manager, error) {
	c := cluster.Cluster{Config: cfg}
	return Manager{Cluster: c}, nil
}

type ManagerInterface interface {
	// Register a checker to the manager and don't start the health checker.
	// Individual health checks can be enabled/disabled using command line flags,
	// e.g. --disable-healthchecks="foo,bar".
	// HealthChecker is an interface that has one method `check`. It is defined below.
	RegisterHealthChecker(healthChecker checker.HealthCheckerInterface) error
	// Start all the registered health checks with the manager.
	StartAllChecks(ctx context.Context) error
}

// Manager is a health check manager.
type Manager struct {
	Cluster  cluster.Cluster
	checkers []checker.HealthCheckerInterface
}

func (mgr *Manager) StartAllChecks(ctx context.Context) error {
	for _, checker := range mgr.checkers {
		checker.Check(ctx, mgr.Cluster)
	}
	return nil
}

func (mgr *Manager) RegisterHealthChecker(healthChecker checker.HealthCheckerInterface) error {
	mgr.checkers = append(mgr.checkers, healthChecker)
	return nil
}
