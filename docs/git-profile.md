## git-profile

Manage and automatically set git user profiles

### Synopsis

git-profile is a simple CLI to manage and automatically set git user profiles based on the project's origin.

Save a profile together with its origin and let git-profile set the attributes next time you clone a new repository.
To make managing names and emails more convenient in general, git-profile offers further commands that will let you
check, unset and set credentials without creating a profile. You also get the option to do these things globally.

Profiles are stored in a TOML config file. Run "git-profile config" to open it directly, or "git-profile list" to
view your profiles. Use "git-profile completion --help" to set up shell autocompletion.


### Options

```
  -h, --help   help for git-profile
```

### SEE ALSO

* [git-profile add](git-profile_add.md)	 - Add a new profile
* [git-profile check](git-profile_check.md)	 - Display the currently set attributes
* [git-profile config](git-profile_config.md)	 - Edit profile configuration file
* [git-profile init](git-profile_init.md)	 - Automatically set attributes for current repository
* [git-profile list](git-profile_list.md)	 - List profiles
* [git-profile rm](git-profile_rm.md)	 - Remove existing profiles
* [git-profile set](git-profile_set.md)	 - Set profile for current repository or globally
* [git-profile tempset](git-profile_tempset.md)	 - Set attributes without defining a profile
* [git-profile unset](git-profile_unset.md)	 - Reset attribute config to none
* [git-profile update](git-profile_update.md)	 - Update one or multiple profiles

