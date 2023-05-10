[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)
# Todo app API with JWT Authentication and MongoDB

This is a simple yet powerful app written in GoLang to showcase the uses of JWT Authentication and MongoDB There is a login/signup feature and a todo list feature. Every user and todo is stored in a MongoDB collection. Every session is secured and managed by JWT Authentication. The passwords are hashed and stored in MongoDB in a industry standard way.

# What is JWT Authentication?
![Rest API and MDB](https://jwt.io/img/logo-asset.svg)

JSON Web Token (JWT) is an open standard (RFC 7519) that defines a compact and self-contained way for securely transmitting information between parties as a JSON object. This information can be verified and trusted because it is digitally signed. JWTs can be signed using a secret (with the HMAC algorithm) or a public/private key pair using RSA or ECDSA.

Although JWTs can be encrypted to also provide secrecy between parties, we will focus on signed tokens. Signed tokens can verify the integrity of the claims contained within it, while encrypted tokens hide those claims from other parties. When tokens are signed using public/private key pairs, the signature also certifies that only the party holding the private key is the one that signed it.

# When is it used?

- **Authorization**: This is the most common scenario for using JWT. Once the user is logged in, each subsequent request will include the JWT, allowing the user to access routes, services, and resources that are permitted with that token. Single Sign On is a feature that widely uses JWT nowadays, because of its small overhead and its ability to be easily used across different domains.

- **Information Exchange**: JSON Web Tokens are a good way of securely transmitting information between parties. Because JWTs can be signed—for example, using public/private key pairs—you can be sure the senders are who they say they are. Additionally, as the signature is calculated using the header and the payload, you can also verify that the content hasn't been tampered with.

# Project Overview

Each user has to sign up and create an account of their own or log into an existing one before they can access the app.
Each user's passwords are hashed and stored in MongoDB in a industry standard way.
![Rest API and MDB](https://i.imgur.com/rpOaQCw.png)

There several operations which you can perform in the app:
- Create a task
- Delete a task
- Update a task
- Finish a task(check)
- Delete all tasks

# Installation

Clone the repository into a directory of your choice Run the command `go mod tidy` to download the necessary packages.
You'll need to add a .env file and add a MongoDB connection string with the name `MONGODB_URL` to access your collection for task and user storage.
You'll also need to add `SECRET_KEY` to the .env file for JWT Authentication.

Run the command `go run main.go` and the project should run on `locahost:8080`

# License

This project is licensed under the terms of the MIT license.
