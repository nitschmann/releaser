# Configuration (Detailed)

releaser loads YAML configuration from:

1. `./.releaser/config.yaml`
2. `$HOME/.releaser/config.yaml`
3. `/etc/releaser/config.yaml`

It merges global defaults with optional project-level overrides matched by path. This document outlines the structure and provides examples.

## Top-level Structure

```yaml
branch:
  allowed_without_type: true
  delimiter: "-"
  name_format: "{{if .Type}}{{ .Type }} {{end}}{{ .BranchName }}"
  types: ["bug", "feature", "fix", "hotfix"]

commit:
  allowed_without_type: true
  message_format: "{{if .Type}}{{ .Type | ToTitle }}: {{end}}{{ .CommitMessage }}"
  types: ["adjustment", "bug", "feature", "fix", "hotfix"]

release:
  first_tag: "v0.0.1"
  target: "main"
  name_format: "Release {{ .ReleaseTag }} ({{ .DTYear }}-{{ .DTMonth }}-{{ .DTDay }})"
  description_format: |
    ## Changelog
    {{range .GitCommitLogs -}}
    * {{ .Message }}
    {{end}}
  upstreams:
    github: {}
    gitlab: {}

git:
  executable: "git"
  remote: "origin"

flags:
  - name: issue
    description: "Issue number"
    required: false
    skip_for_branch: false
    skip_for_commit: false
    skip_for_release: true
```

## Template Variables
Templates are Go `text/template`. releaser provides helpers:
- `ToLower`, `ToUpper`, `ToTitle`
- Date/time: `DTYear`, `DTMonth`, `DTDay`, `DTHour`, `DTMinute`, `DTSecond`
- Branch: `BranchName`, `Type`
- Commit: `CommitMessage`
- Release: `ReleaseTag`, `ReleaseTarget`
- Git: `GitRemote`, `GitRepoHttpURL`, `GitRepoName`, `GitCommitLogs`
- Custom flags via `.Flags.<name>`

## Project Overrides

```yaml
projects:
  - paths: ["/Users/you/work/acme/*", "~/code/acme/*"]
    branch:
      delimiter: "_"
      types: ["feature", "fix"]
    commit:
      message_format: "{{ .CommitMessage }}"
    release:
      target: "main"
    flags:
      - name: issue
        required: true
```

## Validation
- Types must be lowercase alphanumeric and unique
- `delimiter` must be a single character
- Unknown flags in CLI are rejected

## Upstreams
- `release.upstreams` maps upstream names to settings.
- Default token env vars are used unless overridden per upstream entry.

## Examples

Branch naming with flags:
```yaml
branch:
  name_format: "{{if .Flags.issue}}issue {{ .Flags.issue }} {{end}}{{if .Type}}{{ .Type }} {{end}}{{ .BranchName }}"
  delimiter: "-"
```

Commit message without type:
```yaml
commit:
  allowed_without_type: true
  message_format: "{{ .CommitMessage }}"
```

Release description includes commit messages between latest tag and target.
