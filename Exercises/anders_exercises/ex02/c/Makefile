CC=gcc
CFLAGS= -c -Wall 
LDFLAGS= -lpthread
SOURCES= ex02.c
OBJECTS=$(SOURCES:.c=.o) 
EXECUTABLE= ex02

all: $(SOURCES) $(EXECUTABLE) 

$(EXECUTABLE): $(OBJECTS)     
	$(CC) $(LDFLAGS) $(OBJECTS) -o $@ 

.cpp.o:
	$(CC) $(CFLAGS) $< -o $@


..PHONY: all
run:
	./$(EXECUTABLE)

	