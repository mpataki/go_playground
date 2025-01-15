# proto

This sub-project holds proto files for a grpc API.  It publishes packages for go, java, and npm which can be used to implement clients and servers.

Run the code gen:
```sh
buf generate
```
This will generate everything required for the go proto module.

To generate the npm module, run:
```
npm run build
```

To generate the java library, run:
```
./gradlew build
```

