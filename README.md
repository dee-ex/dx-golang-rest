# Golang RESTful API Setup
[![goversion](https://img.shields.io/badge/Go-v1.14.4-blue)](https://golang.org/)
[![pythonversion](https://img.shields.io/badge/Python-v3.9-blue)](https://python.org/)
[![mysqlversion](https://img.shields.io/badge/MySQL-v8.0.22-blue)](https://mysql.com/)
## Table of contents
* [Installation](#installation)
    * [What Is My Setup?](#what-is-my-setup)
    * [Why Should You Use My Setup?](#why-should-you-use-my-setup)
    * [Getting Started](#getting-started)
* [Architecture](#architecture)
    * [`database` Directory](#database-directory)
    * [`entities` Directory](#entities-directory)
    * [`infrastructures` Directory](#infrastructures-directory)
    * [`middlewares` Directory](#middlewares-directory)
    * [`modules` Directory](#modules-directory)
    * [`utils` Directory](#utils-directory)
* [Starting Server](#starting-server)
* [Contributing](#contributing)
* [License](#license)
## Installation
### What Is My Setup?
My Setup is a simple setup for you to create you RESTful API in Golang.
### Why Should You Use My Setup?
Have you found yourself crazy in building your setup to create your RESTful API time after time?  
If you have, the solution of mine is the thing I see that you really need.  
It doesn't have too much abstraction, which lets you feel free to develop your own API.
### Getting Started
No matter you're developing on macOS, Windows or Linux, you can just clone/download this project and use a simple terminal command to create your project base on my Setup.
```
python setup.py -u username -p project_name
```
## Architecture
We will see the architecture in detail by having awareness of function of directories exists in my Setup.
### `database` Directory
All files needed for your migrations is located in `migrations` folder.
### `entities` Directory
This directory takes care your models that you need for mapping data to work with your database.
### `infrastructures` Directory
`sessions.go` in `types` directory is a struct, which is constructed to store session information to connect to your database.  
`mysql` folder consits of two files. `session_construction.go` constructs a session and the other will create a new MySQL from session information.  
Later, you only need to care of construction file to construct your own session.
### `middlewares` Directory
This folder provides midlewares, inspect and filter HTTP requests entering your application.
### `modules` Directory
You build all your modules of your application in here. You can see that I built a example `users` modules based on repository pattern - the architecture of my Setup.  
In case, you may also need `tokens` to create, validate and extract token for your authorization.
### `utils` Directory
Cosisting of global functions for modules.  
I built a some basic http responses like Text, JSON response to fetch the result in `responses\http_responses.go`.
## Starting Server
I did setup some routes in `main.go`. You can run the server in port `6969` by command.
```
go run .
```
If you want to change running port, you can edit in this line in `main.go`
```golang
log.Fatal(http.ListenAndServe(":6969", enhanced_r))
```
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.  
Please make sure to update tests as appropriate.
## License
[MIT](https://choosealicense.com/licenses/mit/)
