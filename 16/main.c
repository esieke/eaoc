#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"
#include "assert.h"

#define STR_LEN 500
int version = 0;

int aocOpen(const char *name, FILE **f)
{
	int nums = 0;
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
			break;
		nums++;
	}

	if (fseek(*f, 0, SEEK_SET) < 0)
		return -2;

	return nums * 4;
}

void initMsg(void *msg, int len, FILE *f)
{
	char(*m)[len] = msg;
	char buf = 0;
	int r = 0;
	for (int i = 0; i < len / 4; i++)
	{
		assert(fscanf(f, "%c", &buf) != EOF);
		char *b = "0000";
		switch (buf)
		{
		case '0':
			b = "0000";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case '1':
			b = "0001";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case '2':
			b = "0010";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case '3':
			b = "0011";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case '4':
			b = "0100";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case '5':
			b = "0101";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case '6':
			b = "0110";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case '7':
			b = "0111";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case '8':
			b = "1000";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case '9':
			b = "1001";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case 'A':
			b = "1010";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case 'B':
			b = "1011";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case 'C':
			b = "1100";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case 'D':
			b = "1101";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case 'E':
			b = "1110";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		case 'F':
			b = "1111";
			memcpy(&(*m)[i * 4], b, 4 * sizeof(char));
			break;
		default:
			assert(0);
			break;
		}
	}
}

int toInt(void *msg, int len)
{
	char(*m)[len] = msg;
	int ret = 0;

	for (int i = 0; i < len; i++)
	{
		if ((*m)[i] == '1')
			ret |= 1 << (len - 1 - i);
	}
	return ret;
}

int parseLiteral(void *msg, int len, int *pos)
{
	char(*m)[len] = msg;

	int resCtr = 0;
	char res[len];
	memset(res, 0, len * sizeof(char));

	int cont = 1;
	do
	{
		cont = toInt(&(*m)[*pos], 1);
		*pos += 1;
		memcpy(&res[resCtr * 4], &(*m)[*pos], 4 * sizeof(char));
		*pos += 4;
		resCtr++;
	} while (*pos < len && cont);

	int value = toInt(res, resCtr * 4);
	return value;
}

int parseLength(void *msg, int len, int *pos, int *id)
{
	char(*m)[len] = msg;

	*id = toInt(&(*m)[*pos], 1);
	*pos += 1;
	assert(*id || *id == 0);

	int ret;
	if (*id) // ID 1 -> 11 Bit number of sub packages
	{
		ret = toInt(&(*m)[*pos], 11);
		*pos += 11;
	}
	if (*id == 0) // ID 0 -> 15 Bit lenght of sub packages
	{
		ret = toInt(&(*m)[*pos], 15);
		*pos += 15;
	}
	return ret;
}

long int calc(int id, int a, int b)
{
	switch (id)
	{
	case 0:
		return a + b;
	case 1:
		return a * b;
	case 2:
		if (a < b)
			return a;
		return b;
	case 3:
		if (a > b)
			return a;
		return b;
	case 5:
		if (a > b)
			return 1;
		return 0;
	case 6:
		if (a < b)
			return 1;
		return 0;
	case 7:
		if (a == b)
			return 1;
		return 0;
	default:
		assert(0);
	}
	return 0;
}

long int parseHeader(void *msg, int len, int *pos)
{
	char(*m)[len] = msg;
	int ver = toInt(&(*m)[*pos], 3);
	*pos += 3;

	version += ver;

	int typeId = toInt(&(*m)[*pos], 3);
	*pos += 3;

	if (typeId == 4)
		return parseLiteral(msg, len, pos);

	int id = 0;
	int n = parseLength(msg, len, pos, &id);

	long int ret = 0;
	if (id == 1)
	{
		for (int i = 0; i < n; i++)
		{
			if (i > 0)
				ret = calc(typeId, ret, parseHeader(msg, len, pos));
			else
				ret = parseHeader(msg, len, pos);
		}
	}
	if (id == 0)
	{
		int spos = *pos, i = 0;
		while ((*pos - spos) < n)
		{
			if (i > 0)
				ret = calc(typeId, ret, parseHeader(msg, len, pos));
			else
				ret = parseHeader(msg, len, pos);
			i++;
		}
	}
	return ret;
}

int main()
{
	FILE *f;
	int l = aocOpen("../../16/input", &f);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	char msg[l];
	memset(msg, 0, l * sizeof(char));

	initMsg(msg, l, f);

	int pos = 0;
	int res = parseHeader(msg, l, &pos);
	printf("puzzle 1: %d\n", version);

	printf("puzzle 2: %d\n", res);

	return 0;
}
