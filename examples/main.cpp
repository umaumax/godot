#include <iostream>
#include <string>

void A();
void B();
void C();
void D();
void E();
void F();

void A() { B(); }
void B() {
  C();
  D();
}
void C() { E(); }
void D() { F(); }
void E() {}
void F() {}

int main(int argc, char* argv[]) {
  A();
  return 0;
}
