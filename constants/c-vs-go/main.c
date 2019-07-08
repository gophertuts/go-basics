#include<stdio.h>

int main() {
    unsigned int u = 1e9;
    long signed int i = 1;
    // warning: format specifies type 'int' but the argument has type 'long' [-Wformat]
    printf("Result: %d\n", i + u);
}
