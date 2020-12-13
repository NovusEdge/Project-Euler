#include "iostream"
#include "math.h"

int Problem006();

int main() {
    clock_t start; 
    start = std::clock();
    int ans = Problem006();
    double time_taken = double(clock() - start) / double(CLOCKS_PER_SEC); 

    printf("\nAnswer: %d\n", ans);
    printf("Time Taken: %f seconds\n\n", time_taken);

    return 0;
}

int Problem006(){
    int sqSum = 0, sumSq = 0;

    for (int i = 1; i <= 100; i++){
        sqSum += pow(i, 2);
        sumSq += i;
    }

    sumSq = pow(sumSq, 2);    
    return abs(sqSum - sumSq);
}