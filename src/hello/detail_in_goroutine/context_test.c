#include <ucontext.h>  
#include <stdio.h>  

void func1(void * arg)  
{  
    int x = 10;
    printf("  arg addr: %p %p \n", arg, &x);
    puts("1");  
    puts("11");  
    puts("111");  
    puts("1111");  

}  
void context_test()  
{  
    char stack[32768];
    printf("stack addr: %p ~ %p\n", stack, &stack[32768-1]);
    ucontext_t child, main_context;

    getcontext(&child); //获取当前上下文  
    child.uc_stack.ss_sp = stack;//指定栈空间  
    child.uc_stack.ss_size = sizeof(stack);//指定栈空间大小  
    child.uc_stack.ss_flags = 0;  
    child.uc_link = &main_context;//设置后继上下文  

    int arg1 = 100, arg2 = 200;
    makecontext(&child,(void (*)(void))func1, 2, &arg1, &arg2);//修改上下文指向func1函数  

    swapcontext(&main_context, &child);//切换到child上下文，保存当前上下文到main  
    puts("main");//如果设置了后继上下文，func1函数指向完后会返回此处  
}  

int main()  
{  
    context_test();  

    return 0;  
}  
