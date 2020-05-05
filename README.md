# Set of net package usage examples
This repo illustrates basic functionality we have discussed in lesson 6  
## grabber
Shows TCP connection to REQUEST HTTP page and dump RESPONSE to a local file  
## slack
http GET request example. Requires `token` and `channel` environment variables.   
_note: If you don't want to create slack app, I share creds from my test bot in channel_
## whois
TCP client to get `whois` data from server.
## echo
TCP server that simply returns a message sent.  
_note: use `telnet 127.0.0.1 8081` to connect_