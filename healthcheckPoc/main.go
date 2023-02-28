package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ziyue-101/playground/healthcheckPoc/checker"
	"github.com/ziyue-101/playground/healthcheckPoc/manager"
	"github.com/ziyue-101/playground/healthcheckPoc/simplecheck"
	"k8s.io/client-go/rest"
)

func main() {

	mgr := manager.NewManager("fooMgr", rest.Config{})
	cm := checker.CheckerMeta{
		Name:      "foo/api",
		Namespace: "foo-monitoring",
		Severity:  "Fatal",
		ErrorType: "ServiceLevel",
	}
	simplelogic := func(ctx context.Context, config *rest.Config) error {
		time.Sleep(5 * time.Second)
		fmt.Printf("A Simple Health Check Finished!")
		return nil
	}
	simpleChecker := simplecheck.NewSimpleHealthChecker(cm, simplelogic)
	mgr.RegisterHealthChecker(simpleChecker)
	mgr.StartAllChecks()
}
