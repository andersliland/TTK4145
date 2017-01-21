#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>


#define BUFLEN 1024
#define NPACK 10
#define PORT 9930
#define localhost "127.0.0.1"


int main(void)
{
  int udpSocket, nBytes;
  char buffer[BUFLEN];
  struct sockaddr_in serverAddr, clientAddr;
  struct sockaddr_storage serverStorage;
  socklen_t addr_size, client_addr_size;
  int i;

  // Create UDP socket
  udpSocket = socket(PF_INET, SOCK_DGRAM, 0);

  //Configure settings in address struct
  serverAddr.sin_family = AF_INET;
  serverAddr.sin_port = htons(PORT);
  serverAddr.sin_addr.s_addr = inet_addr(localhost);
  memset(serverAddr.sin_zero, '\0', sizeof(serverAddr.sin_zero));

  // Bind socket with address struct
  bind(udpSocket, (struct sockaddr *) $serverAddr, sizeof(serverAddr));

  //Initialize size variable to be used later on
  addr_size = sizeof(serverStorage);

  while(1)
  {
    // Try to recieve any incomming UDP datagram.
    // Adress and port of requestin client will be stored
    // on serverStorage variable.
    nBytes = recvfrom(udpSocket, buffer, BUFLEN, 0,
      (struct sockaddr *) &serverStorage, &addr_size);

    // Convert message received to uppdercase
    for(i = 0; i < nBytes-1; i++)
          buffer[i] = toupper(buffer[i]);

    //Send uppercase message back to client, using serverStorage as the address
    sendto(udpSocket, buffer,nBytes, 0, (struct sockaddr*) &serverStorage, addr_size);

  }


  return 0;
}
