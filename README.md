# NPM Scripts Analyzer

This is a Go program that reads the `package.json` file of an npm project and generates a graph of the npm scripts specified in the `scripts` section of the file. The program takes into account that some scripts may invoke other scripts, creating a sequence of calls. It then searches for duplicate script commands, considering both individual scripts and sequences of scripts that are executed in the same order.

## Installation

Make sure you have Go installed on your system. Then, clone this repository and run the following command in the project directory:

```bash
go build
```

This will generate an executable file in the same directory.

## Usage

The program can be run by executing the generated executable file. The program takes an optional command line argument `-p` or `--package-json-path` to specify the path to the `package.json` file. If the argument is not provided, the program will assume that the `package.json` file is located in the current directory.

To run the program with the default `package.json` file location, use the following command:

```bash
./npm-scripts-analyzer
```

To run the program with a specific `package.json` file, use the following command:

```bash
./npm-scripts-analyzer -p /path/to/package.json
```


## Output

The program generates a graph of the npm scripts and outputs it to the console. The graph is in the form of a tree, where each node represents a script command and its children represent the commands that are executed as part of the script. The program takes into account both individual scripts and sequences of scripts that are executed in the same order when generating the graph.

The program also searches for duplicate script commands and outputs the results. If duplicates are found, the program prints a list of the scripts and the number of times they appear, along with the list of parent scripts that execute them. The program considers both individual scripts and sequences of scripts that are executed in the same order when searching for duplicates.

If no duplicates are found, the program prints a message indicating that no duplicates were found.

## License

This program is licensed under the MIT License. See the LICENSE file for more information.



