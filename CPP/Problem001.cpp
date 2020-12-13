#include "iostream"

int Problem001();

int main() {
    clock_t start; 
    start = std::clock();
    int ans = Problem001();
    double time_taken = double(clock() - start) / double(CLOCKS_PER_SEC); 

    printf("\nAnswer: %d\n", ans);
    printf("Time Taken: %f seconds\n\n", time_taken);

    return 0;
}
/* 
Problem001 answers the problem at : https://projecteuler.net/problem=1

* Problem 1:
		If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
		Find the sum of all the multiples of 3 or 5 below 1000.
*/
int Problem001(){
    int ans = 0;

    for (int i = 0; i < 1000; i++) { if (i % 3 == 0 || i % 5 == 0){ ans += i; } }

    return ans;
}