#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"
#include "assert.h"

#define PARS_LEN 500

typedef struct dim_
{
	int lines;
	int rows;
} dim;

int des(const void *a, const void *b)
{
	int ret = 0;
	if (*(u_int64_t *)b > *(u_int64_t *)a)
		ret = 1;
	if (*(u_int64_t *)b < *(u_int64_t *)a)
		ret = -1;
	return ret;
}

int aocOpen(const char *name, FILE **f, dim *d)
{
	int lines = 0;
	int rows = 0;
	int rows_max = 0;
	char buf = 0;

	*f = fopen(name, "r");
	if (*f == NULL)
		return -1;

	while (1)
	{
		int r = fscanf(*f, "%c", &buf);
		rows++;
		if (r == EOF && errno > 0)
			return -2;
		if (buf == '\n' || r == EOF)
		{
			if (rows > rows_max)
				rows_max = rows;
			rows = 0;
			lines++;
		}
		if (r == EOF)
			break;
	}

	if (fseek(*f, 0, SEEK_SET) < 0)
		return -2;

	d->lines = lines;
	d->rows = rows_max;
	return 0;
}

int score(char c)
{
	int s = 0;
	if (c == ')')
		s = 3;
	if (c == ']')
		s = 57;
	if (c == '}')
		s = 1197;
	if (c == '>')
		s = 25137;
	return s;
}

int parser(void *in, int lines, int rows)
{
	char(*i)[lines][rows] = in;
	char p[PARS_LEN];
	u_int64_t ress[lines];
	int ressCtr = 0;
	memset(ress, 0, lines * sizeof(u_int64_t));
	for (int l = 0; l < lines; l++)
	{
		int pc = 0;
		for (int r = 0; r < strlen((*i)[l]); r++)
		{
			char c = (*i)[l][r];
			if (r == 0)
			{
				if (c == '(' || c == '[' || c == '{' || c == '<')
				{
					p[pc] = '('; // initial value can also be [ { <
					pc++;
					p[pc] = c;
				}
				else
				{
					pc = 0;
					break;
				}
			}
			else
			{
				if (p[pc] == '(' && c == ')')
					pc--;
				else if (p[pc] == '[' && c == ']')
					pc--;
				else if (p[pc] == '{' && c == '}')
					pc--;
				else if (p[pc] == '<' && c == '>')
					pc--;
				else if ((p[pc] == ')' || p[pc] == ']' || p[pc] == '}' || p[pc] == '>' ||
						  p[pc] == '(' || p[pc] == '[' || p[pc] == '{' || p[pc] == '<') &&
						 (c == '(' || c == '[' || c == '{' || c == '<'))
				{
					pc++;
					p[pc] = c;
				}
				else
				{
					pc = 0;
					break;
				}
			}
			assert(pc > -1);
			assert(pc < PARS_LEN);
		}
		if (pc > 0)
		{
			u_int64_t res = 0;
			for (int i = pc; i > 0; i--)
			{
				if (p[i] == '(')
					res = res * 5 + 1;
				if (p[i] == '[')
					res = res * 5 + 2;
				if (p[i] == '{')
					res = res * 5 + 3;
				if (p[i] == '<')
					res = res * 5 + 4;
			}
			ress[ressCtr] = res;
			ressCtr++;
		}
	}

	qsort(ress, lines, sizeof(u_int64_t), des);
	return ress[ressCtr / 2];
}

int main()
{
	FILE *f;
	dim d;
	int l = aocOpen("../../10/input", &f, &d);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	char in[d.lines][d.rows];
	memset(in, 0, d.lines * d.rows * sizeof(char));
	for (int l = 0; l < d.lines; l++)
	{
		for (int r = 0; r < d.rows; r++)
		{
			int buf = fgetc(f);
			if (buf == '\n' || buf == EOF)
				break;
			in[l][r] = (char)buf;
		}
	}

	printf("%d\n", parser(&in, d.lines, d.rows));

	return 0;
}
