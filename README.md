# AHM Service-Runner

Wake up all the fronend and backend services with one command

## Steps ( If you want to modify the script )

1. Make sure Golang is installed on your machine. If not, [download golang](https://go.dev/dl/) from this link based on your hardware and do the required steup
2. Clone this repo
3. Make desired changes to the main.go file
4. Run `go run main.go` and check for desired output.
5. Once you got the desired output, run command `go build`. This will compile the source code to a binary named `Service-Runner`.

## Steps ( If you just want to use the utility )

1. Download the Service-Runner file from this repo.
2. Open a new "Terminal" on MacOS and input command `./Service-Runner` from the location where the binary is downloaded.

## When you run...

- If this is the first time you are running, the Terminal will show you an error saying `No value found for AHM_ROOT_DIR`
- This is an Environment variable
- Set this in your `bashrc` or `zshrc` with the value equal to the path to the directory where all the services (frontend and backend) are put.
- Once the variable is set, you can rerun the binary and all the services will be up.

## Note

Use MacOS's "Terminal" application and not "iTerm" as iTerm is a third party applicaton which involves a risk in giving it advanced privileges 
