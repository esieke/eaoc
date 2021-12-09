#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"
#include "assert.h"

#define STR_LEN 500

int counter_low = 1;

typedef struct pos_
{
	int x;
	int y;
} pos;

int des(const void *a, const void *b)
{
	return (*(u_int64_t *)b - *(u_int64_t *)a);
}

void checkBasin(void *grid, void *basin, int max_y, int max_x, int y, int x, int id)
{
	u_int8_t(*g)[max_y][max_x] = grid;
	u_int8_t(*b)[max_y][max_x] = basin;
	u_int8_t v = (*g)[y][x];

	(*b)[y][x] = id;
	for (int yy = y - 1; yy < y + 2; yy++)
	{
		if (yy != y && yy >= 0 && yy < max_y)
		{
			u_int8_t vv = (*g)[yy][x];
			if (vv > v && vv < 9)
			{
				assert((*b)[yy][x] == 0 || (*b)[yy][x] == id);
				(*b)[yy][x] = id;
				checkBasin(grid, basin, max_y, max_x, yy, x, id);
			}
		}
	}
	for (int xx = x - 1; xx < x + 2; xx++)
	{
		if (xx != x && xx >= 0 && xx < max_x)
		{
			u_int8_t vv = (*g)[y][xx];
			if (vv > v && vv < 9)
			{
				assert((*b)[y][xx] == 0 || (*b)[y][xx] == id);
				(*b)[y][xx] = id;
				checkBasin(grid, basin, max_y, max_x, y, xx, id);
			}
		}
	}
}

u_int64_t checkMinOfPos(void *grid, void *basin, int max_y, int max_x, int y, int x)
{
	u_int8_t(*g)[max_y][max_x] = grid;
	u_int8_t v = (*g)[y][x];
	for (int yy = y - 1; yy < y + 2; yy++)
	{
		if (yy != y && yy >= 0 && yy < max_y)
		{
			u_int8_t vv = (*g)[yy][x];
			if (vv <= v)
				return 0;
		}
	}
	for (int xx = x - 1; xx < x + 2; xx++)
	{
		if (xx != x && xx >= 0 && xx < max_x)
		{
			u_int8_t vv = (*g)[y][xx];
			if (vv <= v)
				return 0;
		}
	}

	checkBasin(grid, basin, max_y, max_x, y, x, counter_low);
	counter_low++;
	return (u_int64_t)(v + 1);
}

int checkMin(void *grid, void *basin, int max_y, int max_x)
{
	u_int64_t ret = 0;
	u_int8_t(*g)[max_y][max_x] = grid;
	for (int y = 0; y < max_y; y++)
	{
		u_int8_t buf = 0;
		for (int x = 0; x < max_x; x++)
		{
			ret += checkMinOfPos(grid, basin, max_y, max_x, y, x);
		}
	}
	return ret;
}

u_int64_t getResult(void *basin, int max_y, int max_x, int n)
{
	u_int64_t sums[n];
 	memset(sums, 0, n * sizeof(u_int64_t));

	u_int8_t(*b)[max_y][max_x] = basin;
	for (int y = 0; y < max_y; y++)
	{
		u_int8_t buf = 0;
		for (int x = 0; x < max_x; x++)
		{
			if((*b)[y][x] > 0)
				sums[(*b)[y][x]]++;
		}
	}
	qsort(sums, n, sizeof(u_int64_t), des);

	return sums[0] * sums[1] * sums[2];
}

int aocOpen(const char *name, FILE **f, pos *dim)
{
	int lines = 0;
	int rows = 0;
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
			(*dim).x = rows - 1;
			lines++;
			rows = 0;
		}
		if (r == EOF)
			break;
	}

	if (fseek(*f, 0, SEEK_SET) < 0)
		return -2;
	(*dim).y = lines;
	return lines;
}

int main()
{
	FILE *f;
	pos dim;
	int l = aocOpen("../../9/input", &f, &dim);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	u_int8_t grid[dim.y][dim.x];
	memset(grid, 0, dim.x * dim.x * sizeof(u_int8_t));
	u_int8_t basin[dim.y][dim.x];
	memset(basin, 0, dim.x * dim.x * sizeof(u_int8_t));

	for (int y = 0; y < dim.y; y++)
	{
		u_int8_t buf = 0;
		for (int x = 0; x < dim.x; x++)
		{
			buf = 0;
			int r = fscanf(f, "%c", &buf);
			assert(buf > 47 && buf < 58);
			grid[y][x] = buf - 48;
		}
		buf = 0;
		int r = fscanf(f, "%c", &buf);
		assert(buf == 10 || r == EOF);
	}

	int n = checkMin(&grid, &basin, dim.y, dim.x);

	u_int64_t res = getResult(&basin, dim.y, dim.x, n);

	printf("%lld\n", res);

	return 0;
}
