#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"
#include "assert.h"

#define STR_LEN 500
#define ELS_LEN 2

typedef struct pair_ pair;
typedef struct el_ el;

struct el_
{
	int v;
	pair *p;
};

struct pair_
{
	el *els;
	pair *parent;
};

pair *newPair(pair *parent)
{
	el *els = (el *)malloc(ELS_LEN * sizeof(el));
	assert(els != NULL);
	pair *p = (pair *)malloc(sizeof(pair));
	p->els = els;
	p->parent = parent;
	return p;
}

int pairGetParent(pair *p, pair **ret)
{
	if (p->parent == NULL)
		return -1;
	*ret = p->parent;
	return 0;
}

int pairGetElPair(pair *p, int idx, pair **ret)
{
	assert(idx >= 0 && idx < ELS_LEN);
	if (p->els[idx].p == NULL)
		return -1;
	*ret = p->els[idx].p;
	return 0;
}

int pairGetElValue(pair *p, int idx)
{
	assert(idx >= 0 && idx < ELS_LEN);
	if (p->els[idx].p != NULL)
		return -1;
	return p->els[idx].v;
}

void _pairSetEl(pair *p, pair *newP, int newV, int idx)
{
	assert(idx >= 0 && idx < ELS_LEN);
	p->els[idx].v = newV;
	p->els[idx].p = newP;
}

void pairSetElPair(pair *p, pair *new, int idx)
{
	_pairSetEl(p, new, -1, idx);
}

void pairSetElValue(pair *p, int new, int idx)
{
	_pairSetEl(p, NULL, new, idx);
}

int pairIsRegular(pair *p)
{
	int ret = 1;
	for (int i = 0; i < ELS_LEN; i++)
	{
		if (p->els[i].p != NULL)
		{
			ret = 0;
			break;
		}
	}
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

void pairPrint(pair *p)
{
	printf("[");
	int v = pairGetElValue(p, 0);
	if (v > -1)
		printf("%d,", v);
	pair *pb = NULL;
	v = pairGetElPair(p, 0, &pb);
	if (v > -1)
	{
		pairPrint(pb);
		printf(",");
	}
	v = pairGetElValue(p, 1);
	if (v > -1)
		printf("%d", v);
	v = pairGetElPair(p, 1, &pb);
	if (v > -1)
	{
		pairPrint(pb);
	}
	printf("]");
	if(p->parent == NULL)
		printf("\n");
}

int main()
{
	FILE *f;
	int l = aocOpen("../../18/input", &f);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	pair *head[l];
	for (int i; i < l; i++)
	{
		pair *p = NULL;
		pair *pb = NULL;
		int init = 1;
		char c, lastc = 0;
		while (1)
		{
			int r = fscanf(f, "%c", &c);
			if (c == '\n' || r == EOF)
				break;
			if (c == '[' && init == 0)
			{
				pb = newPair(p);
				int idx;
				assert(c == '[' || c == ',');
				if (lastc == '[')
					idx = 0;
				if (lastc == ',')
					idx = 1;
				pairSetElPair(p, pb, idx);
				pairGetElPair(p, idx, &p);
			}
			if (c == '[' && init != 0)
			{
				p = newPair(NULL);
				head[i] = p;
				init = 0;
			}
			if (c > 47 && c < 58 && lastc == '[')
				pairSetElValue(p, c - 48, 0);
			if (c > 47 && c < 58 && lastc == ',')
				pairSetElValue(p, c - 48, 1);
			if (c == ']')
				pairGetParent(p, &p);
			lastc = c;
		}
		assert(pairGetParent(p, &p) < 0);
	}

	for (int i = 0; i < l; i++)
	{
		pairPrint(head[i]);
	}
	
	
	return 0;
}
