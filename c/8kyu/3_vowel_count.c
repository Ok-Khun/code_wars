#include <stdio.h>
#include <stdbool.h>
#include <string.h>

bool isInArr(char cArr[], int size, char c)
{
    for (int i = 0; i < size; i++)
    {
        if (c == cArr[i])
        {
            return true;
        }
    }
    return false;
}

int main()
{
    char vowels[6] = {'a', 'e', 'i', 'o', 'u'};
    const char *s = "abracadabra";
    int sizeOfStr = strlen(s);
    int result = 0;
    for (int i = 0; i < sizeOfStr; i++)
    {
        if (isInArr(vowels, 5, s[i]))
        {
            result += 1;
        }
    }
    printf("result = %i \n", result);
    // isInArr(vowels, 5, 'a');
}