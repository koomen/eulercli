# eulercli
A CLI for working on [Project Euler](https://projecteuler.net) problems

## Overview

[Project Euler](https://projecteuler.net) is a wonderful, free resource for anyone interested in programming and mathematics.

The project offers a [series](https://projecteuler.net/archives) of "challenging mathematical/computer programming problems".  For those who learn best by doing, solving these problems is a fun and addictive way to learn a new programming language.

[eulercli](https://github.com/koomen/eulercli) is a command-line interface for working on Project Euler problems. You can use it to
- Create template solution programs in julia, python, or go
- Check answers to most problems in the [archive](https://projecteuler.net/archives)
- Echo problem text and hashed solutions for most problems

eulercli is written in [go](https://golang.org/). I built it to give back to the Project Euler community, and to get some hands-on experience with [cobra](https://github.com/spf13/cobra), and [viper](https://github.com/spf13/viper).

## Installation

eulercli requires [go 1.16](https://golang.org/doc/go1.16) or later.

To install it, clone [this repository](https://github.com/koomen/eulercli) and run `go install` in the repository's root directory.

Make sure you've added `GOBIN` to your `PATH`.  See `go help install` for more details on where to find `GOBIN` on your system.

You can verify your installation by running `euler --version`.

## Usage

### Echo problems and (hashed) answers

eulercli will display many euler problems in plaintext:

```sh
$ euler problem 1
If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
```

It will also display the answer to many problems, hashed using MD5.

```sh
$ euler answer 1
e1edf9d1967ca96767dcc2b2d6df69f4
```

### Generate an empty solution program

eulercli uses the go's [`text/template` package](https://golang.org/pkg/text/template/) to generate ready-to-use solution programs:

```sh
$ euler generate 25 --language julia
julia template not found in ./templates directory
Using template: https://github.com/koomen/eulercli/raw/main/templates/julia/solution.jl
Generated boilerplate julia solution for problem 25:
    ./julia/euler25/solution.jl
```

eulercli looks for solution template files in `./templates`.  If this directory does not exist or if appropriate template files aren't found inside it, they are downloaded on-the-fly from this github repository.

Generated programs are saved in in `./<language>/euler<problem_number>/`.  

**Note**: When you specify a language using the `--language` flag, your preference will be stored in the `.eulercfg` config file in your working directory, so you only need to specify the language once.

### Download solution templates

  You can use `euler pull-template` to download template files from the eulercli repository:

```sh
$ euler pull-template --language julia
Downloaded julia template files:
    ./templates/solution.jl
```

Downloading solution templates is useful if you want to modify these templates or add your own. See below for instructions on writing your own solution templates.


### Check answers

eulercli can also be used to check answers:

```sh
$ euler check 1 <answer>
Congratulations! <answer> is the correct answer for problem 1.
```

You can also pipe the results of your euler solution directly to eulercli for answer-checking.  eulercli will echo your program's output to stdout and uses a simple regular expression to monitor the output for the expected answer.

```sh
$ julia julia/euler001/solution.jl | euler check 1
Found solution <answer>
Congratulations! <answer> is the correct answer for problem 1.
```

## Acknowledgements

Problem text taken from David Corbin's [Project Euler Offline](https://github.com/davidcorbin/euler-offline/blob/master/project_euler_problems.txt) and Kyle Keen's [Local Euler](http://kmkeen.com/local-euler/) projects.

Problem solutions taken from Bai Li's [projecteuler-solutions](https://github.com/luckytoilet/projecteuler-solutions) repository.