# Hangman
A multiplayer hangman with fun ascii stick figures.

## Purpose
Initially for fun and to learn Go, I eventually tweaked it 
to add multithreaded/multiplayer functionality.  

## Requirements
Go (I used Go 1.11, maybe 1.9 will work, they're pretty good on 
backwards compatability between versions).

## Running it
* Spin up a server in one terminal
* In another terminal, run the client. If you're connected to the same
 network, everything should work.

Note: If you're running both on only one machine, you might want to add
127.0.0.1 in front of the port number on both the client/server programs before
compiling. It's been a while since I did anything with this, 
but I remember that posing some kind of issue.