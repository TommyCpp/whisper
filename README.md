# Whisper

Whisper is a chat server based on `WebSocket`.

## Feature
* Peer chat as well as Group Chat.
* User authentication
* Multi-thread implemented by `goroutine`
* Encryption with RSA algorithm
* Support custom Middleware.
* Unit test included.

## Install with Docker
Run `docker-compose up` to start service
test data has been included in database.sql

## Config
You may change configuration as you want in `config/config.json`

## Architecture

## Todo
* Use React or other framework to rewrite frontend
* End to end encryption
* Other middleware like word filter etc.