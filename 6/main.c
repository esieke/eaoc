#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"

#define STR_LEN 1000
#define CLU_LEN 7

typedef struct clu_
{
	u_int8_t timer;
	u_int64_t num;
	u_int64_t childs;
} clu;

void initClus(void *clus)
{
	clu(*c)[CLU_LEN] = clus;
	for (int i = 0; i < CLU_LEN; i++)
	{
		(*c)[i].timer = i;
	}
}

void addNum(void *clus, u_int64_t t, u_int64_t n)
{
	clu(*c)[CLU_LEN] = clus;
	for (int i = 0; i < CLU_LEN; i++)
	{
		if (t == (*c)[i].timer)
		{
			(*c)[i].num += n;
			return;
		}
	}
}

u_int64_t getNum(void *clus, u_int64_t t){
		clu(*c)[CLU_LEN] = clus;
	for (int i = 0; i < CLU_LEN; i++)
	{
		if (t == (*c)[i].timer)
		{
			return (*c)[i].num;
		}
	}
}

void addChild(void *clus, u_int64_t t, u_int64_t child)
{
	clu(*c)[CLU_LEN] = clus;
	for (int i = 0; i < CLU_LEN; i++)
	{
		if (t == (*c)[i].timer)
		{
			(*c)[i].childs = child;
			return;
		}
	}
}

void childToNum(void *clus, u_int64_t t)
{
	clu(*c)[CLU_LEN] = clus;
	for (int i = 0; i < CLU_LEN; i++)
	{
		if (t == (*c)[i].timer)
		{
			(*c)[i].num += (*c)[i].childs;
			(*c)[i].childs = 0;
			return;
		}
	}
}

void simDay(void *clus)
{
	clu(*c)[CLU_LEN] = clus;
	for (int i = 0; i < CLU_LEN; i++)
	{
		if ((*c)[i].timer > 0)
			(*c)[i].timer--;
		else
			(*c)[i].timer = 6;
	}
	addChild(clus, 1, getNum(clus, 6));
	childToNum(clus, 5);
}

u_int64_t getPopulation(void *clus)
{
	u_int64_t ret = 0;
	clu(*c)[CLU_LEN] = clus;
	for (int i = 0; i < CLU_LEN; i++)
	{
		ret += (*c)[i].num + (*c)[i].childs;
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

int main()
{
	FILE *f;
	int l = aocOpen("../../6/input", &f);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	u_int64_t in[STR_LEN];
	memset(in, 0, STR_LEN * sizeof(int));
	char buf[STR_LEN];
	memset(buf, 0, STR_LEN * sizeof(char));
	char *cr = fgets(buf, STR_LEN, f);
	if (cr == NULL && errno > 0)
		return 5;
	if (strlen(buf) > STR_LEN - 1)
		return 6; // buffer limit STR_LEN

	int inLen = 0;

	char *t = strtok(buf, ",");
	while (t != NULL)
	{
		int r = sscanf(t, "%d,", &in[inLen]);
		if (r == EOF && errno > 0)
			return 5;
		inLen++;
		t = strtok(NULL, ",");
	}

	clu clus[CLU_LEN];
	memset(clus, 0, CLU_LEN * sizeof(clu));

	initClus(&clus);

	for (int i = 0; i < inLen; i++)
	{
		addNum(&clus, in[i], 1);
	}

	for (int i = 0; i < 256; i++)
	{
		simDay(&clus);
	}

	printf("%llu\n", getPopulation(&clus));

	return 0;
}
