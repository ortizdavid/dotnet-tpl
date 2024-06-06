# dotnet-tpl

`dotnet-tpl` is a CLI tool for generating .NET project templates. This tool supports generating different types of .NET templates such as MVC and WebAPI, providing a quick start for your projects.

## Installation

To use `dotnet-tpl`, you need to have Go installed on your machine. If you don't have Go installed, you can download it from the [official website](https://golang.org/dl/).

1. Clone the repository:

    ```sh
    git clone https://github.com/ortizdavid/dotnet-tpl.git
    ```

2. Navigate to the project directory:

    ```sh
    cd dotnet-tpl
    ```

3. Build the project:

    ```sh
    go build -o dotnet-tpl
    ```

## Usage

The `dotnet-tpl` CLI tool generates .NET templates based on the specified template type.

### Syntax

```sh
dotnet-tpl --template <TEMPLATE_TYPE>
