# Dockerizing from the outside

This is a bit harder than with the monolith, but I'm sure we can do it.

Can you create three new Dockerfile files:

- web.Dockerfile
- router.Dockerfile
- repository.Dockerfile

which will Dockerize each one of the microservices?

You will need to use the `-f` option on `docker build`.
