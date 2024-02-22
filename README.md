# file-stream-tcp-go

How to stream large files over TCP in Go instead of reading the whole file into memory.

## version 1 (no stream) 
```
go run . nostream
```
if suppose the sent bytes in > 2048 (eg 4000)
then it will be sent in chunks
> received 2048 bytes over the network
> received 1952 bytes over the network

## version 2 (stream) 
```
go run . stream
```
if suppose the sent bytes in > 2048 (eg 16000)
then it will be sent in stream (together in one go)
> received 16000 bytes over the network