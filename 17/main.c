#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"
#include "assert.h"

#define STR_LEN 500

#define TARGET_WIDTH 56

typedef struct pos_
{
	int y;
	int x;
} pos;

typedef struct target_
{
	pos b;
	pos e;
} target;

typedef struct probe_
{
	pos mx;
	pos p;
	pos v;
} probe;

void step(probe *p)
{
	p->p.y += p->v.y;
	if (p->p.y > p->mx.y)
		p->mx.y = p->p.y;
	p->v.y -= 1;

	p->p.x += p->v.x;
	if (p->p.x > p->mx.x)
		p->mx.x = p->p.x;
	if (p->v.x > 0)
		p->v.x -= 1;
	if (p->v.x < 0)
		p->v.x += 1;
}

int checkForHit(probe *p, target *t)
{
	if (p->p.x > t->e.x)
		return -1; // x out of range
	if (p->p.y < t->e.y)
		return -1; // y out of range
	if (p->p.x >= t->b.x && p->p.x <= t->e.x &&
		p->p.y <= t->b.y && p->p.y >= t->e.y)
		return 1;
	return 0;
}

int vx(probe *p)
{
	return p->v.x;
}

int vy(probe *p)
{
	return p->v.y;
}

int simu(target *t)
{
	int yMx = 0;

	for (int xi = 1; xi < t->e.x; xi++)
	{
		probe p = {.p.x = 0, .p.y = t->b.y, .v.y = 0, .v.x = xi, .mx.y = 0, .mx.x = 0};
		while (1)
		{
			step(&p);
			p.p.y = t->b.y;
			int cfh = checkForHit(&p, t);
			if (cfh > 0)
			{
				for (int yi = 1; yi < (-1*t->e.y +1); yi++)
				{
					probe p2 = {.p.x = 0, .p.y = 0, .v.y = yi, .v.x = xi, .mx.y = 0, .mx.x = 0};
					while (1)
					{
						step(&p2);
						int cfh2 = checkForHit(&p2, t);
						if (cfh2 > 0)
						{
							if (p2.mx.y > yMx)
								yMx = p2.mx.y;
						}
						if (cfh2 < 0)
							break;
					}
				}
			}
			if (cfh < 0 || vx(&p) <= 0)
				break;
		}
	}
	return yMx;
}

int main()
{
	// target area: x=20..30, y=-10..-5
	// target t = {.b = {.y = -5, .x = 20}, .e = {.y = -10, .x = 30}};
	// target area: x=94..151, y=-156..-103
	target t = {.b = {.y = -103, .x = 94},.e = {.y = -156, .x = 151}};

	int res = simu(&t);
	printf("%d\n", res);
	return 0;
}
