#include <stdio.h>
#include <stdlib.h>

int cmp(const void *p1, const void *p2)
{
    const int *a = (const int *)p1;
    const int *b = (const int *)p2;
    return a[0] < b[0];
}

int main()
{
    int n, m;
    int left = 0;
    int right = 0;
    int middle = 0;
    int res =0;

    scanf("%d %d", &n, &m);

    int *input = malloc(n * sizeof(int));
    for (int i = 0; i < n; i++)
    {
        scanf("%d", &input[i]);
    }

    qsort(input, n, sizeof(int), cmp);

    while (right < n)
    {
        
    }
}