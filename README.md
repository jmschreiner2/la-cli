# la-cli

A CLI application to query Azure Logic App information.

## Overview

`la-cli` is a command-line interface designed to simplify querying for hard-to-find information within Azure Logic Apps. This tool helps you find runs, triggers, and other essential data directly from your terminal. It leverages your existing Azure CLI credentials, so no extra authentication is required.

## Current Status

This project is in the early stages of development. While the basic structure is in place, many of the core features are not yet fully implemented.

## Prerequisites

Before using `la-cli`, ensure you have the following installed and configured:

- [Go](https://golang.org/doc/install) (version 1.18 or higher)
- [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli)

You must be logged in to your Azure account via the Azure CLI:

```sh
az login
```

## Installation

To install `la-cli`, use `go install`:

```sh
go install github.com/jmschreiner2/la-cli@latest
```

## Usage

The base command for the CLI is `la-cli`.

### Global Flags

- `-v`, `--verbose`: Enable verbose output for debugging.
- `-s`, `--subscription-id`: Set the Azure Subscription ID for the command.

### Commands

#### `set`

The `set` command is used to configure `la-cli`.

##### `set subscription`

Set the Azure Subscription ID to be used for all queries.

```sh
la-cli set subscription <YOUR_SUBSCRIPTION_ID>
```

#### `find` (Planned)

The `find` command will be used to locate specific Logic App resources.

##### `find trigger` (Planned)

The `find trigger` command will allow you to search for specific triggers within your Logic Apps.

## Configuration

`la-cli` can be configured via a YAML file located at `$HOME/.la-cli.yaml`. The CLI also supports environment variables.

## Contributing

Contributions are welcome! If you have suggestions or want to contribute to the code, please open an issue or submit a pull request.
