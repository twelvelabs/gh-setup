# gh-setup

:sparkles: A GitHub (`gh`) [CLI](https://cli.github.com) extension to setup new repositories.

## Installation

1. Install the `gh` CLI - see the [installation](https://github.com/cli/cli#installation)

   _Installation requires a minimum version (2.0.0) of the the GitHub CLI that supports extensions._

2. Install this extension:

   ```sh
   gh extension install twelvelabs/gh-setup
   ```

## Usage

Navigate to the repo you would like to setup and run:

```sh
gh setup
```

This will:

- Ensure your local repo has been created:
  - `git init`
  - `git add .`
  - `git commit -m "Initial commit"`
- Ensure your remote repo has been created:
  - `gh repo create --source=. --push`

The extension was designed to be run directly after scaffolding out a new project, but is idempotent (so is safe to run at any time). Each step is only run if needed and prompts before taking action.

## Development

Local development requires [Go](https://go.dev) 1.19:

```sh
brew install go
# Or install manually: https://go.dev/doc/install
```

Then:

```sh
git clone git@github.com:twelvelabs/gh-setup.git
cd ./gh-setup

make setup
make build
make install

# For more tasks
make help
```
