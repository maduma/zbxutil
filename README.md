Small utility to set Zabbix usegroup Read only permissions for all group


## Build
```
go build -o /opt/hrs/bin/zbxutil-fix-permissions
```

## Run
Pass token and zabbix url via environment variable
Pass usergroup ids list as parameter
```
ZABBIX_SERVER_URL=https://zabbix.exemple.com
ZABBIX_API_TOKEN=X...X
export ZABBIX_SERVER_URL ZABBIX_API_TOKEN
/opt/hrs/bin/zbxutil-fix-permissions 23,45
```
