#include "iostream"
#include "vector"
#include "boost/range/irange.hpp"


long Problem010();
std::vector<int> primeSieve(int n);

int main(){
    clock_t start; 
    start = std::clock();
    int ans = Problem010();
    double time_taken = double(clock() - start) / double(CLOCKS_PER_SEC); 

    printf("\nAnswer: %d\n", ans);
    printf("Time Taken: %f seconds\n\n", time_taken);

    return 0;
}

/*
Problem010 answers the problem at : https://projecteuler.net/problem=10

* Problem 10:
		Find the sum of all the primes below two million.
*/
long Problem010() {
    std::vector<int> primes = primeSieve(20);
    long res = 0;

    for (int i: primes){
        printf("%d, ", i);
        res += long(i);
    }

    return res;
}

std::vector<int> primeSieve(int n){
    std::vector<int> pr;
    boost::integer_range<int> nums = boost::irange(1, n);
        

    return pr;
}


