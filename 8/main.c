#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"

#define STR_LEN 500
#define SEG_LEN 8
#define SEG_N_STATES 10
#define DISP_LEN 4

typedef struct display_
{
	int id;
	char pattern[SEG_N_STATES][SEG_LEN];
} display;

typedef struct output_
{
	char fourDigit[DISP_LEN][SEG_LEN];
} output;

typedef struct row_
{
	display disp;
	output out;
} row;

int aocOpen(const char *name, FILE **f)
{
	int lines = 0;
	char buf = 0;

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
	int l = aocOpen("../../8/input", &f);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	// read input
	row rows[l];
	memset(rows, 0, l * sizeof(row));
	for (int i; i < l; i++)
	{
		int r = fscanf(f, "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
					   &rows[i].disp.pattern[0],
					   &rows[i].disp.pattern[1],
					   &rows[i].disp.pattern[2],
					   &rows[i].disp.pattern[3],
					   &rows[i].disp.pattern[4],
					   &rows[i].disp.pattern[5],
					   &rows[i].disp.pattern[6],
					   &rows[i].disp.pattern[7],
					   &rows[i].disp.pattern[8],
					   &rows[i].disp.pattern[9],
					   &rows[i].out.fourDigit[0],
					   &rows[i].out.fourDigit[1],
					   &rows[i].out.fourDigit[2],
					   &rows[i].out.fourDigit[3]);
		if (r == EOF && errno > 0)
			return -2;
	}

	// count unique numbers
	int num = 0;
	for (int i = 0; i < l; i++)
	{
		for (int k = 0; k < DISP_LEN; k++)
		{
			int len = strlen(rows[i].out.fourDigit[k]);
			if (len == 2 || len == 3 || len == 4 || len == 7)
			{
				num += 1;
			}
		}
	}

	printf("%d\n", num);
	return 0;
}
