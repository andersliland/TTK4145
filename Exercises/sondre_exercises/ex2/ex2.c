// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o ex2_c ex2.c -lpthread

#include <pthread.h>
#include <stdio.h>

int num = 0;

// Note the return type: void*
void* threadIncNum(void* mutex) {
	for(int i = 0; i < 1000000; i++) {
		pthread_mutex_lock(mutex);
		num++;
		pthread_mutex_unlock(mutex);
	}
	return NULL;
}

void* threadDecNum(void* mutex) {
	for(int i = 0; i < 1000001; i++) {
		pthread_mutex_lock(mutex);
		num--;
		pthread_mutex_unlock(mutex);
	} 
	return NULL;
}

int main() {
    pthread_mutex_t mutex;
    pthread_mutex_init(&mutex, NULL);

    pthread_t thread1, thread2;
    pthread_create(&thread1, NULL, threadIncNum, (void*) &mutex);
    pthread_create(&thread2, NULL, threadDecNum, (void*) &mutex);

    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);
    pthread_mutex_destroy(&mutex);

    printf("Num: %d\n", num);

    return 0;
}