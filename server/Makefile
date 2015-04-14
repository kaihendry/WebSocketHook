all: server
	sudo ./server

server: main.go
	go build -a -tags netgo -installsuffix netgo .

upload: server
	scp index.html server sg.webconverger.com:

clean:
	rm -f server
