#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"

#define STR_BUF_LEN 200
#define STR_LEN 2500

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

int diff(int a, int b)
{
	int res = a - b;
	if (res < 0)
		res *= -1;
	return res;
}

int main()
{
	FILE *f;
	int l = aocOpen("../../7/input", &f);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	int in[STR_LEN];
	memset(in, 0, STR_LEN * sizeof(int));
	char buf[STR_BUF_LEN];
	int inLen = 0, bufLen = 0;
	while (1)
	{

		int r = fscanf(f, "%c", &buf[bufLen]);
		if (r == EOF && errno > 0)
			return 5;
		if (buf[bufLen] == ',' || r == EOF)
		{
			if (bufLen > STR_BUF_LEN - 1)
				return 1;
			buf[bufLen + 1] = 0;
			in[inLen] = atoi(buf);
			bufLen = -1;
			inLen++;
			if (r == EOF)
				break;
		}
		bufLen++;
	}

	int hPosMax = 0;
	for (int i = 0; i < inLen; i++)
	{
		if (in[i] > hPosMax)
			hPosMax = in[i];
	}

	int fuleCostDimX = inLen;
	int fuleCostDimY = hPosMax + 1;
	int fuelCost[fuleCostDimX][fuleCostDimY];
	memset(fuelCost, 0, fuleCostDimX * fuleCostDimY * sizeof(int));

	// cost map
	for (int x = 0; x < fuleCostDimX; x++)
	{
		for (int y = 0; y < fuleCostDimY; y++)
		{
			fuelCost[x][y] = diff(in[x], y);
		}
	}

	// sum fuel costs
	int fuelCostSum[fuleCostDimY];
	memset(fuelCostSum, 0, fuleCostDimY * sizeof(int));
	for (int y = 0; y < fuleCostDimY; y++)
	{
		for (int x = 0; x < fuleCostDimX; x++)
		{
			fuelCostSum[y] += fuelCost[x][y];
		}
	}

	// search min
	int min = fuelCostSum[0];
	for (int y = 0; y < fuleCostDimY; y++)
	{
		if( fuelCostSum[y] < min)
			min = fuelCostSum[y];
	}
	
	printf("%d\n", min);
	return 0;
}
