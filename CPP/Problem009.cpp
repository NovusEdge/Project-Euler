#include "iostream"
#include "math.h"

int Problem009();
bool isTriplet(int a, int b, int c);

int main(){
    clock_t start; 
    start = std::clock();
    int ans = Problem009();
    double time_taken = double(clock() - start) / double(CLOCKS_PER_SEC); 

    printf("\nAnswer: %d\n", ans);
    printf("Time Taken: %f seconds\n\n", time_taken);

    return 0;
}

/*
Problem009 answers the problem at : https://projecteuler.net/problem=9

* Problem 9:
		There exists exactly one Pythagorean triplet for which a + b + c = 1000 (a < b < c).
		Find the product abc.
*/
int Problem009(){

    for (int i = 3; i < 500; i++){
        for (int j = i; j < 500; j++){
            int k = sqrt(i*i + j*j);
            if (isTriplet(i, j, k)){
                return i*j*k;
            }
        }
    }
    return 0;
}

bool isTriplet(int a, int b, int c){
    return (a<b && b<c) && (a*a + b*b == c*c) && (a+b+c == 1000);
}