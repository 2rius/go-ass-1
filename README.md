# go-ass-1

## The algorithm

The solution we implemented is by [Chandy/Misra](https://en.wikipedia.org/wiki/Dining_philosophers_problem#Chandy/Misra_solution) in which they propose a solution
that is fully decentralized with no starving entites and no deadlocks.

The downside is that there can be a long chain of processes waiting for each other.

By giving the forks to the neighbours that have lower id's than yourself you
start out with a state of a directed acyclic graph. Following the algorithm it won't
turn the graph into a cyclic one and ensures therefor that there will be no deadlocks.

## How to run the program

Run the program from it's root like so:

`go run .`

You will be prompted for how many philosophers and forks there will be at the table.

Then you'll select how many meals they need to eat each.

Have fun. :)
