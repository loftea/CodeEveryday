#include <stdio.h>
#include <stdlib.h>
#include <string.h>
// int boom[1005][1005];
int left[1005];
int right[1005];
int up[1005];
int down[1005];
int main()
{
    int n, m;
    char tmp;
    int cnt = 0;
    int ans = 0;
    scanf("%d %d", &n, &m);
    char * line = malloc(m * sizeof(char));
    for (int i = 1; i < n + 1; i++)
    {
        left[i] = 0;
        right[i] = 0;
    }
    for (int j = 1; j < m + 1; j++)
    {
        up[j] = 0;
        down[j] = 0;
    }

    for (int i = 1; i <= n; i++)
    {
        scanf("%s", line);
        for (int j = 1; j <= m; j++)
        {
            tmp = line[j-1];
            if (tmp == '1')
            {
                // boom[i][j] = 1;
                right[i] = j;
                down[j] = i;
                if (!left[i])
                    left[i] = j;
                if (!up[j])
                    up[j] = i;
                cnt++;
            }
        }
    }

    for (int i = 1; i < n + 1; i++)
    {
        if (left[i])
        {
            ans += right[i] - left[i] + 1 + m;
        }
    }
    for (int i = 1; i < m + 1; i++)
    {
        if (up[i])
        {
            ans += down[i] - up[i] + 1 + n;
        }
    }

    ans -= 4 * cnt;
    printf("%d", ans);
    free(line);
    return 0;
}
