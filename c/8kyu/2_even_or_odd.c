#include<stdio.h>

const char * even_or_odd(int number)
{
    return number % 2 == 0 ? "Even" : "Odd";
}

int main(){
    const char * result = even_or_odd(10);
    printf("result = %s \n", result);
}