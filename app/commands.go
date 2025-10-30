package main

type RedisCommands string

const (
	ping RedisCommands = "ping"
	echo RedisCommands = "echo"
	get  RedisCommands = "get"
	set  RedisCommands = "set"
)
