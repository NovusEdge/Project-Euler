#include "iostream"
#include "string"
#include "vector"

int Problem004();
bool isPallindrome(long n);
std::string reverseString(std::string s);


int main() {
    clock_t start; 
    start = std::clock();
    int ans = Problem004();
    double time_taken = double(clock() - start) / double(CLOCKS_PER_SEC); 

    printf("\nAnswer: %d\n", ans);
    printf("Time Taken: %f seconds\n\n", time_taken);;

    return 0;
}
/*
Problem004 answers the problem at : https://projecteuler.net/problem=4

* Problem 4:
		A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

		Find the largest palindrome made from the product of two 3-digit numbers.
*/
int Problem004(){
    std::vector<int> pallVec;

    for (int i = 999; i > 1; i--){
        for (int j = 999; j > 1; j--){
            if (isPallindrome(i*j)){
                pallVec.push_back(i*j);
            }
        }
    }    

    int res = *std::max_element(pallVec.begin(), pallVec.end());
    
    return res;
}


bool isPallindrome(long n){
    std::string nstr = std::to_string(n);
    std::string cpy = nstr;

    std::reverse(cpy.begin(), cpy.end());

    return nstr == cpy;
}
