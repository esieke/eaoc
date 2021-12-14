#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"
#include "assert.h"

#define STR_LEN 500
#define SEG_LEN 8
#define SEG_N_STATES 11
#define DISP_LEN 4

typedef struct display_
{
	int ids[SEG_N_STATES];
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

int findIn(char *a, char *b)
{
	if (strlen(a) > strlen(b))
		return 0;
	int ret = 0;

	for (int ai = 0; ai < strlen(a); ai++)
	{
		int found = 0;
		for (int bi = 0; bi < strlen(b); bi++)
		{
			if (a[ai] == b[bi])
				found++;
		}
		ret += found;
	}

	if (ret == strlen(a) && ret == strlen(b))
		return 2;
	if (ret == strlen(a))
		return 1;
	return 0;
}

int getPatternById(display *d, int id)
{
	display disp = *d;

	for (int i = 0; i < SEG_N_STATES; i++)
	{
		if (disp.ids[i] == id)
			return i;
	}
	return -1;
}

void setBD(display *d)
{
	char *one = d->pattern[getPatternById(d, 1)];
	char *four = d->pattern[getPatternById(d, 4)];

	int p = 0;
	int f = 0;
	for (int f = 0; f < strlen(four); f++)
	{

		int found = 0;
		for (int o = 0; o < strlen(one); o++)
		{
			if (four[f] == one[o])
				found = 1;
		}
		if (found == 0)
		{
			d->pattern[SEG_N_STATES - 1][p] = four[f];
			p++;
		}
	}
	d->pattern[SEG_N_STATES - 1][p] = 0;
	d->ids[SEG_N_STATES - 1] = 10;
}

void setIdIfLenAndNotId(display *d, int id, int len, int notId)
{
	char *cmp = d->pattern[getPatternById(d, notId)];
	int notFound = 0;
	for (int p = 0; p < SEG_N_STATES - 1; p++)
	{
		if (strlen(d->pattern[p]) == len)
		{
			if (findIn(cmp, d->pattern[p]) == 0)
			{
				d->ids[p] = id;
				notFound++;
			}
		}
	}
	assert(notFound == 1);
}

void setIdIfLenAndId(display *d, int id, int len, int notId)
{
	char *cmp = d->pattern[getPatternById(d, notId)];
	int found = 0;
	for (int p = 0; p < SEG_N_STATES - 1; p++)
	{
		if (strlen(d->pattern[p]) == len)
		{
			if (findIn(cmp, d->pattern[p]) > 0)
			{
				d->ids[p] = id;
				found++;
			}
		}
	}
	assert(found == 1);
}

void setIdIfLenAndNotSet(display *d, int id, int len)
{
	int found = 0;
	for (int p = 0; p < SEG_N_STATES - 1; p++)
	{
		if (strlen(d->pattern[p]) == len)
		{
			if (d->ids[p] < 0)
			{
				d->ids[p] = id;
				found++;
			}
		}
	}
	assert(found == 1);
}

int getOutput(row *r)
{
	int ret = 0;
	int fac = 1000;
	for (int di = 0; di < DISP_LEN; di++)
	{
		for (int pi = 0; pi < SEG_N_STATES - 1; pi++)
		{
			if (findIn(r->out.fourDigit[di], r->disp.pattern[pi]) == 2)
			{
				assert(fac < 10000);
				// printf("%d", r->disp.ids[pi]);
				ret += r->disp.ids[pi] * fac;
				if (fac < 10)
					fac = 100000;
				fac /= 10;
			}
		}
	}
	// printf("\n");
	return ret;
}

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

	// set unique ids
	int num = 0;
	for (int i = 0; i < l; i++)
	{
		rows[i].disp.ids[SEG_N_STATES - 1] = -1;
		for (int k = 0; k < (SEG_N_STATES - 1); k++)
		{
			rows[i].disp.ids[k] = -1;
			int len = strlen(rows[i].disp.pattern[k]);
			if (len == 2)
				rows[i].disp.ids[k] = 1;
			if (len == 3)
				rows[i].disp.ids[k] = 7;
			if (len == 4)
				rows[i].disp.ids[k] = 4;
			if (len == 7)
				rows[i].disp.ids[k] = 8;
		}
	}

	int res = 0;
	for (int i = 0; i < l; i++)
	{
		setBD(&rows[i].disp);
		setIdIfLenAndNotId(&rows[i].disp, 6, 6, 1);
		setIdIfLenAndNotId(&rows[i].disp, 0, 6, 10);
		setIdIfLenAndNotSet(&rows[i].disp, 9, 6);
		setIdIfLenAndId(&rows[i].disp, 5, 5, 10);
		setIdIfLenAndId(&rows[i].disp, 3, 5, 1);
		setIdIfLenAndNotSet(&rows[i].disp, 2, 5);
		res += getOutput(&rows[i]);
	}
	printf("%d\n", res);
	return 0;
}
