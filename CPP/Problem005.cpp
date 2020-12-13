#include "iostream"

int Problem005();
int gcd(int a, int b);
int lcm(int a, int b);

int main() {
    clock_t start; 
    start = std::clock();
    int ans = Problem005();
    double time_taken = double(clock() - start) / double(CLOCKS_PER_SEC); 

    printf("\nAnswer: %d\n", ans);
    printf("Time Taken: %f seconds\n\n", time_taken);

    return 0;
}

int Problem005(){
    int lnum = 1;
    for (int  i = 1; i <= 20; i++){
        lnum = lcm(i, lnum);
    }
    return lnum;
}

int gcd(int a, int b) {
   if (b == 0) { return a; }
   return gcd(b, a % b);
}

int lcm(int a, int b){
    return a * (b / gcd(a, b));
}