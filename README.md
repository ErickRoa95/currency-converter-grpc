# Currency Converter GRPC

CurrencyConverter GRPC is a project to showcase the interaction between a GRPC client and a GRPC Server using GO. 

Server will provide RPC endpoints that would return CurrencyRates information to the GRPC client. 

## Tools
Tools that are used on this project:
- Go 1.23.10 
- Protobuff with [grpc-go](https://grpc.io/docs/languages/go/quickstart/) (libprotoc 28.3)
- [mockgen](https://github.com/uber-go/mock) v0.5.2
- Use [openexchangerates](https://openexchangerates.org/) API for data_dump. 


## Run the project 

Before running the project, make sure that **protoc** command is correctly installed on the system. Afterwards, we need to compile proto files to correctly create go dependencies used on th GRPC server & client. 

For this, we will need to run the following command:
```
make protoc 
```
<br>

A **Makefile** was created to easily recompiled proto files. If you need to update the output files of the protoc compiling step, **please update the file Makefile**. 


> **NOTE:**  Everytime the **currency.proto** file is updated, the command `make protoc` needs to be excuted again. 

<br>

## Run GRPC SERVER
After correctly compiling the proto files, we can now run the GRPC server using the following command:
```
make server  
```

This will start a GRPC server with default values. In case you need to use another values for the initalization, we support the following flags for the server:

<br>

| Flag | description | default value |
| ---- | ----------- | ------------- |
| port | Server port |  50051        |


To be able to run the server with a custom value:
```
go run server/servergo -port=<port_server>
```
<br>

> **NOTE:**  In the case that you update currency.proto file, you will need to restart the server. 


## Run GRPC Client 
For the GRPC client, we support a couple of flags that will be needed to make RPC calls to the server. 

| Flag | description | default value |
| ---- | ----------- | ------------- |
| port | Server port |  50051        |
| code | CountryCode | MXN |
| rpc | Call Specific RPC endpoint | Converter |
| amount | Amount to convert | 10.00 |

The GRPC Server currently has 2 RPC endpoint:
- **Converter** : Which provide the CurrencyRates for an specific Country Code. 
- **Exchange**: Will provide a the Exchange calcualtion of the specific amount that we send to the Server. 


## Example - GRPC Client for Converter
To execute the GRPC Client for the CONVERTER call we use the following command:
```
go run client/client.go -code=MXN -rpc=Converter
```

This will return the following output:
```
== Call Response: CurrencyName:"Mexican Peso"  CountryCode:"MXN"  Base:"USD"  Amount:19.00649
```

## Example - GRPC Client for Exchange Examp
To execute the GRPC Client for the Exchange call we use the following command:
```
go run client/client.go -code=MXN -amount=1000.0 -rpc=Exchange
```
 
This will return the following output:
```
 == Call Response: CountryCode:"MXN"  Base:"USD"  Amount:1000  Exchange:52.613605
```