## git-profile unset

Reset attribute config to none

### Synopsis

Resets git attributes for current repository or globally.
If you unset local config, git will default to your global config.
If you unset global config, git will have no default credentials.

Examples:
  # Unset local repository attributes
  git-profile unset

  # Unset global attributes
  git-profile unset --global


```
git-profile unset [flags]
```

### Options

```
  -g, --global   Unset the credentials globally instead of for the current repository
  -h, --help     help for unset
```

### SEE ALSO

* [git-profile](git-profile.md)	 - Manage and automatically set git user profiles

