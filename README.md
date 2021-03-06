# Project-Euler-100 
A good ol' repository for project-euler solutions

## Index
  - [About](#about)
  - [Code Format](#code-format)

## About
This reporsitory was created (initially) in light of the ProjectEuler100 challange, but quickly turned into a general repository for all solutions.
* Project Euler: https://projecteuler.net/about
* Problems: https://projecteuler.net/archives
* Back to index: [Index](#index)


## Code Format

### Contents:
  - [Python](#python)
  - [Go](#go)
  - [Perl](#perl)
  - [C++](#c)

#### Python:
* Most of the `python` code is written in a format similar to the following:
```py
from time import time
... #importing libs

start = time() # to keep timestamp of when the processes started

...
# code/processes
...

if __main__ == "__name__":
    print(f"\nAnswer: { ans }")
    print(f"Time Taken: { time() - start }\n")
```

* Back To Contents: [Contents](#contents)

#### Go:
* Most of the `go` code is written in a format similar to the following:
```go
package euler

import (
    "time"
    ...
)

/*
Problem discription with link to the it.
*/
func Problem...(){
    vars ...
    start := time.Now() // to keep a track of when the processes started
    
    ...
    // processes ...
    ...

    end = := time.Now()

    fmt.Printf("\nAnswer to Problem... : %d\n", answer_var)
    fmt.Printf("Time Taken: %f seconds\n\n", end.Sub(start).Seconds())
}
```
Another thing to mention is that all helper functions are contained within `Functions.go` in package `euler` i.e. `Go/Solutions/projeuler`, and text/data files required to solve certain problems are in `Go/Solutions`.

* Back To Contents: [Contents](#contents)

#### Perl:
* Most of the `perl` code is written in a format similar to the following:
```perl
#importing libs ...

$start = time(); # timestamp of when the processes started...

...
# Processes
...

$exc_time = time() - $start;
print "\nAnswer: $answer_var\n";
print "Time Taken: $exc_time\n\n";
```
* Back To Contents: [Contents](#contents)

#### C++:
* Most of the `C++` code is written in a format similar to the following:
```C++
#include "iostream"
...

int Problem...();
// function declarations ...

int main() {
    clock_t start; 
    start = std::clock();
    int ans = Problem...();
    double time_taken = double(clock() - start) / double(CLOCKS_PER_SEC); 

    printf("\nAnswer: %d\n", ans);
    printf("Time Taken: %f seconds\n\n", time_taken); 
    
    return 0;
}

/*
    Problem description
    ...
*/
int Problem...(){
    ...
}
// function definitions :
...
``` 
* Back To Contents: [Contents](#contents)
* Back to index: [Index](#index)