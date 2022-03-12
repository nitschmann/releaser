package release

// Release data structure to be used in upstreams
type Release struct {
	// Description (changelog) of the the release
	Description string `json:"description"`
	// Name of the release
	Name string `json:"name"`
	// Tag of the new release
	Tag string `json:"tag"`
	// Target branch or commit
	Target string `json:"target"`
	// RepoHttpURL is the full HTTP(s) URL of the repository for the supposed release
	RepoHttpURL string `json:"repo_http_url"`
	// RepoName is the name (project) of the repository
	RepoName string `json:"repo_name"`

	IsDraft      bool `json:"is_draft"`
	IsPreRelease bool `json:"is_pre_release"`
}
