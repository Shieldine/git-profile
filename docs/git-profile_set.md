## git-profile set

Set profile for current repository or globally

### Synopsis

Change the current repository's profile to <profile-name>, or set it globally with --global flag.

This command will apply the name and email from the specified profile to your git configuration.
If the profile doesn't exist, you'll be prompted to create it.

Examples:
  # Set a profile for the current repository
  git-profile set work

  # Set a profile globally
  git-profile set personal --global


```
git-profile set <profile-name> [flags]
```

### Options

```
  -g, --global   Set the profile globally instead of for the current repository
  -h, --help     help for set
```

### SEE ALSO

* [git-profile](git-profile.md)	 - Manage and automatically set git user profiles

