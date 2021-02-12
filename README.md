# env-utils
For all your environment and secret needs

## To Build
From the top level directory run 
```
go install .
go build
```
you should then be able to call the command with `env-utils`

## Example Usage
There are currently 2 main ways to use this command line tool:
```
env-utils non-secret populate "sample/non-secret.jsonc"

env-utils secret populate "sample/secret.jsonc"
```
Both of which populate env vars based on the contents of the jsonc files provided.
They both share the following flags:

* `-p` or `-persist` which defaults to `true` and causes the program to write out the env vars to a file named `.secret_env` or `.non_secret_env` in the same directory as the command was called from
* `-o` or `-output` lets you rename the output file and place it anywhere you want

the `secret` command also has the flag `-r` or `-region` which defaults to `us-east-1` and sets the appropriate AWS region.

To persist the secrets to your env after you run the commands, simply run `source .secret_env` or `source .non_secret_env`