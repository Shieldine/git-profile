## git-profile rm

Remove existing profiles

### Synopsis

Remove one or multiple profiles from the configuration.

Use --all flag to remove all profiles.
Use other flags to remove all profiles containing a specific name, email or origin.

Provide <profile-name> to remove only the profile called <profile-name>.
<profile-name> and filtering flags cannot be provided together.

This action cannot be undone.

Examples:
  # Remove a specific profile
  git-profile rm myprofile

  # Remove all profiles
  git-profile rm --all

  # Remove all profiles with a specific email
  git-profile rm --email user@example.com

  # Remove all profiles with a specific name
  git-profile rm --name "John Doe"

  # Remove all profiles with a specific origin
  git-profile rm --origin github.com


```
git-profile rm [profile-name] [flags]
```

### Options

```
  -a, --all             Remove all profiles
  -e, --email string    Remove profiles with email
  -h, --help            help for rm
  -n, --name string     Remove profiles with name
  -o, --origin string   Remove profiles with origin
```

### SEE ALSO

* [git-profile](git-profile.md)	 - Manage and automatically set git user profiles

