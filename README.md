# Generator

`generator` is a command-line interface (CLI) application written in Golang that is used to generate random passwords. It has the ability to generate passwords containing alphabetical characters, numeric characters, special symbols, and base 64 encoded passwords. This application is intended to be used for development purposes where you need to generate a password quickly, such as for generating a JWT secret.

## Installation

To install `generator` follow the steps below:

1. Make sure you have Golang installed on your system. If not, download it from the [official website](https://go.dev/dl/) (Or just download the binary that is releavant to your operating system with Github releases)

2. Open the terminal and type the following command
  ```
  go install github.com/VinukaThejana/generator@latest
  ```

## Usage

`generator` canbe used to generate different types of passwords depending on your requirements. By default, the generated password length is 32 characters. However, you can specify the length of the password by using the `--length` flag.

## Flags

`generator` has the following flags that can be used globally

* `--length`: Specify the length of the password to be generated. Default value is 32.
* `--help`: Shows help for the application and each command.

## Subcommands

`generator` has the following subcommands:

* `alpha`: *Generates a password that only contains alphabetical characters*
   <br />&nbsp;&nbsp;&nbsp;&nbsp;Example :-
    Generate a alphabetical password that contains 20 characters
    ```
    generator alpha --length=20
    ```
    
* `numeric`: *Generates a password that only contains numeric characters*
  <br />&nbsp;&nbsp;&nbsp;&nbsp;Example :-
    Generate a numeric only password that only contains 100 characters
    ```
    generator numeric --length=100
    ```
    
* `symbols`: *Generates a password that only contains symbols*
  <br />&nbsp;&nbsp;&nbsp;&nbsp;Example :-
    Generate a symbol only password that contains 16 characters
    ```
    generator symbols --length=16
    ```
    
* `base64`: *Generated a base64 encoded passcode*
  <br />&nbsp;&nbsp;&nbsp;&nbsp;Example :-
  Generate a base64 encoded password (Not providing the length means that it have 32 characters)
  ```
  generator base64
  ```

## License

`generator` is licensed under the [MIT License](https://opensource.org/license/mit/)


