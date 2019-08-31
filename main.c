#include <stdio.h>
#include "adder/libadder.h"

void a( GoInt a ){
  // printf("Hello %d", a);
}

int main() {
  printf("C says: about to call Go...\n");
  int total = GoAdder(1, 7, &a);
  printf("C says: Go calculated our total as %i\n", total);
  return 0;
}
