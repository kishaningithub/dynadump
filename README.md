# dynadump

[![Build Status](https://github.com/kishaningithub/dynadump/actions/workflows/build.yml/badge.svg)](https://github.com/kishaningithub/dynadump/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/kishaningithub/dynadump)](https://goreportcard.com/report/github.com/kishaningithub/dynadump)
[![Latest release](https://img.shields.io/github/release/kishaningithub/dynadump.svg)](https://github.com/kishaningithub/dynadump/releases)

Export your dynamoDB data as JSON

<!-- TOC -->
* [dynadump](#dynadump)
  * [Installation](#installation)
    * [Using Homebrew (Mac and linux)](#using-homebrew-mac-and-linux)
    * [Using docker](#using-docker)
    * [Using pkgx](#using-pkgx)
    * [Others](#others)
  * [Examples](#examples)
    * [Dump dynamodb data as json](#dump-dynamodb-data-as-json)
  * [Usage](#usage)
  * [Contributing](#contributing)
<!-- TOC -->

## Installation

### Using Homebrew (Mac and linux)

```bash
brew install kishaningithub/tap/dynadump
```

### Using docker

```bash
alias dynadump="docker run -i ghcr.io/kishaningithub/dynadump:latest"

dynadump table_name
```

### Using pkgx

```bash
pkgx dynadump@latest table_name
```

### Others

Head over to the [releases page](https://github.com/kishaningithub/dynadump/releases) and download a binary for your platform

## Examples

### Dump dynamodb data as json

```bash
$ dynadump table_name > data.json
```

### Extracting a specific field

```bash
$ dynadump table_name | jq -r '.field'
```

## Usage

```bash
$ dynadump --help
Export your dynamoDB data as JSON

Usage:
  dynadump table_name [flags]

Examples:
## Dump data from table_name in JSON format
dynadump table_name > data.json

## Extracting a specific field
dynadump table_name | jq -r '.field'

Flags:
  -h, --help      help for dynadump
  -v, --version   version for dynadump
```

## Contributing

PRs are always welcome!. Refer [CONTRIBUTING.md](./CONTRIBUTING.md) for more information



