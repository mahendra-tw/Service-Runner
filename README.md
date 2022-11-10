# AHM Service-Runner

Wake up all the fronend and backend services with one command

## Steps

1. Make sure Golang is installed on your machine. If not, [download golang](https://go.dev/dl/) from this link based on your hardware and do the required steup
2. Clone this repo
3. Make desired changes to the main.go file
4. If you want to compile the script yourself:
  1. Run `go run main.go` and check for desired output.
  2. Once you got the desired output, run command `go build`. This will compile the source code to a binary named `Service-Runner`.
5. If you want to run the provided binary, run command `./<The binary name>`.

## When you run...

- If this is the first time you are running, the Terminal will show you an error saying `No value found for AHM_ROOT_DIR` and AHM_PACKAGE_MANAGER.
- These are Environment variables.
- Set these in your `bashrc` or `zshrc` with the value equal to 
  1. the path to the directory where all the services (frontend and backend) are put for AHM_ROOT_DIR and
  2. the package manager (npm or yarn) for AHM_PACKAGE_MANAGER.
- Once the variables are set, you can rerun the binary and all the services will be up.

## Note

Use MacOS's "Terminal" application and not "iTerm" as iTerm is a third party applicaton which involves a risk in giving it advanced privileges 

## Future Scope

1. The application currently opens up different window for each service. This can be done in one window and multiple tabs.
2. OR services can be run as a background task.
