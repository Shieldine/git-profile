## git-profile list

List profiles

### Synopsis

Display profiles currently present in your config.

Provide a profile name to list the attributes of the specified profile.
Use flags to filter for a specific origin, name or email.

Examples:
  # List all profiles
  git-profile list

  # Show details of a specific profile
  git-profile list myprofile

  # List all profiles with a specific email
  git-profile list --email user@example.com

  # List all profiles with a specific name
  git-profile list --name "John Doe"

  # List all profiles with a specific origin
  git-profile list --origin github.com


```
git-profile list [profile-name] [flags]
```

### Options

```
  -e, --email string    List profiles with matching email
  -h, --help            help for list
  -n, --name string     List profiles with matching name
  -o, --origin string   List profiles with matching origin
```

### SEE ALSO

* [git-profile](git-profile.md)	 - Manage and automatically set git user profiles

