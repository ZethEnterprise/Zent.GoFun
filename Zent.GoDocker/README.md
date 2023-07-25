Hello fellow Nerdies!

It's time for the next project (yeah! I know that I've committed it all at once, but that was before I actually throught of throwing it up into the cloud for ... well... Fun.

So I wanted to look into how to throw my Go-code into a Docker Container. So... As any (lazy) developer would do, you'll Google it and find a believable page to try out.
In my case I found this one: https://blog.devgenius.io/go-docker-hello-world-f092ecf7cead

It is quite interesting read - give it a go! (giggles)

But to summarize some of the most important commands I'll choose these:

-------------------------------------------------------
|                                                     |</br>
| $go mod init hellogohttp/m/v2                       |
| //It generates a go.mod file, which is a module file|
| //and it states which Go version it is running in   |
|                                                     |
| $go run main.go                                     |
| //Open new terminal since this one is hosting       |
| //A web page on port 8080                           |
|                                                     |
| $curl http://localhost:8080/helloworld              |
| Hello, World!                                       |
|                                                     |
| $docker build -t hello_go_http .                    |
| $docker run -p 8080:8080 -t hello_go_http           |
| //Open new terminal since this one is hosting       |
| //A web page on port 8080                           |
|                                                     |
| $curl http://localhost:8080/helloworld              |
| Hello, World!                                       |
|                                                     |
-------------------------------------------------------

So. Now we have made a textbased "Hello World" example and the same hosted in a Docker Image on a web page. How fun. Third project is going to be GUI! (YAY!)
