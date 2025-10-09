## releaser release new

Generate new release

### Synopsis

Generate the data for a new release. Pass the option argument in case you want only parts of it.

```
releaser release new [OPTION] [flags]
```

### Options

```
      --first-tag string        First tag if none is present yet (default "v0.0.1")
  -g, --git-executable string   Git executable (default "git")
      --git-remote string       Git remote (default "origin")
  -h, --help                    help for new
      --json                    Print the result as JSON
      --tag string              Specific new tag of the release
      --target string           Target (branch / commit) of the release (default "main")
```

### SEE ALSO

* [releaser release](releaser_release.md)	 - Manage releases

