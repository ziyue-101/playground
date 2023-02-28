package cluster

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type Cluster struct {
	Config *rest.Config
	// watchers will populate the below fields.
	WatchEvents map[schema.GroupVersionKind][]watch.Event
	Snapshot    map[schema.GroupVersionKind][]runtime.Object
}
