# Quake Log Parser

This repo contains the source code of the Quake Log Parser project.

## Requirements

| Name | Version | Notes | Mandatory
|------|---------|---------|---------|
| [golang](https://golang.org/dl/) | >= go1.15.14 | Main programming language | true
| [sh/bash] | depending on OS. Anyway, you should be able do execute any .sh file | Used to lint checks, test processes and some console interface customizations | true
| [make](https://www.gnu.org/software/make/) | depending on OS. Anyway, you should be able do execute make commands to run the project, tests and localenvironment | n/a | true

# Usage

### Start Local
Inside /quake-log-parser, run:
```bash
make run                  # Execute the algorithm
```

# Testing

```bash
make test                 # Run all unit tests
```