package manager

import (
	"context"

	"github.com/ziyue-101/playground/healthcheckPoc/checker"
	"k8s.io/client-go/rest"
)

func NewManager(name string, config *rest.Config) (Manager, error) {
	return Manager{config: config}, nil
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
	config   *rest.Config
	checkers []checker.HealthCheckerInterface
}

func (mgr *Manager) GetRestConfig(ctx context.Context) *rest.Config {
	return mgr.config
}

func (mgr *Manager) StartAllChecks(ctx context.Context) error {
	for _, checker := range mgr.checkers {
		checker.Check(ctx, mgr.config)
	}
}

func (mgr *Manager) RegisterHealthChecker(healthChecker checker.HealthCheckerInterface) error {
	mgr.checkers = append(mgr.checkers, healthChecker)
	return nil
}
