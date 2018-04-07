Author: Kieran O'Halloran
Student ID: G00326425
This is my implementation of a regular expression engine in Go. This is for my 3rd year Graph Theory module in GMIT.

Project Description
You must write a program in the Go programming language [2] that can
build a non-deterministic finite automaton (NFA) from a regular expression,
and can use the NFA to check if the regular expression matches any given
string of text. You must write the program from scratch and cannot use the
regexp package from the Go standard library nor any other external library.
A regular expression is a string containing a series of characters, some
of which may have a special meaning. For example, the three characters
“.”, “|”, and “∗” have the special meanings “concatenate”, “or”, and “Kleene
star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1,
and 1∗ means any number of 1’s. These special characters must be used in
your submission.
Other special characters you might consider allowing as input are brackets
“()” which can be used for grouping, “+” which means “at least one of”, and
“?” which means “zero or one of”. You might also decide to remove the
concatenation character, so that 1.0 becomes 10, with the concatenation
implicit.
You may initially restrict the non-special characters your program works
with to 0 and 1, if you wish. However, you should at least attempt to expand
these to all of the digits, and the characters a to z, and A to Z.
You are expected to be able to break this project into a number of smaller
tasks that are easier to solve, and to plug these together after they have been
completed.

How to Run The Program
Assumes that Git and Go are installed along with the prerequisites. If not, they can be found from https://golang.org/dl/ and https://git-scm.com/downloads

1. Clone the Repository

> git clone https://github.com/123kieran/Graph_Theory
2. Change Directory to the Folder

Open the terminal/command line and navigate into the folder 
eg. > cd Grapg_Theory
2. Run the program

> go run main.go

