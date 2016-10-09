All: rpc_client rpc_server

GCFLAGS="-N -l"
rpc_client:rpc_client.go
	go fmt $<
	go build -gcflags $(GCFLAGS) $< 

rpc_server:rpc_server.go
	go fmt $<
	go build -gcflags $(GCFLAGS) $<

.PHONY:clean

clean:
	-rm rpc_client rpc_server


