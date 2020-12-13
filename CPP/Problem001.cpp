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

int Problem001(){
    int ans = 0;

    for (int i = 0; i < 1000; i++) { if (i % 3 == 0 || i % 5 == 0){ ans += i; } }

    return ans;
}