#include <stdio.h>
#include <time.h>

int fibb(int n);
int Problem002();

int main() {
    time_t start = time(NULL);

    printf("\nAnswer: %d\n", Problem002());
    printf("Time Taken: %ld seconds\n\n", time(NULL) - start);
    
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