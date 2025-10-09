## releaser branch

Generate new branch

### Synopsis


Generate a new branch name based on the configuration and check it out (optional).
In case the branch name is invalid (based on git standards) it will raise a warning.


```
releaser branch [NAME] [flags]
```

### Options

```
  -c, --checkout                Checkout the new branch
  -g, --git-executable string   Git executable (default "git")
  -h, --help                    help for branch
  -t, --type string             Specify explicit type
```

### SEE ALSO

* [releaser](README.md)	 - CLI tool for smart and rule based Git branch, commit and release log management

