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
	int res = 0;
	char(*i)[lines][rows] = in;
	char p[PARS_LEN];
	for (int l = 0; l < lines; l++)
	{
		int pc = 0;
		for (int r = 0; r < strlen((*i)[l]); r++)
		{
			char c = (*i)[l][r];
			if (r == 0)
			{
				if (
					c == '(' ||
					c == '[' ||
					c == '{' ||
					c == '<')
				{
					pc++;
					p[pc] = c;
				}
				else
				{
					res += score(c);
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
					res += score(c);
					break;
				}
			}
			assert(pc > -1);
			assert(pc < PARS_LEN);
		}
		if( pc > 0 )
			printf("incomplete\n");
	}

	return res;
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

		// char *cr = fgets(in[l], d.rows, f);
		// assert(cr != NULL);
	}

	printf("%d\n", parser(&in, d.lines, d.rows));

	return 0;
}
