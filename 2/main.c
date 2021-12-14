#include "stdio.h"
#include "errno.h"
#include "string.h"

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

struct pos
{
	int x;
	int y;
	int aim;
};

int main()
{
	FILE *f;
	int l = aocOpen("../../2/input", &f);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	struct pos in[l];
	for (int i; i < l; i++)
	{
		char c[STR_LEN];
		int val;
		int r = fscanf(f, "%s %d", &c, &val);
		if (r == EOF && errno > 0)
			return -2;

		if (strcmp(c, "forward") == 0)
		{
			if (i == 0)
			{
				in[i].x = val;
				in[i].y = 0;
				in[i].aim = 0;
			}
			else
			{
				in[i].x = in[i - 1].x + val;
				in[i].y = in[i - 1].y + in[i - 1].aim * val;
				in[i].aim = in[i - 1].aim;
			}
		}

		int upFlg = strcmp(c, "up") == 0;
		int downFlg = strcmp(c, "down") == 0;
		if (upFlg || downFlg)
		{
			if (i == 0)
			{
				in[i].x = 0;
				in[i].y = 0;
				in[i].aim = val;
			}
			else
			{
				in[i].x = in[i - 1].x;
				in[i].y = in[i - 1].y;
				if (upFlg)
					in[i].aim = in[i - 1].aim - val;
				if (downFlg)
					in[i].aim = in[i - 1].aim + val;
			}
		}
	}

	printf("%d\n", in[l-1].x * in[l-1].y); // 1855892637
	return 0;
}
