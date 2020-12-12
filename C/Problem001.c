#include <stdio.h>
#include <time.h>
int Problem001();

int main() {
    time_t start = time(NULL);

    printf("\nAnswer: %d\n", Problem001());
    printf("Time Taken: %ld seconds\n\n", time(NULL) - start);
    
    return 0;
}

int Problem001(){
    int ans = 0;
    for (int i = 0; i < 1000; i++) { if (i % 3 == 0 || i % 5 == 0){ ans += i; } }
    return ans;
}