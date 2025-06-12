# git-profile
A simple CLI to manage and automatically set git user profiles.

## Why git-profile?
Some developers use their computers for both work-related and private projects.
This usually involves having *at least* two different sets of credentials
for git. If your private project and/or work involve multiple platforms to keep
your source-code on, that increases the number of credential sets you need to manage.
Setting those each time you clone a project can be quite tiresome and - if you forget or misspell something -
lead to the need of amending commits.

With git-profile, you need to type in your attributes exactly *once*.
They get saved in a *profile* along with the project's origin. Upon calling git-profile
in a repository, it will automatically pick a profile based on the origin
and simply set those attributes for you - no need to even remember a profile name!

At the same time, you still get a few little extra commands to have manual control
over your attributes. There are also options to manipulate the global config.


## Installation
There are two installation scripts included in this repository: one for UNIX-based systems and one for Windows users.
Both scripts automatically add the executable to PATH if requested.

### For Linux/MacOS users:

If you have cloned the repository, simply run:
```shell
sh install.sh
```
If not, run:
```shell
curl -fsSL https://raw.githubusercontent.com/Shieldine/git-profile/main/install.sh | bash
```

The executable is located in `~/.local/bin/` <br />
The config file will be created in `~/.config/git-profile/` on first run.

### For Windows users:

If you have cloned the repository, simply run:
```shell
.\install.ps1
```
If not, run:
```shell
powershell -c "irm https://raw.githubusercontent.com/Shieldine/git-profile/main/install.ps1 | iex"
```
The executable is located in `\AppData\Local\Programs\git-profile`

## Getting started

```bash
$ git-profile
Usage:
  git-profile [command]

Available Commands:
  add         Add a new profile
  check       Display the currently set credentials
  completion  Generate the autocompletion script for the specified shell
  config      Edit profile configuration file
  help        Help about any command
  init        Automatically set credentials for current repository
  list        List profiles
  rm          Remove existing profiles
  set         Set profile for current repository or globally
  tempset     Set credentials without defining a profile
  unset       Reset credential config to none
  update      Update an existing profile

Flags:
  -h, --help      help for git-profile
  -v, --version   version for git-profile

Use "git-profile [command] --help" for more information about a command.
```

### Common Workflows

#### Setting up profiles
1. **Create a new profile**:
   ```bash
   git-profile add work --name "John Doe" --email "john@company.com" --origin github.com
   ```

2. **List your profiles**:
   ```bash
   git-profile list
   ```

3. **Update an existing profile**:
   ```bash
   git-profile update work --email "new.email@company.com"
   ```

#### Using profiles in repositories
1. **Automatically set attributes based on repository origin**:
   ```bash
   git-profile init
   ```

2. **Manually set a specific profile**:
   ```bash
   git-profile set personal
   ```

3. **Check current attributes**:
   ```bash
   git-profile check
   ```

4. **Set temporary attributes without creating a profile**:
   ```bash
   git-profile tempset --name "Temp Name" --email "temp@example.com"
   ```

### Tips
- Run `git-profile init` in any repository you want to handle attributes in. The CLI will guide you from there on.
- Other than `init`, the most important commands are: `add`, `list`, `rm` and `update`
- For some more convenience in handling repositories that you want to play with, take a look at `check`, `set`, `unset` and `tempset`
- `check`, `set`, `unset` and `tempset` also support a `--global` flag to manipulate the global git config


## Development
This project is in active development.

If you find any bugs and/or have feature suggestions, feel free to
create issues and pull requests.

Before submitting an issue, please check if it hasn't shown up in other
issues to avoid duplicates.

## License

This project is licensed under Apache 2.0.
You can find the license [here](./LICENSE).
