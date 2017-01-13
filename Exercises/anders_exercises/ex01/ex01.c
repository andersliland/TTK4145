// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld_c.c -lpthread

#include <pthread.h>
#include <stdio.h>

int Num = 0;

// Note the return type: void*
void* someThreadFunction(){
    printf("Hello from a thread!\n");
    return NULL;

}

void* decNumFunc(){
	for(int j = 0; j < 1000000; j++)
		Num++;
	return NULL;
}

void* incNumFunc(){
	for(int j = 0; j < 1000000; j++)
		Num--;
	return NULL;
	
}


int main(){
    pthread_t someThread;
    // creates a thread (scheduled by the OS)
    pthread_create(&someThread, NULL, someThreadFunction, NULL);

    pthread_t decNumThread_id;
    pthread_create(&decNumThread_id, NULL, decNumFunc, NULL);

    pthread_t incNumThread_id;
    pthread_create(&incNumThread_id, NULL, incNumFunc, NULL);


	// suspends main thread until threads in argument is completed    
    pthread_join(someThread, NULL);
    pthread_join(decNumThread_id, NULL);
    pthread_join(incNumThread_id, NULL);

    printf("Hello from main!\n");
	printf("Num: %i \n", Num);    
    return 0;
    
}
