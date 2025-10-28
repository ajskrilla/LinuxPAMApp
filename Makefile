# Makefile
CC       := gcc
CFLAGS   := -fPIC -Wall -I/usr/include/security
LDFLAGS  := -shared -lpam

SRC      := src/pam_local.c
OBJ      := pam_local.o
MODULE   := pam_local.so

all: $(MODULE)

# 1) Compile .c → .o
$(OBJ): $(SRC)
	$(CC) $(CFLAGS) -c -o $@ $<

# 2) Link .o → .so
$(MODULE): $(OBJ)
	$(CC) $(LDFLAGS) -o $@ $<

clean:
	rm -f $(OBJ) $(MODULE)
