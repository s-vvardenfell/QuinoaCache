# QuinoaCache <br>
Redis server with grpc API for Quinoa project<br>
Use `gen.sh` to generate grpc-code from proto/service.proto

Config example:<br>
```yaml
host: localhost
redis_port: "6379"
server_port: "50051"
pasword: ""
db_num: 0
with_reflection: true # allows to use grpcui
logrus:
  log_level: 4
  to_file: false
  to_json: false
  log_dir: "logs/logs.log"
```
