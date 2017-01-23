// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld_c.c -lpthread

#include <pthread.h>
#include <stdio.h>

int Num = 0;
pthread_mutex_t lock;

void* decNumFunc(){
	for(int j = 0; j < 10; j++)
	{
		pthread_mutex_lock(&lock);

		Num--;
		printf("decNum: %i \n", Num);
		pthread_mutex_unlock(&lock);
	}
	return NULL;
}

void* incNumFunc(){
	for(int j = 0; j < 11; j++)
	{
		pthread_mutex_lock(&lock);
		Num++;
		printf("incNum: %i \n", Num);
		pthread_mutex_unlock(&lock);
	}
	return NULL;

}


int main(){

	if(pthread_mutex_init(&lock, NULL) != 0){
		printf("Mutex init failed \n");
		return 1;
	}

   // creates a thread (scheduled by the OS)
    pthread_t decNumThread_id;
    pthread_create(&decNumThread_id, NULL, decNumFunc, NULL);

    pthread_t incNumThread_id;
    pthread_create(&incNumThread_id, NULL, incNumFunc, NULL);

	// suspends main thread until threads in argument is completed
    pthread_join(decNumThread_id, NULL);
    pthread_join(incNumThread_id, NULL);

    //destroys the mutex
    pthread_mutex_destroy(&lock);

	printf("Num: %i \n", Num);
    return 0;

}
