#include "iostream"
#include "vector"

int Problem007();
std::vector<int> primeSieve(int n);
std::vector<int> range(int start, int end);

int main(){
    clock_t start; 
    start = std::clock();
    int ans = Problem007();
    double time_taken = double(clock() - start) / double(CLOCKS_PER_SEC); 

    printf("\nAnswer: %d\n", ans);
    printf("Time Taken: %f seconds\n\n", time_taken);

    return 0;
}

/* 
Problem007 answers the problem at : https://projecteuler.net/problem=7

* Problem 7:

		By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
		What is the 10001st prime number?
*/
int Problem007(){
    std::vector<int> primes = primeSieve(1000000);
    int res = primes[10000];
    return res;
}

std::vector<int> primeSieve(int n){
    std::vector<int> result = std::vector<int>();

    if (n < 2) { return result; }

    std::vector<bool> input(n + 1, true);
    int sqrtN = (int)sqrt(n);

    for (int i = 2; i <= sqrtN; i ++) {
        if (! input[i]) {
            continue;
        }
        
        for (int j = i * i; j <= n; j += i) {
            input[j] = false;
        }
    }

    result.push_back(2);

    for (int i = 3; i <= n; i += 2) {
        if (input[i]) {
            result.push_back(i);
        }
    }

    return result; 
}

std::vector<int> range(int start, int end){
    std::vector<int> res;
    for (int i = start; i < end; i++){
        res.push_back(i);
    }
    return res;
}