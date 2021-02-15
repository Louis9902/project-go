# Go Project Template

## Project Layout

In the root of the project should be at least one file with the `.go` extension.
This file will be imported uf the root namespace it imported. If the project is
an application it should be a file with the name of the project. This file will
invoke the `cmd/<appname>/main.go` command.

If the project is a library it should be a file with the name `config.go` which
will hold some common information about version etc.