#include<stdio.h>
#include<stdint.h>

uint64_t fac(uint64_t n);

int main(void) {
  int n;
  n = getchar();
  printf("%d", fac(n));
  return 0;
}

uint64_t fac(uint64_t n) {
  if (n <= 1){
    return 1;
  }

return (n * fac((n - 1)));
}


