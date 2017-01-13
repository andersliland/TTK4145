CC=gcc
CFLAGS= -c -Wall 
LDFLAGS= -lpthread
SOURCES= ex01.c
OBJECTS=$(SOURCES:.c=.o) 
EXECUTABLE= ex01

all: $(SOURCES) $(EXECUTABLE) 

$(EXECUTABLE): $(OBJECTS)     
	$(CC) $(LDFLAGS) $(OBJECTS) -o $@ 

.cpp.o:
	$(CC) $(CFLAGS) $< -o $@


..PHONY: run
run:
	./$(EXECUTABLE)

	