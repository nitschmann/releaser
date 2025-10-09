## releaser release publish

Publish a release

### Synopsis

Publish a release to a specific upstream.

```
releaser release publish [UPSTREAM] [flags]
```

### Options

```
      --api-token string        API token for the choosen upstream
      --draft                   Release should be published as draft
      --first-tag string        First tag if none is present yet (default "v0.0.1")
  -g, --git-executable string   Git executable (default "git")
      --git-remote string       Git remote (default "origin")
  -h, --help                    help for publish
      --pre-release             Release should be published as pre-release
      --tag string              Specific new tag of the release
      --target string           Target (branch / commit) of the release (default "main")
  -y, --yes                     Automatically answer prompt questions with 'yes'
```

### SEE ALSO

* [releaser release](releaser_release.md)	 - Manage releases

