## git-profile config

Edit profile configuration file

### Synopsis

Open and edit the config file containing all profiles.
You can manually type in new profiles by using the following scheme:

```
[[profiles]]
  profile_name = ""
  name = ""
  email = ""
  origin = ""
  signing_key = ""
  signing_format = ""
  ```

Examples:
  # Edit config with default editor (vim)
  git-profile config

  # Edit config with a specific editor
  git-profile config --editor nano

  # Edit config with VS Code
  git-profile config --editor code


```
git-profile config [flags]
```

### Options

```
  -e, --editor string   Specify the editor to use (e.g. nano, code). Vim will be used as default
  -h, --help            help for config
```

### SEE ALSO

* [git-profile](git-profile.md)	 - Manage and automatically set git user profiles

