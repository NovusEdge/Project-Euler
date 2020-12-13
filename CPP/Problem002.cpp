#include "iostream"

int fibb(int n);
int Problem002();

int main() {
    clock_t start; 
    start = std::clock();
    int ans = Problem002();
    double time_taken = double(clock() - start) / double(CLOCKS_PER_SEC); 

    printf("\nAnswer: %d\n", ans);
    printf("Time Taken: %f seconds\n\n", time_taken);

    return 0;
}

int Problem002() {
    int res = 0;
    int i = 2; int k;

    while (1) {
        k = fibb(i); 
        if (k > 4000000){break;}
        if (k%2==0){ res += k; }
        i++;
    }
    return res;
}

int fibb(int n){
    if ( n == 0 || n == 1 ) { return n; }
    if ( n == 2 ) {return 1;}
    return fibb(n-1) + fibb(n-2);
}