# Tiggo

The name is a nod to [*tig*](https://jonas.github.io/tig/) in go ergo Tiggo. 

Tiggo provides a minimal and fast interface for the most common git commands.

## Features
- Git log Viewer: scrollable commit history with cursor selection

## Installation
```bash
go install github.com/Viriathus1/tiggo@latest
```

## Usage
### `git log`
```bash
tiggo log
```
- Naviate commits using ↑ ↓ or j/k.
- Use `q` to quit

## TODO
- [ ] Combined git status, add and commit view
- [ ] Commit detail view (`Enter` to show `git show <hash>`)
- [ ] Configurable themes for personal preference

## Licence
MIT Licence