# go-web-authentication
Basic Go Authentication

 * Sessions

 * Cookies

 * CSRF tokens

 * usernames/passwords
  
**Registration**: 

`$ curl -X POST -i -w '\n' "localhost:9080/register?username=y44k0v21&password=password123"`

**Login**: 

`$ curl -X POST -i -w '\n' "localhost:9080/login?username=y44k0v21&password=password123"`

**Protected**:
`$ curl -X POST -i -w '\n'  "localhost:9080/protected?username=y44k0v21&password=password123" \ -H "X-CSRF-Token:63x6go3IcOxPDgtXvsDpaiNuf-b1RaP0WTp9YjbPeeg=" \
--cookie "session_token=LjQy0UZbeQHlyg0zptqQwgZ1Q2mWEfxYQKJt0njrWnQ="`

**Logout**:
`$ curl -X POST -i -w '\n'  "localhost:9080/logout?username=yanky33&password=password123" -H "X-CSRF-Token:63x6go3IcOxPDgtXvsDpaiNuf-b1RaP0WTp9YjbPeeg=" \ 
--cookie "session_token=LjQy0UZbeQHlyg0zptqQwgZ1Q2mWEfxYQKJt0njrWnQ="`