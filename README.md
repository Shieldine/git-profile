# git-profile

[![Go Report Card](https://goreportcard.com/badge/github.com/Shieldine/git-profile)](https://goreportcard.com/report/github.com/Shieldine/git-profile)
[![Latest Release](https://img.shields.io/github/v/release/Shieldine/git-profile)](https://github.com/Shieldine/git-profile/releases/latest)
[![License](https://img.shields.io/github/license/Shieldine/git-profile)](./LICENSE)

A simple CLI to manage and automatically set git user profiles.

## Table of contents

- [Why git-profile?](#why-git-profile)
- [Installation](#installation)
- [Getting started](#getting-started)
- [Common workflows](#common-workflows)
- [Configuration file](#configuration-file)
- [Shell completion](#shell-completion)
- [Command reference](#command-reference)
- [Troubleshooting / FAQ](#troubleshooting--faq)
- [Tips](#tips)
- [Development](#development)
- [License](#license)

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

Additionally, git-profile is available as a homebrew tap. If you are a macOS user and have homebrew installed, run the
following:

```shell
brew tap Shieldine/tools
brew install git-profile
```

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

The executable is located in `\AppData\Local\Programs\git-profile` <br />
The config file (`config.toml`) is created next to the executable on first run.

### Uninstalling

- **Homebrew:** `brew uninstall git-profile` (and `brew untap Shieldine/tools` if you no longer need the tap)
- **Linux/macOS (installed via `install.sh`):** remove the binary from `~/.local/bin/` and, if you want to drop your
  saved profiles too, delete `~/.config/git-profile/`
- **Windows (installed via `install.ps1`):** delete the `\AppData\Local\Programs\git-profile` folder, which contains
  both the executable and `config.toml`

Uninstalling never touches your actual git config (`~/.gitconfig` or a repository's `.git/config`) - whatever
name/email/signing key was last set there stays as-is. Use `git-profile unset` beforehand if you want to clear it too.

## Getting started

```bash
$ git-profile
Usage:
  git-profile [command]

Available Commands:
  add         Add a new profile
  check       Display the currently set attributes
  completion  Generate the autocompletion script for the specified shell
  config      Edit profile configuration file
  help        Help about any command
  init        Automatically set attributes for current repository
  list        List profiles
  rm          Remove existing profiles
  set         Set profile for current repository or globally
  tempset     Set attributes without defining a profile
  unset       Reset attribute config to none
  update      Update one or multiple profiles

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

   Optionally attach a commit signing key. When set, `init`/`set` will also configure
   the repository (or global config) to sign commits with that key:
   ```bash
   # GPG key
   git-profile add work --name "John Doe" --email "john@company.com" --origin github.com --signing-key ABCD1234

   # SSH key
   git-profile add work --name "John Doe" --email "john@company.com" --origin github.com --signing-key ~/.ssh/id_ed25519.pub --signing-format ssh
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

## Configuration file

All profiles live in a single TOML file, created automatically on first run:

| Platform      | Location                                                                                                            |
|---------------|---------------------------------------------------------------------------------------------------------------------|
| Linux / macOS | `~/.config/git-profile/config.toml`                                                                                 |
| Windows       | `config.toml` next to the installed `git-profile.exe` (typically `\AppData\Local\Programs\git-profile\config.toml`) |

Run `git-profile config` to open it in your `$EDITOR` (or pass `--editor`, e.g. `git-profile config --editor code`).
Each profile is a `[[profiles]]` table:

```toml
[[profiles]]
profile_name = "work"
name = "John Doe"
email = "john@company.com"
origin = "github.com"
signing_key = "ABCD1234"
signing_format = ""
```

`signing_key`/`signing_format` are optional - leave them as empty strings for profiles that shouldn't configure
commit signing. You can edit the file directly, but prefer `add`/`update`/`rm` where possible so the file stays
valid TOML.

## Shell completion

git-profile uses [Cobra](https://github.com/spf13/cobra), so autocompletion scripts are built in for
bash, zsh, fish and PowerShell. See setup instructions for your shell with:

```bash
git-profile completion --help
```

For example, on zsh with a Homebrew-managed completion directory:

```bash
git-profile completion zsh > "${fpath[1]}/_git-profile"
```

or, to load it just for the current session:

```bash
source <(git-profile completion bash)
```

## Command reference

The `--help` flag on any command (e.g. `git-profile add --help`) is the source of truth and always matches your
installed version. A generated, browsable copy of the same reference also lives in [`docs/`](docs/git-profile.md)
for quick linking - start at [`docs/git-profile.md`](docs/git-profile.md).

## Troubleshooting / FAQ

**`init`/`set` says "Repository already has correct credentials. Nothing to do." but I changed my profile.**
`init`/`set` only reconfigure git when something actually differs. Double check the profile was saved with
`git-profile list <profile-name>`, or just re-run `git-profile check` to see what's currently applied.

**My commits aren't being signed even though I set a `signing-key`.**
`init`/`set` only configure `user.signingkey`, `gpg.format` and `commit.gpgsign` - they don't verify that
the key itself is valid or that `gpg`/`ssh-agent` is set up correctly on your machine. Run `git commit -S` manually
to see git's own error, and confirm the key value matches what `gpg --list-secret-keys` (or your SSH key file) expects.

**`git-profile init` doesn't find a profile / picks the wrong one.**
Profiles are matched by exact string comparison against the repo's remote `origin` URL. Run `git remote get-url origin`
and compare it against `git-profile list`; if they don't match exactly, update the profile's origin with
`git-profile update <name> --origin ...`.

**I want to remove a signing key from a profile without deleting the whole profile.**
Run `git-profile update <profile-name>` interactively and type `none` when prompted for the signing key.

**"not a git repository" error.**
Commands that touch local repo config (`check`, `set`, `tempset`, `unset`) require you to be inside a git repository
unless you pass `--global`.

**How do I completely undo what git-profile has set?**
Run `git-profile unset` (add `--global` for the global config) to clear `user.name`/`user.email`. This doesn't touch
`gpg.format`/`commit.gpgsign` - unset those with plain `git config --unset` if needed.

### Tips

- Run `git-profile init` in any repository you want to handle attributes in. The CLI will guide you from there on.
- Other than `init`, the most important commands are: `add`, `list`, `rm` and `update`
- For some more convenience in handling repositories that you want to play with, take a look at `check`, `set`, `unset`
  and `tempset`
- `check`, `set`, `unset` and `tempset` also support a `--global` flag to manipulate the global git config
- If a profile has a signing key set, `init` and `set` will configure `user.signingkey`, `gpg.format` (when specified)
  and `commit.gpgsign` for you

## Development

This project is in active development.

If you find any bugs and/or have feature suggestions, feel free to
create issues and pull requests.

Before submitting an issue, please check if it hasn't shown up in other
issues to avoid duplicates.

If you change any command's flags, description or examples, regenerate the Markdown reference in `docs/`:

```bash
go run ./tools/gendocs
```

## License

This project is licensed under Apache 2.0.
You can find the license [here](./LICENSE).
