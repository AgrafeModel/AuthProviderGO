# AuthProviderGO - A fast and easy way to deploy an authentication provider in Go
<span style="color:red"><b>WARNING:</b></span> This project is not yet ready for production use. It is still in the early stages of development and may contain bugs or security vulnerabilities. Use at your own risk!

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Dependencies](#dependencies)
- [Contributing](#contributing)


## Introduction
### What is AuthProviderGO?
AuthProviderGO is a fast and easy way to deploy an authentication provider in Go. It is a simple and lightweight web server that provides a RESTful API for user authentication. It is designed to be easy to use and easy to deploy, with minimal configuration required.

### Why AuthProviderGO?
There are many authentication providers available, but most of them are complex and difficult to deploy. AuthProviderGO is designed to be simple and lightweight, with minimal configuration required. It is easy to deploy and easy to use, making it ideal for small projects and prototyping.
And to be more honest, I needed a simple authentication provider for a project I was working on, so I decided to build one myself. I hope you find it useful too!

## Features
- Simple and lightweight
- Easy to deploy
- RESTful API
- User registration
- JWT-based authentication
- Password hashing
- Configurable

And more to come!

## Installation

To install AuthProviderGO, you can clone the repository and build the binary yourself, or you can download the pre-built binary from the releases page. (coming soon)


### Building from source
To build AuthProviderGO from source, you will need to have Go installed on your system. You can download and install Go from the official website: [https://golang.org/](https://golang.org/)

Once you have Go installed, you can clone the repository and build the binary using the following commands:

```bash
git clone
cd AuthProviderGO
go build
```

## Usage

To edit the configuration, you can create a `.env` that contains sensitive information like the JWT secret and the database connection string. You can also use the `config.yaml` file to configure the server to fit your needs.

## Dependencies
This project uses the following dependencies:
- [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- [github.com/a-h/templ](https://github.com/a-h/templ)
- [github.com/joho/godotenv](https://github.com/joho/godotenv)
- [github.com/spf13/viper](https://github.com/spf13/viper)

## Contributing
If you would like to contribute to AuthProviderGO, please feel free to submit a pull request or open an issue. I am always looking for ways to improve the project and make it more useful for others.
