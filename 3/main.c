#include "stdio.h"
#include "stdlib.h"
#include "errno.h"

#define STR_LEN 500
// input 
#define BIT_LEN 12
// input training
// #define BIT_LEN 5

int asc(const void *a, const void *b)
{
	return (*(int *)a - *(int *)b);
}

int des(const void *a, const void *b)
{
	return (*(int *)b - *(int *)a);
}

int aocOpen(const char *name, FILE **f)
{
	int lines = 0;
	char buf;

	*f = fopen(name, "r");
	if (*f == NULL)
		return -1;

	while (1)
	{
		int r = fscanf(*f, "%c", &buf);
		if (r == EOF && errno > 0)
			return -2;
		if (buf == '\n' || r == EOF)
			lines++;
		if (r == EOF)
			break;
	}

	if (fseek(*f, 0, SEEK_SET) < 0)
		return -2;

	return lines;
}

int main()
{
	FILE *f;
	int l = aocOpen("../../3/input", &f);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	int in[l];
	for (int i; i < l; i++)
	{
		char buf[STR_LEN];
		int cI = -1;
		int ctr = (BIT_LEN - 1); // 5 Bit Number
		in[i] = 0;
		while (1)
		{
			cI = fgetc(f);
			if (cI < 0 || (char)cI == '\n')
				break;
			if ((char)cI == '1')
				in[i] |= 1 << ctr;
			ctr--;
		}
	}

	int ll;
	ll = l;
	for (int k = BIT_LEN - 1; k >= 0; k--)
	{
		int ones = 0;
		int zeros = 0;
		for (int i = 0; i < ll; i++)
		{
			ones += (1 << k & in[i]) ? 1 : 0;
			zeros += (1 << k & in[i]) ? 0 : 1;
		}
		if (ones >= zeros)
		{
			qsort(in, ll, sizeof(int), des);
			ll = ones;
		}
		else
		{
			qsort(in, ll, sizeof(int), asc);
			ll = zeros;
		}
	}
	int oxyGenRating = in[0];

	ll = l;
	for (int k = BIT_LEN - 1; k >= 0; k--)
	{
		int ones = 0;
		int zeros = 0;
		for (int i = 0; i < ll; i++)
		{
			ones += (1 << k & in[i]) ? 1 : 0;
			zeros += (1 << k & in[i]) ? 0 : 1;
		}
		if (zeros <= ones)
		{
			qsort(in, ll, sizeof(int), asc);
			ll = zeros;
		}
		else
		{
			qsort(in, ll, sizeof(int), des);
			ll = ones;
		}
	}
	int co2ScrubRating = in[0];

	printf("%d\n", oxyGenRating * co2ScrubRating);
	return 0;
}
