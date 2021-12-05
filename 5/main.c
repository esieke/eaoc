#include "stdio.h"
#include "errno.h"
#include "string.h"

#define STR_LEN 500

typedef struct pos_
{
	int x;
	int y;
} pos;

typedef struct line_
{
	pos begin;
	pos end;
	pos v;
} line;

typedef struct map_
{
	pos dim;
	void *vents; // 2Dim Array
} map;

int setV(line *l)
{
	int dX = l->end.x - l->begin.x;
	int dXAbs = dX < 0 ? dX * -1 : dX;
	int dY = l->end.y - l->begin.y;
	int dYAbs = dY < 0 ? dY * -1 : dY;

	if (!(
			l->begin.x == l->end.x ||
			l->begin.y == l->end.y ||
			(dXAbs == dYAbs && dXAbs != 0)))
		return -1;

	if (l->begin.x == l->end.x)
	{
		l->v.y = dY / dYAbs;
		l->v.x = 0;
		return 0;
	}

	if (l->begin.y == l->end.y)
	{
		l->v.x = dX / dXAbs;
		l->v.y = 0;
		return 0;
	}

	l->v.x = dX / dXAbs;
	l->v.y = dY / dYAbs;

	return 0;
}

void drawLine(map *m, line *l)
{
	if (setV(l) < 0)
		return;

	int (*vents)[m->dim.x][m->dim.y] = m->vents;

	pos p = l->end;
	(*vents)[p.x][p.y] += 1;
	p = l->begin;
	while (!(p.x == l->end.x && p.y == l->end.y))
	{
		(*vents)[p.x][p.y] += 1;
		p.x += l->v.x;
		p.y += l->v.y;
	}
}

void printMap(map *m)
{
	int (*vents)[m->dim.x][m->dim.y] = m->vents;
	for (int x = 0; x < m->dim.x; x++)
	{
		for (int y = 0; y < m->dim.y; y++)
		{
			if ((*vents)[x][y] == 0)
				printf(".");
			else
				printf("%d",(*vents)[x][y]);
		}
		printf("\n");
	}
}

int result(map *m)
{
	int (*vents)[m->dim.x][m->dim.y] = m->vents;
	int res = 0;
	for (int x = 0; x < m->dim.x; x++)
	{
		for (int y = 0; y < m->dim.y; y++)
		{
			if ((*vents)[x][y] > 1)
				res++;
		}
	}
	return res;
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
	int l = aocOpen("../../5/input", &f);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	line lines[l];
	memset(lines, 0, l * sizeof(line));
	map map;
	memset(&map, 0, sizeof(map));
	for (int i; i < l; i++)
	{
		// read line
		char buf[STR_LEN];
		char *cr = fgets(buf, STR_LEN, f);
		if (cr == NULL && errno > 0)
			return 5;
		if (strlen(buf) > STR_LEN - 1)
			return 6; // buffer limit STR_LEN

		int r = sscanf(buf, "%d,%d -> %d,%d",
					   &lines[i].begin.x,
					   &lines[i].begin.y,
					   &lines[i].end.x,
					   &lines[i].end.y);

		if (lines[i].begin.x > map.dim.x)
			map.dim.x = lines[i].begin.x;
		if (lines[i].begin.y > map.dim.y)
			map.dim.y = lines[i].begin.y;
		if (lines[i].end.x > map.dim.x)
			map.dim.x = lines[i].end.x;
		if (lines[i].end.y > map.dim.y)
			map.dim.y = lines[i].end.y;

		if (r == EOF && errno > 0)
			return -2;
	}
	map.dim.x++;
	map.dim.y++;

	int vents[map.dim.x][map.dim.y];
	memset(vents, 0, map.dim.x * map.dim.y * sizeof(int));
	map.vents = &vents;

	for (int i = 0; i < l; i++)
	{
		drawLine(&map, &lines[i]);
	}

	// printMap(&map);

	printf("%d\n", result(&map));

	return 0;
}
