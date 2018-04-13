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

## Available methods

All the specific documentation for the API is also available through `/swagger`
* `v1/monitor/status [GET]`: Gets the status of the monitor
* `v1/monitor/status [POST] {testString}`: Tests a POST request to the monitor to see if it works
* `v1/monitor/version [GET]`: Checks current version of the monitor and if it is outdated.
* `v1/monitor/update [GET]`: Checks current version of the monitor and updates it 
automatically if it is outdated. *return(CURRENT_VERSION/STATUS/GETEXCEPTIONS)*
* `v1/node/start/clean [POST]`: Makes a clean start of the node. *return(STATUS, GETEXCEPTIONS)*
* `v1/node/start/latest [GET]`: Checks the last time the node process was restarted
* `v1/node/log/json [GET]`: Gets a complete log with useful information from a validator node in
a JSON format
* `v1/node/log/istanbul [GET]`: Gets a complete log with useful information about the istanbul protocol (validators,)
* `v1/node/log/raw [GET]`: Gets a complete log in a raw format from the validator node
* `v1/node/propose [POST] {id, value}`: It proposes a node with an `id` and a `value`
for the propose
* `v1/node/restart [POST]`: Restarts the node
* `v1/node/network/restart [POST] {nodeType, nodeName}`: Used for network restarts. It updates
the node and reinitializes it for the new test-net conditions. *(EROR DE VALIDACIÓN DE NODOS 400)*
* `v1/node/start/default [POST]`: Starts the node.
* `v1/node/status [GET]` Gets the status of the node
* `v1/node/stop [POST]`: Stops the node
* `v1/node/update [GET]`: Updates every configuration file of the node from the alastria-node
repository if they are outdated
* `v1/node/version [GET]`: Checks the current version of the node
* `v1/node/coinbase [GET]`: Checks the nodes coinbase
* `v1/node/transactions [DELETE]`: Cleans the transaction queue of the node
* `v1/node/ismining [GET]`: Checks if the node is mining and, if this is not the case, it starts the miner



