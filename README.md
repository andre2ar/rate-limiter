# Rate Limiter

Rate limiter middleware created to run over fiber web server.

## Running the server

The live reloading server can be run by the following commands:

``make up``

or

``docker-compose up``

There is only one end-point:

``localhost:8080/api/v1/session``

This end-point will only create a JWT and return it.

Examples request can be found on api folder.

## Changing cache

By default, Redis is used to handle caching. Inside the folder pkg/cache a cache interface can be found, by implementing
this interface Redis can be changed by any other cache system. 

Cache client is created on main.go file, there you can create and pass your new implementation if needed.