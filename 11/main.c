#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"
#include "assert.h"

#define STR_LEN 500
#define STEPS 100
#define FLASH_LEVEL 10

typedef struct dim_
{
	int y;
	int x;
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
		if (r == EOF && errno > 0)
			return -2;
		if (buf == '\n' || r == EOF)
		{
			if ((rows - 1) > rows_max)
				rows_max = rows - 1;
			rows = 0;
			lines++;
		}
		if (r == EOF)
			break;
		rows++;
	}

	if (fseek(*f, 0, SEEK_SET) < 0)
		return -2;

	d->y = lines;
	d->x = rows_max;
	return 0;
}

int lim(int v, int max)
{
	if (v < 0)
		v = 0;
	if (v > max)
		v = max;
	return v;
}

void incr(void *grid, int y_max, int x_max)
{
	int(*g)[y_max][x_max] = grid;

	for (int y = 0; y < y_max; y++)
	{
		for (int x = 0; x < x_max; x++)
		{
			if ((*g)[y][x] < FLASH_LEVEL)
				(*g)[y][x]++;
		}
	}
}

void incrSlice(void *grid, int y_max, int x_max, int y_in, int x_in)
{
	int(*g)[y_max][x_max] = grid;
	int ys = lim(y_in - 1, y_max);
	int ye = lim(y_in + 2, y_max);
	int xs = lim(x_in - 1, x_max);
	int xe = lim(x_in + 2, x_max);

	(*g)[y_in][x_in]++;
	for (int y = ys; y < ye; y++)
	{
		for (int x = xs; x < xe; x++)
		{
			if ((*g)[y][x] < FLASH_LEVEL)
			{
				(*g)[y][x]++;
			}
		}
	}
}

int flash(void *grid, int y_max, int x_max)
{
	int(*g)[y_max][x_max] = grid;
	int ret = 0;

	for (int y = 0; y < y_max; y++)
	{
		for (int x = 0; x < x_max; x++)
		{
			if ((*g)[y][x] >= FLASH_LEVEL)
			{
				ret++;
				(*g)[y][x] = 0;
			}
		}
	}
	return ret;
}

int incrAdjacent(void *grid, int y_max, int x_max)
{
	int(*g)[y_max][x_max] = grid;

	int done = 0;
	while (done < 1)
	{
		done = 1;
		for (int y = 0; y < y_max; y++)
		{
			for (int x = 0; x < x_max; x++)
			{
				if ((*g)[y][x] == FLASH_LEVEL)
				{
					incrSlice(grid, y_max, x_max, y, x);
					done = 0;
				}
			}
		}
	}
}

int main()
{
	FILE *f;
	dim d;
	int l = aocOpen("../../11/input", &f, &d);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	int grid[d.y][d.x];
	memset(grid, 0, d.y * d.x * sizeof(char));
	for (int l = 0; l < d.y; l++)
	{
		for (int r = 0; r < d.x + 1; r++)
		{
			int buf = fgetc(f);
			if (buf == '\n' || buf == EOF)
				break;
			grid[l][r] = (char)buf - 48;
		}
	}

	int ret = 0;
	for (int s = 0; s < STEPS; s++)
	{
		incr(&grid, d.y, d.x);
		incrAdjacent(&grid, d.y, d.x);

		ret += flash(&grid, d.y, d.x);
	}

	printf("%d\n", ret);
	
	return 0;
}
