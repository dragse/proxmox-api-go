# proxmox-api-go

The Proxmox API in golang. This Framework is still work in progress.

## Features

### Cluster
- [x] Version Endpoint
- [x] Get Next VM ID

### Node
- [x] List all Nodes
- [x] List all VMs of the Node
- [x] Get Node Status of a specific Node
- [x] Get Time Configuration of a specific Node
- [x] Update Time Zone of a specific node

### VM
- [x] Create VM
- [x] Create VNC TCP Proxy
- [x] Get VM Status
- [x] Update VM Status
- [x] Get Firewall Log
- [x] Get Firewall Options
- [x] Update Firewall Options
- [x] List IPSets
- [x] Create IPSets
- [x] Delete IPSets
- [x] Add CIDR to IPSet
- [x] Remove CIDR from IPSet
- [x] Copy a specific VM

### Pools
- [x] Create Pool
- [x] List All Pools
- [x] Show Pool Details

## Installation

``go get github.com/DragSE/proxmox-api-go``

## Usage

Create at first a client object. This objects manage and handle the http connections to the nodes

```go
proxClient := client.NewProxmoxClient()
```

After that you need to add all Nodes you want to connect. It is recommended to connect to multiple nodes if you have a 
cluster and want high availability  
The Cluster work with api tokens, you need to generate before.

```go
session := client.ProxmoxSession{
    Hostname:  "hostname of the proxmox node",
    Username:  "username@pve!token-name",
    Token:     "apitoken",
}

err := proxClient.AddSession(&session)

if err != nil {
    log.Fatal(err)
}
```

The last step is to create the Cluster and init the information

````go
proxCluster := proxmox.NewProxmoxCluster(proxClient)

err = proxCluster.InitInformation()

if err != nil {
    log.Fatal(err)
}

````

Now you can use the different methods to get or change Information of the cluster 

Support
-------

If you are having issues, please let us know.
We have a mailing list located at: proxmox-apit@dragse.de

License
-------

The project is licensed under the GPL-3.0 License.
