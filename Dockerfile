FROM iron/go

# CMD mkdir /app

WORKDIR /app

# Now just add the binary
# ADD squiddleb.it_linux_386 /app/

ENTRYPOINT ["./squiddleb.it_linux_386"]