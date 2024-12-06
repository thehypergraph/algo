CC = gcc
CFLAGS = -Wall -Wextra
SRC_DIR = .
SERIES_DIR = series
SRCS = $(SRC_DIR)/main.c $(SERIES_DIR)/fib.c
OBJS = $(SRCS:.c=.o)
TARGET = algo

all: $(TARGET)

$(TARGET): $(OBJS)
	$(CC) $(OBJS) -o $(TARGET)

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

clean:
	rm -f $(OBJS) $(TARGET)

run: $(TARGET)
	./$(TARGET)

.PHONY: all clean run
