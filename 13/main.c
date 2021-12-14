#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"
#include "assert.h"

#define STR_LEN 100

typedef struct pos_
{
	int y;
	int x;
} pos;

typedef struct inst_
{
	char dir;
	int n;
} inst;

typedef struct input_
{
	pos dim;
	pos dimAct;
	int posLen;
	void *pos; // pos[]
	void *map; // char[dim.y][dim.x]
	int instLen;
	void *inst; // intst[]
} input;

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

void parseInput(input *in, FILE *f)
{
	pos(*p)[in->posLen] = in->pos;
	inst(*i)[in->instLen] = in->inst;

	char buf[STR_LEN];
	int ci;

	int pi = 0;
	while (1)
	{
		ci = 0;
		memset(buf, 0, STR_LEN * sizeof(char));
		while (1)
		{
			char c = (char)fgetc(f);
			if (c == '\n')
				goto parse_folds;
			if (c == ',')
				break;
			buf[ci] = c;
			ci++;
		}
		int x;
		sscanf(buf, "%d", &x);

		ci = 0;
		memset(buf, 0, STR_LEN * sizeof(char));
		while (1)
		{
			char c = (char)fgetc(f);
			if (c == '\n')
				break;
			buf[ci] = c;
			ci++;
		}
		int y;
		sscanf(buf, "%d", &y);

		(*p)[pi].y = y;
		(*p)[pi].x = x;
		pi++;
	}
parse_folds:
	int ii = 0;
	while (1)
	{
		char dir = 0;
		ci = 0;
		memset(buf, 0, STR_LEN * sizeof(char));
		int c = 0;
		while (1)
		{
			c = fgetc(f);
			if (c == EOF)
				break;
			if ((char)c == '\n')
				break;
			if (dir > 0 && (char)c != '=')
			{
				buf[ci] = (char)c;
				ci++;
			}
			if ((char)c == 'y' || (char)c == 'x')
				dir = (char)c;
		}
		int n;
		sscanf(buf, "%d", &n);

		(*i)[ii].n = n;
		(*i)[ii].dir = dir;
		ii++;
		if (c == EOF)
			break;
	}
}

void setDim(input *in)
{
	pos(*p)[in->posLen] = in->pos;
	in->dim.y = 0;
	in->dim.x = 0;
	for (int i = 0; i < in->posLen; i++)
	{
		if ((*p)[i].y > in->dim.y)
			in->dim.y = (*p)[i].y;
		if ((*p)[i].x > in->dim.x)
			in->dim.x = (*p)[i].x;
	}
	in->dim.y++;
	in->dim.x++;
	in->dimAct.y = in->dim.y;
	in->dimAct.x = in->dim.x;
}

void initMap(input *in)
{
	pos(*p)[in->posLen] = in->pos;
	char(*m)[in->dim.y][in->dim.x] = in->map;

	for (int yi = 0; yi < in->dim.y; yi++)
	{
		for (int xi = 0; xi < in->dim.x; xi++)
		{
			for (int pi = 0; (*p)[pi].x > -1; pi++)
			{
				if ((*p)[pi].y == yi && (*p)[pi].x == xi)
					(*m)[yi][xi] = '#';
			}
		}
	}
}

void foldMap(input *in, int n)
{

	pos(*p)[in->posLen] = in->pos;
	inst(*i)[in->instLen] = in->inst;
	char(*m)[in->dim.y][in->dim.x] = in->map;

	if (!((*i)[n].dir == 'y' || (*i)[n].dir == 'x'))
		return;

	char mc[in->dim.y][in->dim.x];
	memcpy(mc, (*m), in->dim.y * in->dim.x * sizeof(char));
	memset((*m), '.', in->dim.y * in->dim.x * sizeof(char));

	if ((*i)[n].dir == 'y')
	{
		int up_dim = (*i)[n].n;
		int low_dim = (in->dimAct.y - (*i)[n].n - 1);
		int diff = 0;
		if (low_dim > up_dim)
			diff = low_dim - up_dim;
		for (int yi = 0; yi < up_dim; yi++)
		{
			for (int xi = 0; xi < in->dimAct.x; xi++)
			{
				(*m)[yi + diff][xi] = mc[yi][xi];
			}
		}
		if (up_dim > low_dim)
			diff = up_dim - low_dim;
		int offset = in->dimAct.y - 1;
		for (int yi = 0; yi < low_dim; yi++)
		{
			for (int xi = 0; xi < in->dimAct.x; xi++)
			{
				if ((*m)[yi + diff][xi] != '#')
					(*m)[yi + diff][xi] = mc[offset - yi][xi];
			}
		}
		in->dimAct.y = up_dim;
	}

	if ((*i)[n].dir == 'x')
	{
		int left_dim = (*i)[n].n;
		int right_dim = (in->dimAct.x - (*i)[n].n - 1);
		int diff = 0;
		if (right_dim > left_dim)
			diff = right_dim - left_dim;
		for (int yi = 0; yi < in->dimAct.y; yi++)
		{
			for (int xi = 0; xi < left_dim; xi++)
			{
				(*m)[yi][xi + diff] = mc[yi][xi];
			}
		}
		if ( left_dim > right_dim)
			diff = left_dim - right_dim;
		int offset = in->dimAct.x - 1;
		for (int yi = 0; yi < in->dimAct.y; yi++)
		{
			for (int xi = 0; xi < right_dim; xi++)
			{
				if ((*m)[yi][xi + diff] != '#')
					(*m)[yi][xi + diff] = mc[yi][offset - xi];
			}
		}
		in->dimAct.x = left_dim;
	}
}

void foldAllMap(input *in)
{
	for (int i = 0; i < in->instLen; i++)
	{
		foldMap(in, i);
	}
}

void printMap(input *in)
{
	char(*m)[in->dim.y][in->dim.x] = in->map;

	for (int yi = 0; yi < in->dimAct.y; yi++)
	{
		for (int xi = 0; xi < in->dimAct.x; xi++)
		{
			printf("%c", (*m)[yi][xi]);
		}
		printf("\n");
	}
}

int countMap(input *in)
{
	char(*m)[in->dim.y][in->dim.x] = in->map;
	int ret = 0;

	for (int yi = 0; yi < in->dimAct.y; yi++)
	{
		for (int xi = 0; xi < in->dimAct.x; xi++)
		{
			if ((*m)[yi][xi] == '#')
				ret++;
		}
	}

	return ret;
}

int main()
{
	FILE *f;
	int l = aocOpen("../../13/input", &f);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	input in;
	in.posLen = l;
	pos p[l];
	memset(p, -1, l * sizeof(pos));
	in.pos = &p;

	in.instLen = l;
	inst i[l];
	memset(i, 0, l * sizeof(inst));
	in.inst = i;

	parseInput(&in, f);
	setDim(&in);

	char map[in.dim.y][in.dim.x];
	memset(map, '.', in.dim.y * in.dim.x * sizeof(char));
	in.map = map;
	initMap(&in);
	foldAllMap(&in);

	printMap(&in);

	// ####...##..##..#..#...##..##...##..#..#.
	// #.......#.#..#.#..#....#.#..#.#..#.#..#.
	// ###.....#.#..#.####....#.#....#..#.####.
	// #.......#.####.#..#....#.#.##.####.#..#.
	// #....#..#.#..#.#..#.#..#.#..#.#..#.#..#.
	// #.....##..#..#.#..#..##...###.#..#.#..#.

	return 0;
}
