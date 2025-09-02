# Logger

This package is a replacement for `common` package logging functions.
All wdc projects should eventually switch to this.

In order to read docker logs easier bind log directory using docker-compose
Use Cfg.ContainerIdVarName to create unique dir for each container. To avoid any file name conflicts.

## Log reader usage
```bash
go run cmd/log_reader/main.go --level 51 --file ../../email_sender/log/m-lptp/26_Sep_24_12_59_11.jsonl
go run cmd/log_reader/main.go --level 51 --file ../../email_sender/log/m-lptp/26_Sep_24_12_59_11.jsonl --start '2024/09/26 12:59:12-05:00'
go run cmd/log_reader/main.go --level 51 --file ../../email_sender/log/m-lptp/26_Sep_24_12_59_11.jsonl --end '2024/09/26 12:59:12-05:00'
go run cmd/log_reader/main.go --level 51 --file ../../email_sender/log/m-lptp/26_Sep_24_12_59_11.jsonl --start '2024/09/26 12:59:12-05:00' --end '2024/09/26 12:59:14-05:00'
```
