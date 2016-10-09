#include <ucontext.h>
#include <stdio.h>
#include <unistd.h>

int main(int argc, char** argv) {
    ucontext_t context;  
    int x = 10;
    getcontext(&context);  
    printf("Hello world. %d\n", x++);  
    sleep(1);  
    setcontext(&context);  

    return 0;
}
