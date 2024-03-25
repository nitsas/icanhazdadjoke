# icanhazdadjoke

A simple Go program that consumes the icanhazdadjoke API: https://icanhazdadjoke.com/api

## How to build

```sh
git clone https://github.com/nitsas/icanhazdadjoke
cd icanhazdadjoke
go build
```

## Usage

Assuming that you have built the program and named the result `icanhazdadjoke` (the default):
```sh
# Print a random joke:
$ icanhazdadjoke
Joke cxHYg3gFQf:
I got fired from the transmission factor, turns out I didn't put on enough shifts...

# Print a specific joke by id:
$ icanhazdadjoke EYoz51DtHtc
Joke EYoz51DtHtc:
What do computers and air conditioners have in common? They both become useless when you open windows.
```
