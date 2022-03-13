package upstream

import "github.com/nitschmann/releaser/pkg/release/upstream/github"

var r Registry

// Registry is the registered list of upstreams
type Registry struct {
	registry map[string]Upstream
}

func init() {
	r = NewRegistry()
	// Register all upstreams here
	r.add("github", github.New())
}

// GetRegistry returns the current and globaly available Registry instance
func GetRegistry() Registry {
	return r
}

// NewRegistry inits a new instance of Registry with default values
func NewRegistry() Registry {
	return Registry{
		registry: make(map[string]Upstream),
	}
}

func (r Registry) add(name string, u Upstream) {
	r.registry[name] = u
}

// Get an entry of the registry
func (r Registry) Get(name string) Upstream {
	return r.registry[name]
}

// Names of the registered upstreams
func (r Registry) Names() []string {
	var list []string

	for name := range r.registry {
		list = append(list, name)
	}

	return list
}
