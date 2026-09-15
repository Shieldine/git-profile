## git-profile check

Display the currently set attributes

### Synopsis

Check what attributes are currently set in the current project or globally.

This command displays the name and email currently configured in git.
Use the --global flag to check the global git configuration instead of the local repository configuration.

Examples:
  # Check local repository attributes
  git-profile check

  # Check global attributes
  git-profile check --global


```
git-profile check [flags]
```

### Options

```
  -g, --global   Check the global credentials instead of the current repository
  -h, --help     help for check
```

### SEE ALSO

* [git-profile](git-profile.md)	 - Manage and automatically set git user profiles

