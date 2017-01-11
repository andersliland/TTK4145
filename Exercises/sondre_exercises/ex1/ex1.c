// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld_c.c -lpthread

#include <pthread.h>
#include <stdio.h>

int num = 0;

// Note the return type: void*
void* threadIncNum() {
	for(int i = 0; i < 1000000; i++) {
		num++;
	}
	return NULL;
}

void* threadDecNum() {
	for(int i = 0; i < 1000000; i++) {
		num--;
	} 
	return NULL;
}

int main(){
    pthread_t thread1;
    pthread_create(&thread1, NULL, threadIncNum, NULL);

    pthread_t thread2;
    pthread_create(&thread2, NULL, threadDecNum, NULL);

    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);

    printf("Num: %d\n", num);

    return 0;
}