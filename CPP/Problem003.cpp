#include "iostream"
#include "vector"
#include "time.h"
#include "math.h"

int Problem003();
std::vector<int> factors(long n);
bool isPrime(int n);

int main() {
    time_t start = time(NULL);

    printf("\nAnswer: %d\n", Problem003());
    printf("Time Taken: %ld seconds\n\n", time(NULL) - start); // needs work (segfault)

    return 0;
}

int Problem003(){
    std::vector<int> primeFactors;
    primeFactors = factors(600851475143);
    
    for(auto it = primeFactors.begin(); it != primeFactors.end(); ++it) {
        if (!isPrime(*it)){ primeFactors.erase(it); }
    }
    return int(*std::max_element(std::begin(primeFactors), std::end(primeFactors)));
}

bool isPrime(int n){
    if (n == 2) { return true; }

    int sq = int(sqrt(n));

    for (int i = 2; i <= sq; i++){
        if (n%i==0) { return false; }
    }
    
    return true;
}

std::vector<int> factors(long int n){
    std::vector<int> factList;
    int sq = int(sqrt(n)); 

    for (int i = 2; i <= sq; i++){
        if (n%i == 0){
            factList.push_back(i);
            factList.push_back(int(n/i));
        }
    }
    
    return factList;
}