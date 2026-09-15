## git-profile add

Add a new profile

### Synopsis

Define a new profile with the short name <profile-name>.
Passing the profile name as an arg is optional. If not provided, you will
be asked to provide one.

You will be asked to provide your attributes and an origin.
Use flags to provide them directly.
The origin of your current repository will already be filled in
and subject to confirm or change.

You can also attach a commit signing key to the profile. If set, running
"git-profile init" or "git-profile set" will configure the repository to
sign commits with that key automatically.

Examples:
  # Add a profile interactively
  git-profile add

  # Add a profile with a specific name
  git-profile add myprofile

  # Add a profile with flags
  git-profile add myprofile --name "John Doe" --email "john@example.com" --origin "github.com"

  # Add a profile with auto-detected origin
  git-profile add myprofile --name "John Doe" --email "john@example.com" --origin auto

  # Add a profile with a GPG signing key
  git-profile add work --name "John Doe" --email "john@company.com" --origin github.com --signing-key ABCD1234

  # Add a profile with an SSH signing key
  git-profile add work --name "John Doe" --email "john@company.com" --origin github.com --signing-key ~/.ssh/id_ed25519.pub --signing-format ssh


```
git-profile add [profile-name] [flags]
```

### Options

```
  -e, --email string            Set the email directly
  -h, --help                    help for add
  -n, --name string             Set the name directly
  -o, --origin string           Set the origin directly. Type "auto" to accept origin of the current repository
      --signing-format string   Set the signing format (openpgp, ssh, x509). Defaults to git's own default (openpgp)
  -s, --signing-key string      Set the commit signing key directly
```

### SEE ALSO

* [git-profile](git-profile.md)	 - Manage and automatically set git user profiles

