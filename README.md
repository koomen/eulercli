# eulercli
A CLI for working on [Project Euler](https://projecteuler.net) problems

## Overview

[Project Euler](https://projecteuler.net) is a wonderful, free resource for anyone interested in programming and mathematics.

The project offers a [series](https://projecteuler.net/archives) of "challenging mathematical/computer programming problems".  For those who learn best by doing, solving these problems is a fun and addictive way to learn a new programming language.

[eulercli](https://github.com/koomen/eulercli) is a command-line interface for working on Project Euler problems. You can use it to
- Create template solution programs in julia -- `eulercli generate 42 --language julia`
- Check the output of your solution program -- `julia mysolution.jl | eulercli check`
- Echo problem text and hashed solutions -- `eulercli problem 42`

eulercli is written in [go](https://golang.org/). I built it to give back to the Project Euler community, and to get some hands-on experience with [cobra](https://github.com/spf13/cobra).

## Installation

eulercli requires [go 1.16](https://golang.org/doc/go1.16) or later.

To install it, run

```sh
go install github.com/koomen/eulercli@latest
```

Make sure you've added `GOBIN` to your `PATH`.  See `go help install` for more details on where to find `GOBIN` on your system.

You can verify your installation by running `eulercli --version`.

## Usage

### Echo problem text and (hashed) answers

eulercli will display many euler problems in plaintext:

```sh
$ eulercli problem 1
If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
```

It will also display the answer to many problems, hashed using [MD5](https://en.wikipedia.org/wiki/MD5).

```sh
$ eulercli answer 1
e1edf9d1967ca96767dcc2b2d6df69f4
```

### Generate an empty solution program

eulercli uses the go's [`text/template` package](https://golang.org/pkg/text/template/) to generate ready-to-use solution programs:

```sh
$ eulercli generate 25 --language julia
Writing generated solution files to ./julia
  Wrote file ./julia/src/euler0025/solution.jl
Have fun!
```

### Download solution program template files

eulercli looks for solution template files in `./eulercli_templates`.  If this directory does not exist or if appropriate template files aren't found inside it, they can be downloaded on-the-fly from this github repository.

  You can use `eulercli pull` to trigger this download manually, or update local templates with newer versions if they exist:

```sh
$ eulercli pull
Downloading templates from https://github.com/koomen/eulercli
  Wrote file eulercli_templates/julia/initenv.jl
  Wrote file eulercli_templates/julia/src/euler{{.PaddedProblemNum}}/solution.jl
Successfully pulled template solution files to eulercli_templates
```

Downloading solution templates is useful if you want to modify these templates or add your own. See below for instructions on writing your own solution templates.

### Check your answers (using command line arguments)

eulercli can also be used to check answers:

```sh
$ eulercli check 1 <answer>
Congratulations, <answer> is the correct answer to problem 1!\n
```
### Check your answers (using pipes)

You can also pipe the results of your solution program directly to eulercli for answer-checking.  eulercli will echo your program's output to stdout and check it for the correct answer when your program terminates:

```sh
$ julia solution.jl | eulercli check
Scanning stdin for problem number and correct answer...
-------------------------------------------------------------------------------


solution.jl - Solve Project Euler problem 1
Usage: solution.jl [--profile|-p] [--benchmark|p]
Options:
    --benchmark,-b      Benchmark your solution
    --profile,-p        Profile your solution

 Activating new environment at `~/git/projecteuler/julia/src/projecteulerenv/Project.toml`
Solving project euler problem 1...
Obtained solution <answer>


-------------------------------------------------------------------------------
Extracted problem number 1 from input
Detected answer <answer> in input. 
Congratulations, this is the correct answer to problem 1!
```

How does this work? eulercli scans the output of your solution program.  

If the problem number is not provided as an argument, the regular expression `[Pp]roblem\s*(\d+)` is used to extract the problem number, so as long as e.g. `"Problem 5"` or `"problem 5"` appears somewhere in your program's output, eulercli will be able to infer the problem number.

eulercli also scans your program's output for the correct answer.  If it is found anywhere it your program's output, you should see:

```sh
Detected answer <answer> in input. Congratulations, this is the correct answer to problem 1!
```

If the correct answer is not found in your program's output, you'll see:

```sh
Failed to find correct answer for problem 1 in input.
```

### Adding your own solution program templates

eulercli can use any template files in `./eulercli_templates` to generate solution programs. 

If you'd like to create templates for a new language or modify the templates for an existing language, you can do so by saving them in the `./eulercli_templates/<language>` directory.  

Template solution files and filenames can include [`text/template` package](https://golang.org/pkg/text/template/) directives with the following fields:

- `{{.ProblemNum}}` - the problem number (e.g. "42")
- `{{.PaddedProblemNum}}` - the problem number, padded with 0s (e.g. "0042")
- `{{.ProblemText}}` - the problem text (e.g. "The nth term of the sequence of triangle...")
- `{{.Answer}}` - The correct answer to the problem (e.g. "123")
- `{{.AnswerMD5}}` - The correct answer to the problem, hashed using [MD5](https://en.wikipedia.org/wiki/MD5) (e.g. "ba1f2511fc30423bdbb183fe33f3dd0f")

For example, calling

```sh
$ eulercli generate 42 --language julia
```

will render the following template file

```
./eulercli_templates/julia/src/euler{{.PaddedProblemNum}}/solution.jl
```

to the target output file

```
./julia/src/euler0042/solution.jl
```

If you find ways to improve existing template files or create useful new template files for an as-yet-unsupported language, consider [contributing to this project](#contributing)

## Contributing

Code contributions to eulercli are encouraged and appreciated! If you'd like to contribute, clone this repository, commit your proposed changes, and create a pull request in this repository.

## Acknowledgements

Problem text taken from David Corbin's [Project Euler Offline](https://github.com/davidcorbin/euler-offline/blob/master/project_euler_problems.txt) and Kyle Keen's [Local Euler](http://kmkeen.com/local-euler/) projects.

Problem solutions taken from Bai Li's [projecteuler-solutions](https://github.com/luckytoilet/projecteuler-solutions) repository.
