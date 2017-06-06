#include <pthread.h>
#include <stdlib.h>
#include <stdio.h>

typedef unsigned int boolean;
#define false 0
#define true (!false)


typedef struct {
    boolean first_process_active;
    boolean second_process_active;
    boolean fist_process_finished;
    boolean second_process_finished;

    pthread_mutexattr_t mutex_att;
    pthread_mutex_t mutex;

    pthread_condattr_t cond_att;
    pthread_cond_t no_first_process, no_second_process;
    pthread_cond_t atomic_action_ends1, atomic_action_ends2;
} monitor;


void entry_1(monitor *M){
    PTHREAD_MUTEX_LOCK((*M).mutex);
    //PTHREAD_MUTEX_LOCK(&M->mutex);
        while(M->first_process_active)
            PTHREAD_COND_WAIT(&M->no_first_process, &M->mutex);
        M->first_process_active = true;
}

void exit_1(monitor *M){
    while(!M->second_process_finished)
        PTHREAD_COND_WAIT(&M->atomic_action_ends2, &M->mutex);
    PTHREAD_COND_SIGNAL(&M->atomic_action_ends1);
    M->first_process_active = false;
    PTHREAD_COND_SIGNAL(&M->no_first_process);
    PTHREAD_MUTEX_UNLOCK(&M->mutex);
}


int main(){
    printf("Hello world\n");
}
