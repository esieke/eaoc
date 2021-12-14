#include "stdio.h"
#include "errno.h"
#include "string.h"
#include "stdlib.h"
#include "assert.h"

#define STR_LEN 10
#define MAX_CHILDS 100
#define MAX_LOGS 100

typedef struct cave_
{
	char name[STR_LEN];
	char childs[MAX_CHILDS][STR_LEN];
} cave;

int desStr(const void *a, const void *b)
{
	return strcmp(a, b);
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

int check(cave *c, void *caves, int cavesLen, void *logs, int logsLen)
{
	cave(*cs)[cavesLen] = caves;
	char(*l)[logsLen][STR_LEN] = logs;

	int ret = 0;
	for (int i = 0; c->childs[i][0] != 0; i++)
	{
		int inLog = 0;
		for (int li = 0; li < logsLen; li++)
		{
			if (strcmp(c->childs[i], (*l)[li]) == 0)
				inLog = 1;
		}
		int isStart = 0;
		if (inLog == 0)
		{
			if (strcmp(c->childs[i], "start") == 0)
			{
				isStart = 1;
				ret++;
			}
		}
		// childCave
		if (inLog == 0 && isStart == 0)
		{
			// getCave
			cave nextC;
			for (int ci = 0; ci < cavesLen; ci++)
			{
				if (strcmp(c->childs[i], (*cs)[ci].name) == 0)
					nextC = (*cs)[ci];
			}
			// copy Log to new log
			char newLogs[logsLen][STR_LEN];
			memset(newLogs, 0, logsLen * STR_LEN * sizeof(char));
			for (int li = 0; li < logsLen; li++)
			{
				strcpy(newLogs[li], (*l)[li]);
				if ((*l)[li][0] == 0)
				{
					if (c->name[0] > 97 && c->name[0] < 123)
						strcpy(newLogs[li], c->name);
					break;
				}
			}
			ret += check(&nextC, caves, cavesLen, newLogs, logsLen);
		}
	}
	return ret;
}

int checkTwice(cave *c, void *caves, int cavesLen, void *logs, int logsLen, void *twice, int ctr)
{
	cave(*cs)[cavesLen] = caves;
	char(*l)[logsLen][STR_LEN] = logs;
	char(*t)[STR_LEN] = twice;
	int lctr = ctr;

	int ret = 0;
	for (int i = 0; c->childs[i][0] != 0; i++)
	{
		int inLog = 0;
		for (int li = 0; li < logsLen; li++)
		{
			if (strcmp(c->childs[i], (*l)[li]) == 0)
				inLog = 1;
		}
		if (*t != NULL)
		{
			if (strcmp(*t, c->childs[i]) == 0)
			{
				if (lctr == 0)
				{
					inLog = 0;
					t = NULL;
				}
				if (lctr > -1)
					lctr--;
			}
		}

		int isStart = 0;
		if (inLog == 0)
		{
			if (strcmp(c->childs[i], "start") == 0)
			{
					ret++;
				isStart = 1;
			}
		}
		// childCave
		if (inLog == 0 && isStart == 0)
		{
			// getCave
			cave nextC;
			for (int ci = 0; ci < cavesLen; ci++)
			{
				if (strcmp(c->childs[i], (*cs)[ci].name) == 0)
					nextC = (*cs)[ci];
			}
			// copy Log to new log
			char newLogs[logsLen][STR_LEN];
			memset(newLogs, 0, logsLen * STR_LEN * sizeof(char));
			for (int li = 0; li < logsLen; li++)
			{
				strcpy(newLogs[li], (*l)[li]);
				if ((*l)[li][0] == 0)
				{
					if (c->name[0] > 97 && c->name[0] < 123)
						strcpy(newLogs[li], c->name);
					break;
				}
			}
			ret += checkTwice(&nextC, caves, cavesLen, newLogs, logsLen, *t, lctr);
		}
	}

	return ret;
}

int main()
{
	FILE *f;
	int l = aocOpen("../../12/input", &f);
	if (l < 0)
	{
		printf("open input file failed with error\n");
		return 1;
	}

	char in[l][2][STR_LEN];
	memset(in, 0, l * 2 * STR_LEN * sizeof(char));
	for (int i = 0; i < l; i++)
	{
		int r = 0;
		while (1)
		{
			assert(r < STR_LEN);
			int buf = fgetc(f);
			if ((char)buf == '-')
			{
				in[i][0][r] = 0;
				break;
			}
			in[i][0][r] = (char)buf;
			r++;
		}
		r = 0;
		while (1)
		{
			assert(r < STR_LEN);
			int buf = fgetc(f);
			if ((char)buf == '\n' || buf == EOF)
			{
				in[i][1][r] = 0;
				break;
			}
			in[i][1][r] = (char)buf;
			r++;
		}
		// print raw input
		printf("%s-%s\n", in[i][0], in[i][1]);
	}

	// get all caves
	int allNodesLen = l * 2;
	char allNodes[allNodesLen][STR_LEN];
	memset(allNodes, 0, STR_LEN * allNodesLen * sizeof(char));
	for (int i = 0; i < allNodesLen; i++)
	{
		strcpy(allNodes[i], in[i / 2][i % 2]);
	}

	// number of caves
	qsort(allNodes, allNodesLen, STR_LEN * sizeof(char), desStr);
	int cavesLen = 1;
	for (int i = 1; i < allNodesLen; i++)
	{
		if (strcmp(allNodes[i], allNodes[i - 1]))
			cavesLen++;
	}

	// set cave name
	cave caves[cavesLen];
	memset(caves, 0, cavesLen * sizeof(cave));
	strcpy(caves[0].name, allNodes[0]);
	cavesLen = 1;
	for (int i = 1; i < allNodesLen; i++)
	{
		if (strcmp(allNodes[i], allNodes[i - 1]))
		{
			strcpy(caves[cavesLen].name, allNodes[i]);
			cavesLen++;
		}
	}

	// set cave childs
	for (int c = 0; c < cavesLen; c++)
	{
		int childLen = 0;
		for (int i = 0; i < l; i++)
		{
			assert(childLen < MAX_CHILDS);
			if (strcmp(caves[c].name, in[i][0]) == 0)
				strcpy(caves[c].childs[childLen++], in[i][1]);
			if (strcmp(caves[c].name, in[i][1]) == 0)
				strcpy(caves[c].childs[childLen++], in[i][0]);
		}
	}

	// getCave
	cave caveEnd;
	for (int ci = 0; ci < cavesLen; ci++)
	{
		if (strcmp(caves[ci].name, "end") == 0)
			caveEnd = caves[ci];
	}
	// copy Log to new log
	char newLog[MAX_LOGS][STR_LEN];
	memset(newLog, 0, MAX_LOGS * STR_LEN * sizeof(char));
	strcpy(newLog[0], caveEnd.name);
	int ret = 0;
	int retCheck = check(&caveEnd, caves, cavesLen, newLog, MAX_LOGS);
	ret += retCheck;
	for (int ci = 0; ci < cavesLen; ci++)
	{
		if (strcmp(caves[ci].name, "start") != 0 &&
			strcmp(caves[ci].name, "end") != 0 &&
			caves[ci].name[0] > 97 &&
			caves[ci].name[0] < 123)
		{
			char twice[STR_LEN];
			strcpy(twice, caves[ci].name);
			int ctr = 0;
			int ret_buf = -1;
			while (ctr < 10)
			{
				ret_buf = checkTwice(&caveEnd, caves, cavesLen, newLog, MAX_LOGS, twice, ctr);
				ctr++;
				ret += ret_buf - retCheck;
			}
		}
	}

	printf("%d\n", ret);
	return 0;
}
