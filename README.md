# Lösenord.xyz

Skapad av [Markus Tenghamn](http://tenghamn.com)

This is my first project with golang, I personally feel that the code is a bit messy as
I am learning but things seem to work well at the moment. I would be open to feedback and
criticism and if anyone is interested feel free to contact me if you are interested in developing
the project further.

The project uses the Revel framework.

See the project live at http://lösenord.xyz

Credit to this blog post for helping me get Revel running in Docker http://jbeckwith.com/2015/05/08/docker-revel-appengine/

Credit for the nice password generating function goes to https://github.com/cmiceli/password-generator-go

## How to run this locally

Install docker

Clone this project

Then run the following commands in the project directory

`docker build -t password-fun .`

`docker run -it -p 8080:8080 password-fun`

The app should now be running on http://127.0.0.1:8080