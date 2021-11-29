#envinject

This is a go package and a CLI tool that walks a directory reading all files looking for tokens in the following format:

`$${{ENV_VAR}}`

which it then replaces with the value of that environment variable.