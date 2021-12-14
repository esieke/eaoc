#include "stdio.h"
#include "errno.h"

#define STR_LEN 500

int aocOpen(const char *name, FILE **f)
{
	int lines;
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
	int l = aocOpen("../../1/input", &f);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	int count = 0;
	int in[l];
	for (int i = 1; i < l; i++)
	{
		int r = fscanf(f, "%d", &in[i]);
		if (r == EOF && errno > 0)
			return -2;
	}

	for (int i = 3; i < l; i++)
	{
		if (in[i] + in[i-1] + in[i-2] > in[i-1] + in[i-2] + in[i-3])
			count++;
	}
	printf("%d\n", count);

	return 0;
}
