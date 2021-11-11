# proxmox-api-go

The Proxmox API in golang. This Framework is still work in progress.

## Installation

``go get github.com/DragSE/proxmox-api-go``

## Usage

Create at first a cluster object from where you can work with

```
proxCluster := proxmox.NewProxmoxCluster()
```

After that you need to add all Nodes you want to connect. It is recommended to connect to multiple nodes if you have a 
cluster and want high availability, if one node is offline or something like it.  
The Cluster work with api tokens, you need to generate

```
session := client.ProxmoxSession{
    Hostname:  "hostname of the proxmox node",
    Username:  "username@pve!token-name",
    Token:     "apitoken",
}
```

The last step is to init the cluster information

````
err = proxCluster.InitInformation()

if err != nil {
    log.Fatal(err)
}
````

Now you can use the different methods to get or change Information of the cluster 
