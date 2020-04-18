# Accounting System - Paxos Challenge
## Introduction

This project consist of an API and a React Frontend aiming to fullfil
this 
[specification](https://agileengine.bitbucket.io/fsNDJmGOAwqCpzZx/)
)

For Development of the frontend just head to client and you can use the available scripts

To run the application just use 

#### `cd client && yarn run build`

after that

#### `./go-executable-build.sh name_you_want`

Choose your executable based on your system from build folder

Run it

Open [http://localhost:3000](http://localhost:3000)

## API

A simple API providing base funtionality to 
operate with a user account adding or substracting money from it

See [swagger](https://agileengine.bitbucket.io/fsNDJmGOAwqCpzZx/api/#/transactions/getTransactionById
)

## Client
A simpistic frontend made with create-react-app aiming to provide
a list of the account transactions.

## Available Scripts in Client

Inside of the client, you can run:

### `yarn start`

Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

### `yarn test`

Launches the test runner in the interactive watch mode.<br />
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `yarn build`

Builds the app for production to the `build` folder.<br />

## Available Scripts in Main directory

### `go run main.go`
Runs the application, serving the static files inside de build folder of the client

### `./go-executable-build.sh name_you_want`

Builds the application for the major platforms to be distributed
The output will be in the build folder inside the client

## Technologies

* Golang
* React
* Gin Gonic
* Material UI


