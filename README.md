# Tri

A simple TODO list app for the CLI.

## Setup

- `chmod +x setup.sh`
- `./setup.sh`

## Usage

```text
Tri is designed to make organizing your tasks fast and easy straight through your terminal

Usage:
  tri [command]

Available Commands:
  add         Add a new Todo task
  completion  Generate the autocompletion script for the specified shell
  done        Completes a TODO item
  help        Help about any command
  list        Lists TODO items

Flags:
      --config string   config file (default is $HOME/.tri/tricfg.yaml)
  -h, --help            help for tri

Use "tri [command] --help" for more information about a command.
```

## Configuration

You can edit the $HOME/.tri/tricfg.yaml config file, use the --config flag file to use your own
config file instead, or you can override the config file by setting environmental variables:

`export TRI_DATAFILE = "$HOME/input/my-todo-list.json"`

