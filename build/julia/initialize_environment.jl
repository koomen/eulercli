"""
init.jl

Run this code to initialize your Project Euler julia environment.  Problem solution
programs should activate the projecteulerenv environment.

    import Pkg
    Pkg.activate("projecteulerenv")

New modules and packages can be added "in-line" in a solution program or in this
program via 

    Pkg.add("MyPackage")
"""

import Pkg

Pkg.activate("projecteulerenv")

Pkg.add("BenchmarkTools")
Pkg.add("Combinatorics")
Pkg.add("DataStructures")
Pkg.add("MD5")
Pkg.add("Multisets")
Pkg.add("Primes")
