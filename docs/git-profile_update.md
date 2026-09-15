## git-profile update

Update one or multiple profiles

### Synopsis

Updates profiles based on provided criteria.

When a profile name is provided, updates only that specific profile.
Without a profile name, updates all profiles matching the filter criteria.

Examples:
  # Update a specific profile interactively
  git-profile update myprofile

  # Update a specific profile with flags
  git-profile update myprofile --name "New Name" --email "new@example.com"

  # Update all profiles with a specific email
  git-profile update --old-email old@example.com --email new@example.com

  # Update all profiles with a specific name
  git-profile update --old-name "Old Name" --name "New Name"

  # Update all profiles with a specific origin
  git-profile update --old-origin github.com --origin gitlab.com

  # Set or change a profile's signing key
  git-profile update myprofile --signing-key ABCD1234

  # Remove a profile's signing key (interactively, answer "none" to the prompt)
  git-profile update myprofile


```
git-profile update [profile-name] [flags]
```

### Options

```
  -e, --email string             Set the new email value
  -h, --help                     help for update
  -n, --name string              Set the new name value
      --old-email string         Filter profiles by email
      --old-name string          Filter profiles by name
      --old-origin string        Filter profiles by origin
      --old-signing-key string   Filter profiles by signing key
  -o, --origin string            Set the new origin value. Type "auto" to use current repository's origin
      --signing-format string    Set the new signing format (openpgp, ssh, x509)
  -s, --signing-key string       Set the new signing key value. Type "none" during an interactive update to remove it
```

### SEE ALSO

* [git-profile](git-profile.md)	 - Manage and automatically set git user profiles

