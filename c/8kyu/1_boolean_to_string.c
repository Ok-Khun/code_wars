#include <stdio.h>
#include <stdbool.h>

const char *boolean_to_string(bool b)
{
    if (b == true)
    {
        return "true";
    }
    return "false";
}

int main()
{
    const char *result = boolean_to_string(true);
    printf("result = %s \n", result);
}