#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#define DEF_CARDS_NUM (52)

void swap(int* x, int* y) {
    int tmp = *x;
    *x = *y;
    *y = tmp;
}

int shuffle(int* arr, int len) {
    int pos = 0;
    for (int i=len-1; i>0; i--) {
        pos = random() % i;
        swap(arr+i, arr+pos);
    }
}

int main() {
    srandom(time(NULL));
    int cards[DEF_CARDS_NUM];
    for (int i=1; i<DEF_CARDS_NUM+1; i++) {
        cards[i-1] = i;
    }

    shuffle(cards, DEF_CARDS_NUM);

    for (int i=0; i<DEF_CARDS_NUM; i++) {
        printf("%02d ", cards[i]);
    }
    putchar('\n');

    return 0;
}
