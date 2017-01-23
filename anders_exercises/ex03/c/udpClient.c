#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <netinet/in.h>


#define PORT 1234
#define BUFLEN 1024
#define INET_ADDR "127.0.0.1"

int main()
{
  int clientSocket, nBytes;
  char buffer[BUFLEN];
  struct sockaddr_in serverAddr;
  socklen_t addr_size;


  // Create UDP socekt
  clientSocket = socket(PF_INET, SOCK_DGRAM, 0);

  // Configure settings in adress struct
  serverAddr.sin_family = AF_INET;
  serverAddr.sin_port = htons(PORT);
  serverAddr.sin_addr.s_addr = inet_addr(INET_ADDR);
  memset(serverAddr.sin_zero, '\0', sizeof(serverAddr.sin_zero));

  // Initialize size variable to be used later on
  addr_size = sizeof(serverAddr);


}
