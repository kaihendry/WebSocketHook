# Web socket to Web hook

Some devices are behind a NAT/Firewall so they can't get Web hooks posted to
them directly.

So a workaround is to use Web sockets, whereby clients connect to a service and
listens & waits for a Webhook event. When the event is fired, the client
performs a Web request itself, after being sent the Web hook URL over the Web
socket.

## Other implementations

* <http://requestcatcher.com/>
* <http://web.sockethook.io/> but for some reason the socket closes... [it should stay open](https://github.com/factor-io/websockethook/issues/5)

For more fault tolerance, i.e. when server goes down: <https://github.com/joewalnes/reconnecting-websocket>
