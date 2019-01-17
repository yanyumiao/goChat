# go-chat

##### Version:
[v0](https://github.com/yanyumiao/gochat/tree/master/v0)  
[v1](https://github.com/yanyumiao/gochat/tree/master/v1)  use channel to realize broadcast, better than v0    

##### Demo:
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


