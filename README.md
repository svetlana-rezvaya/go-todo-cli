# go-todo-cli

[![GoDoc](https://godoc.org/github.com/svetlana-rezvaya/go-todo-cli?status.svg)](https://godoc.org/github.com/svetlana-rezvaya/go-todo-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/svetlana-rezvaya/go-todo-cli)](https://goreportcard.com/report/github.com/svetlana-rezvaya/go-todo-cli)
[![Build Status](https://app.travis-ci.com/svetlana-rezvaya/go-todo-cli.svg?branch=master)](https://app.travis-ci.com/svetlana-rezvaya/go-todo-cli)
[![codecov](https://codecov.io/gh/svetlana-rezvaya/go-todo-cli/branch/master/graph/badge.svg)](https://codecov.io/gh/svetlana-rezvaya/go-todo-cli)

The utility for maintaining a to-do list with terminal user interface.

## Installation

```
$ go get github.com/svetlana-rezvaya/go-todo-cli
```

## Usage

```
$ go-todo-cli -h | -help | --help
$ go-todo-cli [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-storage STRING` &mdash; path to a storage file (default: `storage.data`).

## Commands

- filtering commands:
  - `list [done|to do]` &mdash; show all notes or notes filtered by specified status;
  - `find TEXT` &mdash; filter notes by specified text and show them;
  - `date TIMESTAMP` &mdash; filter notes by specified timestamp (in RFC #822 with numeric zone) and show them;
- updating commands:
  - `add TEXT` &mdash; add a new unchecked note with specified text and current timestamp;
  - `check ID` &mdash; check the note with specified ID;
  - `uncheck ID` &mdash; uncheck the note with specified ID;
  - `delete ID` &mdash; delete the note with specified ID;
- misc. commands:
  - `help` &mdash; show the help message;
  - `exit` &mdash; exit.

To combine filtering commands, use the vertical bar (`|`).

## Output Example

```
> list
10001 [_] 30 Jan 21 18:20 +0300 one
10002 [_] 30 Jan 21 18:20 +0300 two
> check 10001
> check 10002
> add three
> add four
> add five
> list
10001 [x] 30 Jan 21 18:20 +0300 one
10002 [x] 30 Jan 21 18:20 +0300 two
10003 [_] 31 Jan 21 18:20 +0300 three
10004 [_] 31 Jan 21 18:20 +0300 four
10005 [_] 31 Jan 21 18:20 +0300 five
> list to do
10003 [_] 31 Jan 21 18:20 +0300 three
10004 [_] 31 Jan 21 18:20 +0300 four
10005 [_] 31 Jan 21 18:20 +0300 five
> list to do | find f
10004 [_] 31 Jan 21 18:20 +0300 four
10005 [_] 31 Jan 21 18:20 +0300 five
>
```

## License

The MIT License (MIT)

Copyright &copy; 2021 svetlana-rezvaya
