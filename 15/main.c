#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"
#include "assert.h"

#define STR_LEN 500
#define UINT32_MAX 4294967295

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

	d->y = lines;
	d->x = rows_max - 1;
	return 0;
}

int main()
{
	FILE *f;
	dim d;
	int l = aocOpen("../../15/input", &f, &d);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	assert(d.y == d.x);
	u_int32_t in[d.y][d.x];
	memset(in, 0, d.y * d.x * sizeof(u_int32_t));
	u_int32_t cost[d.y][d.x];
	memset(cost, UINT32_MAX, d.y * d.x * sizeof(u_int32_t));
	for (int y = 0; y < d.y; y++)
	{
		for (int x = 0; x < d.x; x++)
		{
			int buf = fgetc(f);
			in[y][x] = (u_int32_t)(buf - 48);
		}
		int buf = fgetc(f);
		if (buf == EOF)
			break;
	}

	cost[d.y - 1][d.x - 1] = in[d.y - 1][d.x - 1];
	u_int32_t v = 0;
	for (u_int32_t i = 0; i < d.x - 1; i++)
	{
		u_int32_t corner = d.x - 2 - i;
		v = in[corner][corner];
		for (u_int32_t xi = d.x - 1; xi + 1 >= corner + 1; xi--)
		{
			if (xi > corner)
			{
				cost[corner][xi] = cost[corner + 1][xi] + in[corner][xi];
			}
		}
		for (u_int32_t xi = d.x - 1; xi + 1 >= corner + 1; xi--)
		{
			if (xi < d.x - 1)
			{
				u_int32_t v = cost[corner][xi + 1] + in[corner][xi];
				if (v < cost[corner][xi])
					cost[corner][xi] = v;
			}
		}

		for (u_int32_t yi = d.y - 1; yi + 1 >= corner + 1; yi--)
		{
			if (yi > corner)
			{
				cost[yi][corner] = cost[yi][corner + 1] + in[yi][corner];
			}
		}
		for (u_int32_t yi = d.y - 1; yi + 1 >= corner + 1; yi--)
		{
			if (yi < d.y - 1)
			{
				u_int32_t v = cost[yi + 1][corner] + in[yi][corner];
				if (v < cost[yi][corner])
					cost[yi][corner] = v;
			}
		}
	}

	// for (u_int32_t yi = 0; yi < d.y; yi++)
	// {
	// 	for (u_int32_t xi = 0; xi < d.x; xi++)
	// 	{
	// 		printf("%d\t", in[yi][xi]);
	// 	}
	// 	printf("\n");
	// }

	// printf("\n");

	// for (u_int32_t yi = 0; yi < d.y; yi++)
	// {
	// 	for (u_int32_t xi = 0; xi < d.x; xi++)
	// 	{
	// 		if (cost[yi][xi] == UINT32_MAX)
	// 			cost[yi][xi] = 0;
	// 		printf("%d\t", cost[yi][xi]);
	// 	}
	// 	printf("\n");
	// }

	printf("%d\n", cost[0][0] - in[0][0]);
	return 0;
}
