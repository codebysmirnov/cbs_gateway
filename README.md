# CBS_GATEWAY

___

1. Config build on many simple parts.
    1. JSON - format.
    2. including.
    3. env variables.
2. Proxy.
3. Swagger generator.
4. Security.
5. Mock 
   1. response data
   2. response http code

---

## Run proxy

1. create global.json config file from example.global.json template on project root.
      1. configure entrypoint: write host and port for proxy listening.
      2. logging don't need setting, there is has just one mode 'stdout'.
2. build and run go binary file.