#include<stdio.h>
#include<stdlib.h>

int fib(int n);

int fib(int n) {
  if (n <= 1) {
    return n;
  }

  return fib(n - 1) + fib(n - 2);
}


int main(){
  printf("Enter the number to calculate \n");

  int n = getchar();
  printf("%d", fib(n));

  return 0;
}
