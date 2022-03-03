package release

// Release data structure to be used in upstreams
type Release struct {
	// Target branch or commit
	Target string `json:"target"`
	// RepoHttpURL is the full HTTP(s) URL of the repository for the supposed release
	RepoHttpURL string `json:"repo_http_url"`
	// RepoName is the name (project) of the repository
	RepoName string `json:"repo_name"`
}
