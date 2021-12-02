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
			in[i].x = val;
			in[i].y = 0;
		}
		if (strcmp(c, "down") == 0)
		{
			in[i].x = 0;
			in[i].y = val;
		}
		if (strcmp(c, "up") == 0)
		{
			in[i].x = 0;
			in[i].y = -val;
		}
	}

	struct pos posAct = {.x = 0, .y = 0};
	for (int i = 0; i < l; i++)
	{
		posAct.x += in[i].x;
		posAct.y += in[i].y;
	}
	
	printf("%d\n", posAct.x * posAct.y);
	return 0;
}
