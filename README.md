# go-web-authentication
Basic Go Authentication

 * Sessions

 * Cookies

 * CSRF tokens

 * usernames/passwords
  
Registration: 

`$ curl -X POST -i -w '\n' "localhost:9080/register?username=y44k0v21&password=password123"`

Login: 

`$ curl -X POST -i -w '\n' "localhost:9080/login?username=y44k0v21&password=password123"`