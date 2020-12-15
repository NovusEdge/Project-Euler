#include "boost/math/special_functions/sign.hpp"
#include "iostream"

using namespace std;

int main(int argc, char* argv[]){
  cout << copysign(4.2, 7.9) << endl;
  return 0;
}
