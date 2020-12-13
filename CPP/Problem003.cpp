#include "iostream"
#include "vector"
#include "math.h"

int Problem003();
std::vector<int> factors(long n);
bool isPrime(int n);

int main() {
    clock_t start; 
    start = std::clock();
    int ans = Problem003();
    double time_taken = double(clock() - start) / double(CLOCKS_PER_SEC); 

    printf("\nAnswer: %d\n", ans);
    printf("Time Taken: %f seconds\n\n", time_taken);

    return 0;
}

int Problem003(){
    std::vector<int> primeFactors;
    primeFactors = factors(600851475143);
    int flag = 0;
    
    for(std::vector<int>::iterator it = primeFactors.begin(); it != primeFactors.end(); ++it) {
        if (flag < *it && isPrime(*it)){
            flag = *it;
        }
    }
    return flag;
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