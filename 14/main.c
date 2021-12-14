#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"
#include "assert.h"

#define STR_LEN 500
#define UINT8_SIZE 256
#define UINT64_SIZE 65536

int des(const void *a, const void *b)
{
	if (*(u_int64_t *)b > *(u_int64_t *)a)
		return 1;
	if (*(u_int64_t *)b < *(u_int64_t *)a)
		return -1;
	return 0;
}

int asc(const void *a, const void *b)
{
	if (*(u_int64_t *)b > *(u_int64_t *)a && *(u_int64_t *)a == 0)
		return 1;
	if (*(u_int64_t *)b > *(u_int64_t *)a)
		return -1;
	if (*(u_int64_t *)b < *(u_int64_t *)a && *(u_int64_t *)b == 0)
		return -1;
	if (*(u_int64_t *)b < *(u_int64_t *)a)
		return 1;
	return 0;
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
	int l = aocOpen("../../14/input", &f);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	u_int8_t tmp[STR_LEN];
	memset(tmp, 0, STR_LEN * sizeof(u_int8_t));
	u_int64_t polymer[UINT64_SIZE];
	memset(polymer, 0, UINT64_SIZE * sizeof(u_int64_t));
	u_int64_t polymerBuf[UINT64_SIZE];
	memset(polymerBuf, 0, UINT64_SIZE * sizeof(u_int64_t));
	u_int64_t sum[UINT8_SIZE];
	memset(sum, 0, UINT8_SIZE * sizeof(u_int64_t));
	u_int8_t rules[UINT64_SIZE];
	memset(rules, 0, UINT64_SIZE * sizeof(u_int8_t));

	u_int16_t i = 0;
	u_int8_t buf = 0;
	u_int8_t state = 0;
	u_int16_t pair = 0;
	while (1)
	{
		int r = fscanf(f, "%c", &buf);
		if (r == EOF)
			break;
		if (state == 2)
		{
			if (i == 8)
				i = 0;
			if (i == 0)
				pair |= (u_int16_t)buf << 8;
			if (i == 1)
				pair |= (u_int16_t)buf;
			if (i == 6)
			{
				rules[pair] = buf;
				pair = 0;
			}
			i++;
		}
		if (state == 1)
		{
			state = 2;
		}
		if (state == 0)
		{
			if (buf == '\n')
			{
				state = 1;
				i = 0;
			}
			else
			{
				tmp[i] = buf;
				i++;
			}
		}
	}

	// initialize sum
	for (u_int8_t i = 0; i < strlen(tmp); i++)
	{
		sum[tmp[i]]++;
	}

	// initialize polymer
	for (u_int8_t i = 0; i < strlen(tmp) - 1; i++)
	{
		u_int16_t r = 0;
		r |= (u_int16_t)tmp[i] << 8;
		r |= (u_int16_t)tmp[i + 1];
		polymer[r]++;
	}

	u_int16_t polymerMin = ((u_int16_t)'A' << 8) | ((u_int16_t)'A');
	u_int16_t polymerMax = (((u_int16_t)'X' << 8) | ((u_int16_t)'X')) + 1;
	for (int step = 0; step < 40; step++)
	{
		for (u_int16_t pi = polymerMin; pi < polymerMax; pi++)
		{
			if (polymer[pi] != 0)
			{
				u_int64_t n = polymer[pi];
				u_int8_t ll = (u_int8_t)(pi >> 8);
				u_int8_t rl = (u_int8_t)pi;
				u_int8_t nl = rules[pi];
				sum[nl] += n;
				polymerBuf[((u_int16_t)ll << 8) | ((u_int16_t)nl)] += n;
				polymerBuf[((u_int16_t)nl << 8) | ((u_int16_t)rl)] += n;
			}
		}
		memcpy(polymer, polymerBuf, UINT64_SIZE * sizeof(u_int64_t));
		memset(polymerBuf, 0, UINT64_SIZE * sizeof(u_int64_t));
	}

	qsort(sum, UINT8_SIZE, sizeof(u_int64_t), des);
	u_int64_t max = sum[0];
	qsort(sum, UINT8_SIZE, sizeof(u_int64_t), asc);
	u_int64_t min = sum[0];

	// this was awesome
	printf("%lld\n", max - min);

	return 0;
}
