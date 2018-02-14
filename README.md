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
* Check current version of the monitor
```
./monitor.sh version
```
* Check latest version available for the monitor in the repository
```
./monitor.sh latest
```

## Available commands
All the specific documentation for the API is also available through 
* `/monitor/status [GET]`
* `/monitor/status [POST]`
* `/monitor/version [GET]`: Checks current version of the monitor and if it is outdated.
* `/monitor/versionUpdate [GET]`: Checks current version of the monitor and updates it 
automatically if it is outdated
* `/node/cleanstart [GET]`: Makes a clean start of the node
* `/node/lastrestart [GET]`: Checks the last time the node process was restarted
* `/node/logsjson [GET]`: Gets a complete log with useful information from a validator node in
a JSON format
* `/node/logsraw [GET]`: Gets a complete log in a raw format from the validator node
* `/node/propose [POST] {id, value}`: It proposes a node with an `id` and a `value`
for the propose
* `/node/restart [GET]`: Restarts the node
* `/node/restartNetwork [POST] {nodeType, nodeName}`: Used for network restarts. It updates
the node and reinitializes it for the new test-net conditions.
* `/node/start [GET]`: Starts the node.
* `/node/status [GET]` Gets the status of the node
* `/node/stop [GET]`: Stops the node
* `/node/update [GET]`: Updates every configuration file of the node from the alastria-node
repository if they are outdated
* `/node/version [GET]`: Checks the current version of the node