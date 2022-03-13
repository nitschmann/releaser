package release

// UpstreamResult is the result of a published release for an upstream
type UpstreamResult struct {
	ID  *int64  `json:"id,omitempty"`
	URL *string `json:"url,omitempty"`
}
