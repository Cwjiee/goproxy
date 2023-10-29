# goproxy
A proxy server built with golang

# features
- sends request from client to server and forward response back to client
- logging
  - log details of incoming requests for easier debugging
- copies the header and body of request and responses
- manipulate headers
  - manipulate user agent in request to enhance user privacy and anonymity
- content filtering
  - censors positive words and replaces it to "****"

# testing
- used a [to-do list api](https://github.com/Cwjiee/todo-list-api) to serve as backend server
- used Insomnia for client side testing
