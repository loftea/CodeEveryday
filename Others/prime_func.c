#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int prime[10000]; // 0表示没计算过 1表示是素数 -1表示不是

void initprime(){
    for (int i = 0; i < 10000; i++)
    {
        prime[i]=0;
    }
    prime[2]=1;
}


int is_prime(int a)
{
    if ((a < 10000) && (prime[a] != 0))
    {
        if (prime[a] == 1){
            return 1;
        }
        if (prime[a] ==-1){
            return 0;
        }
    }

    for (int i = 1; 4* i * i <= a; i++)
    {
        if ((a % (2*i-1)) == 0)
        {
            if (a < 10000)
            {
                prime[a] = -1;
            }
            return 0;
        }
    }

    if (a<10000){
        prime[a]=1;
    }

    return 1;
}

int count(int x, int *array, int length)
{
    int ans = 0;
    for (int i = 0; i < length; i++)
    {
        if (array[i] % x == 0)
            ans++;
    }
    return ans;
}

int solve(int l, int r, int *array, int length)
{
    int ans = 0;
    for (int i = l; i < r + 1; i++)
    {
        if (is_prime(i))
        {
            ans += count(i, array, length);
            if (i>2){
                i++;
            }
        }
    }
    return ans;
}

int main()
{
    initprime();
    int n,m;
    int tmp = 0;
    int l, r;
    int *array;

    scanf("%d", &n);
    array = (int *)malloc(n * sizeof(int));

    for (int i = 0; i < n; i++)
    {
        scanf("%d", &array[i]);
    }

    scanf("%d", &m);

    for (int i = 0; i < m; i++)
    {
        scanf("%d %d", &l, &r);
        tmp = solve(l, r, array, n);
        printf("%d\n", tmp);
    }

    // free(ton);
    free(array);
    return 0;
}
