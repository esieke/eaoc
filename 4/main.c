#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"

#define STR_LEN 1000
#define IN_LEN 500

typedef struct field_
{
	int val;
	int id;
	int marked;
} field;

typedef struct board_
{
	field field[5][5];
	field acsField[5][5];
	int sumMarked;
	int won;
} board;

int desBoard(const void *a, const void *b)
{
	return (((board *)b)->sumMarked - ((board *)a)->sumMarked);
}

int ascField(const void *a, const void *b)
{
	return (((field *)a)->val - ((field *)b)->val);
}

int calcResult(board *b, int row, int col)
{
	b->won = 1;
	int res = 0;
	for (int r = 0; r < 5; r++)
	{
		for (int c = 0; c < 5; c++)
		{
			if (b->field[r][c].marked == 0)
				res += b->field[r][c].val;
		}
	}

	return res * b->field[row][col].val;
}

int checkResult(board *b, int row, int col)
{
	if (b->won)
		return -1;
	int check = 0;
	for (int c = 0; c < 5; c++)
	{
		if (b->field[row][c].marked)
			check++;
		if (check == 5)
			return calcResult(b, row, col);
	}
	check = 0;
	for (int r = 0; r < 5; r++)
	{
		if (b->field[r][col].marked)
			check++;
		if (check == 5)
			return calcResult(b, row, col);
	}
	return -1;
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
	int l = aocOpen("../../4/input", &f);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	int inLen = 0;
	int in[IN_LEN];
	int boardsLen = (l - 1) / 6;
	board boards[boardsLen];
	memset(&boards, 0, boardsLen * sizeof(board));

	for (int i = 0; i < 1 + boardsLen; i++)
	{
		if (i == 0)
		{
			// read line
			char buf[STR_LEN];
			char *cr = fgets(buf, STR_LEN, f);
			if (cr == NULL && errno > 0)
				return 5;
			if (strlen(buf) > STR_LEN - 1)
				return 6; // buffer limit STR_LEN

			char *t = strtok(buf, ",");
			while (t != NULL)
			{
				int r = sscanf(t, "%d,", &in[inLen]);
				if (r == EOF && errno > 0)
					return 5;
				inLen++;
				t = strtok(NULL, ",");
			}
		}
		else
		{
			field *a = (field *)(boards[i - 1].field);
			for (int k = 0; k < 25; k++)
			{
				int r = fscanf(f, "%d", &((a + k)->val));
				if (r == EOF && errno > 0)
					return 5;
				(a + k)->id = k;
			}
			memcpy(boards[i - 1].acsField, boards[i - 1].field, 25 * sizeof(field));
			qsort(boards[i - 1].acsField, 25, sizeof(field), ascField);
		}
	}

	int lastRes = 0;
	for (int i = 0; i < inLen; i++)
	{
		for (int k = 0; k < boardsLen; k++)
		{
			field *ret;
			field f = {.val = in[i]};
			ret = (field *)bsearch(&f, boards[k].acsField, 25, sizeof(field), ascField);
			if (ret != 0)
			{
				int r = (ret->id) / 5;
				int c = ret->id % 5;
				boards[k].field[r][c].marked = 1;
				boards[k].sumMarked++;
				if (boards[k].sumMarked > 4)
				{
					int res = checkResult(&boards[k], r, c);
					if (res > -1)
					{
						lastRes = res;
					}
				}
			}
		}
	}

	printf("%d\n", lastRes);
	return 0;
}
