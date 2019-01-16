# go-chat

##### Version
v0:  
server.go  
client.go  

v1:
server1.go  
client1.go  

##### Demo v0
$ go run server.go  
New client: 127.0.0.1:56621  
New client: 127.0.0.1:56627  
[A]:Hello  
[A]:Good morning  
[B]:Good morning nice to see you  

$ go run client.go  
Input nickname:  
A  
Your nickname: A  
Hello  
Good morning  
[B]:Good morning nice to see you  

$ go run client.go  
Input nickname:  
B  
Your nickname: B  
[A]:Hello  
[A]:Good morning  
Good morning nice to see you  


