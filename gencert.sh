#/bin/sh

openssl req -x509 -out ./cert/localhost.crt -keyout ./cert/localhost.key  -newkey rsa:2048 -nodes -sha256 -subj '/CN=localhost'