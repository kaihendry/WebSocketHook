# Web socket to Web hook

Some devices are behind a NAT/Firewall so they can't get Web hooks posted to
them.

So a workaround is to use Web sockets, to connect to a service and listens &
waits for a Webhook event. When the event is fired, the client performs a Web
request itself, after being sent the Web hook URL over the Web socket.

## Walkthough

The client could and probably should be a Web application, though for testing
[client.go] can be used.

The client connects to [server.go] and listens for a URL fired by an event at
the server. Once it has the URL, it returns the URL & exits. Otherwise the
client tries to maintain a socket connection until a URL is received.

## Other implementations

<http://web.sockethook.io/> but for some reason the socket closes... [it should stay open](https://github.com/factor-io/websockethook/issues/5)
