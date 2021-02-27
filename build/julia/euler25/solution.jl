"""
Problem 25
https://projecteuler.net/problem=25
==========

Here is the problem text 2


Solution
========

[Explain your solution here]

"""


# Print usage instructions
println("""
$(PROGRAM_FILE) - Solve Project Euler problem 25
Usage: $(PROGRAM_FILE) [--profile|-p] [--benchmark|p]
Options:
    --benchmark,-b      Benchmark your solution
    --profile,-p        Profile your solution
""")


# Activate the projecteuler package and add/import useful packages
import Pkg
Pkg.activate("../projecteulerenv")

using BenchmarkTools
using Profile
using MD5

const ANSWER_MD5 = answer123

# Solution code

"""
`solve()`

Solve the problem
"""
function solve()
    return 0
end

# Boilerplate timing code

if "--benchmark" in ARGS || "-b" in ARGS
    println("Benchmarking solution...")
    display(@benchmark solve())
    println()
elseif "--profile" in ARGS || "-p" in ARGS
    println("Profiling solution...")
    @profile solve()
    Profile.print(combine=true, sortedby=:count, mincount=5, maxdepth=8)
else
    println("Solving...")
    solution = solve()
    iscorrect = bytes2hex(md5(string(solution))) == ANSWER_MD5
    println("Obtained $(iscorrect ? "correct" : "incorrect") solution $(solution)")
end