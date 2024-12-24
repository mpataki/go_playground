# proto

This sub-project holds proto files for a grpc API.  It publishes packages for go, java, and npm which can be used to implement clients and servers.

Run the code gen:
```sh
buf generate
```

To generate the npm module:
```
npm run build
```

Generate the java lib using gradle from the repo root. 
```
./gradlew build
```

