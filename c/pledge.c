#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <err.h>

int main()
{
	if (pledge("stdio", NULL) != 0) {
		err(1, "pledge");
	}
	getchar();
	exit(0);
}
