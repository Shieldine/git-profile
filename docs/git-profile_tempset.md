## git-profile tempset

Set attributes without defining a profile

### Synopsis

Set git attributes for the current repository or globally without saving them in a profile.
The attributes can be passed as flags right away.
If you don't pass them, you will be asked to provide a name and an email.

Examples:
  # Set temporary attributes interactively
  git-profile tempset

  # Set temporary attributes with flags
  git-profile tempset --name "John Doe" --email "john@example.com"

  # Set temporary global attributes
  git-profile tempset --global --name "John Doe" --email "john@example.com"


```
git-profile tempset [flags]
```

### Options

```
  -e, --email string   Pass the email directly
  -g, --global         Set the credentials globally instead of for the current repository
  -h, --help           help for tempset
  -n, --name string    Pass the name directly
```

### SEE ALSO

* [git-profile](git-profile.md)	 - Manage and automatically set git user profiles

