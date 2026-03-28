# ketch-cli

Ketch command line interface.

## Obtaining

### Downloading a release

The easiest way to obtain the Ketch CLI is to download from the [latest release](https://github.com/ketch-com/ketch-cli/releases/latest).
Binaries for MacOS(Darwin), Linux and Windows are available. Download and unzip the file for your operating system.

### Building from source code

You can also build from source code using the following:
```shell
git clone git@github.com:ketch-com/ketch-cli.git
cd ketch-cli
go build -o ketch ./cmd/ketch/main.go
```

## Usage

To obtain a token, run the following command (use `ketch.exe` on Windows):

```shell
ketch login
```

This will output something similar to the following:
```shell
Now, go to https://ketch.us.auth0.com/activate?user_code=HZZK-JZJM and confirm the following code:

         +-----------+
         | HZZK-JZJM |
         +-----------+

```

Open the link in your web browser and enter the code. Once completed, the token is printed to the standard output, such as:
```shell
eyJhbGciOiJSUzI1NiIsInR5cCI6Ik8239879487298472xmM3pBWDlFU3NGZi05c1V6diJ9.eyJpc3Muytfaf76Af786aof76fo8sdf6osdf6so8f6sf7s6o8f76asofayf78s6fosiyasof87eas6YyNzc1MDk0LCJhenAiOiJqOWdlbWl6c1hpczVJY1VnOTMxc0JqR295R1N4YlQxYSJ9.os987foFLUIluzydflfyldisutflsdiuftslfiutuftdliut736r3l7ltd83l6
```

For easy use later, you can set the key to an environment variable. For example, using the sample token above:
```shell
export KETCH_TOKEN="eyJhbGciOiJSUzI1NiIsInR5cCI6Ik8239879487298472xmM3pBWDlFU3NGZi05c1V6diJ9.eyJpc3Muytfaf76Af786aof76fo8sdf6osdf6so8f6sf7s6o8f76asofayf78s6fosiyasof87eas6YyNzc1MDk0LCJhenAiOiJqOWdlbWl6c1hpczVJY1VnOTMxc0JqR295R1N4YlQxYSJ9.os987foFLUIluzydflfyldisutflsdiuftslfiutuftdliut736r3l7ltd83l6"
```

Setting an environment variable saves having to specify the token on subsequent calls to the CLI.

On Linux/MacOS, you can set an environment variable automatically using the following:
```shell
export KETCH_TOKEN=$(ketch login)
```

All further examples will assume that a `KETCH_TOKEN` environment variable has been set.

### Transponder

The Transponder is multi-tenanted, so you can use a single transponder for multiple organizations in the Ketch platform.
The token obtained via [login](#Login) will be scoped to a single organization (the one you logged into when creating the
token).

To configure a transponder, you need to know the URL of the transponder and the computer where you run the CLI needs to
be able to access that URL. You can first check if the URL is accessible in your web browser or using `curl`. Once you have
confirmed that the URL is correct, you can then set an environment variable to save from having to specify the transponder URL
on subsequent executions of the CLI.

For example, if your transponder URL is `https://transponder-stage-uswest2.mycompany.com/`, then you can use the following:
```shell
export KETCH_URL="https://transponder-stage-uswest2.mycompany.com/"
```

All further examples will assume that a `KETCH_URL` environment variable has been set.

#### List Connections

Connections belong to organizations and provide secrets and configuration details for establishing connections to data systems.

To list the connections belonging to the current organization, use the following command:

```shell
ketch transponder ls
```

#### Configure Connection

To configure a connection, two preparation steps are required:
1. [Create the connection](https://docs.ketch.com/hc/en-us/articles/5883595869335-Connecting-the-Ketch-Transponder-to-Ketch#install-database-provider-0-2) in the Ketch console and note the connection code (we will use `my_connection` in the examples below)
2. Collect the required [configuration properties](https://docs.ketch.com/hc/en-us/articles/5922260652439-Database-Provider-Configuration-Parameters) for the data system you are configuring

Once you have the prerequisite information, you can run configure such as:
```shell
ketch transponder configure my_connection -P 'username=myuser' -P 'password=*****' ....other properties....
```

Potential problems:
1. The connection hasn't been configured in the Ketch console
2. The connection has been configured in a different organization than the current KETCH_TOKEN
3. Configuration properties are incomplete or incorrect
4. The transponder is unable to access the data system

#### Rotate Organization API Key

To update an organization's API key used by the Transponder, use the following:

```shell
ketch transponder rotate
```

This command will connect to the transponder and then the transponder will connect to the Ketch platform, attempting to
rotate the current API key configured. If successfully rotated, it will replace the old API key. If there is a problem
rotating the API key, then the previous API key will not be replaced and an error will be displayed. The access token
obtained via login will be used to create the API key. Therefore, your user account must have permissions to manage
API keys.

Potential problems:
1. The transponder is unable to connect to the Ketch platform
2. The organization code does not exist
3. The token does not have permission to manage API keys
4. The API key does not have the appropriate permissions
