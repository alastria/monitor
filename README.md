# Alastria Monitor
This piece os software publishes a set of commands through a REST API for management purposes. The endpoint is exposed in **port 8443** for HTTPS and **port 8080** for HTTP.

**NOTE**: For its correct operation using HTTP and HTTPS, node machines should have ports 8443/8080 opened. The current version of the monitor is only HTTP compatible. In future releases, only HTTPS will be enabled.

## Install and run the monitor
All the installation and runnning operations for the monitor are managed through the script `scripts/monitor.sh` from [alastria-node](https://github.com/alastria/alastria-node)
* Install the monitor (if we ran the latest version of `scripts/bootstrap.sh` this should work without further configurations)
```
./monitor.sh build
```
* Run the monitor
```
./monitor.sh start
```
