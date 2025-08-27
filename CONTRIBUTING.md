# Contributing to ranger-go

We're excited to have you contribute to our project! This document outlines the process you should follow to contribute effectively.

## How to Contribute

Basic instructions about where to send patches, check out source code, and get development support:

- **Patches**: Please raise patches as Pull Requests.
- **Source Code**: You can check out the source code at https://github.com/G-Research/ranger-go.git.
- **Support**: For development support, please reach out by [raising an issue](https://github.com/G-Research/ranger-go/issues).

## Getting Started

Before you start contributing, please follow these steps:

- **Installation Steps**: Clone this repository.
- **Pre-requisites**: Go v1.24+.

## Team

Understand our team structure and guidelines:

- For details on our team and roles, please see the [MAINTAINERS.md](MAINTAINERS.md) file.

## Building Dependencies

Don't forget to install all necessary dependencies:

- **Installation Steps**: [Insert detailed steps for installing dependencies on your platform here].

## Building the Project

Ensure you can build the project successfully:

- **Build Scripts/Instructions**: Ensure this project can be built by running `go build`.

## Workflow and Branching

Our preferred workflow and branching structure:

- We recommend using [git flow](https://nvie.com/posts/a-successful-git-branching-model/).

## Testing Conventions

Our approach to testing:

- **Test Location**: All tests can be found in [ranger_test.go](ranger_test.go).
- **Running Tests**: Tests can be run locally using `go test`.

## Coding Style and Linters

Our coding standards and tools:

- **Coding Standards**: Code should be formatted using `go fmt`.
- **Linters**: We use [golangci-lint](https://github.com/golangci/golangci-lint) to lint this codebase.
- **Static analysis**: Code is checked using `go vet`

## Writing Issues

Guidelines for filing issues:

- **Where to File Issues**: Please file issues [raising an issue](https://github.com/G-Research/ranger-go/issues).
- **Issue Conventions**: Follow our conventions outlined in [ISSUE_TEMPLATE.md](ISSUE_TEMPLATE.md).

## Writing Pull Requests

Guidelines for pull requests:

- **Where to File Pull Requests**: Submit your pull requests by forking this repository and raising [here](https://github.com/G-Research/ranger-go/compare).
- **PR Conventions**: Follow our conventions outlined in [PULL_REQUEST_TEMPLATE.md](PULL_REQUEST_TEMPLATE.md).

## Reviewing Pull Requests

How we review pull requests:

- **Review Process**: Maintainers will be notified of new pull requests and will be reviewed periodically.
- **Reviewers**: Our reviews are conducted by [our maintainers](MAINTAINERS.md).

## Shipping Releases

Our release process:

- **Cadence**: We ship releases when new features or bug fixes are available.
- **Responsible Parties**: Releases are managed by [our maintainers](MAINTAINERS.md).

## Documentation Updates

How we handle documentation:

- **Documentation Location**: Our documentation can be found in this project's [README](README.md).
- **Update Process**: Documentation is updated whenever significant changes are made.
