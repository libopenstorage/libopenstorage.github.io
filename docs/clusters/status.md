# Status
Get status from a specific node or its peers

## Interface

### Status Types
TODO Add link here

### Node Status
Returns the status of the node running the REST service as seen by the Cluster Provider for a given listener. If listenerName is empty it returns the status of THIS node maintained by the Cluster Provider. At any time the status of the Cluster Provider takes precedence over the status of listener. Precedence is determined by the severity of the status.

{% codetabs name="Golang", type="go" -%}
nodeStatus, err := m.NodeStatus()
if err != nil {
    return nil
}
fmt.Printf("Status: %s\n", nodeStatus)
{%- language name="Python", type="py" -%}
TBD
{%- endcodetabs %}

### Peer Status
Returns the statuses of all peer nodes as seen by the Cluster Provider for a given listener. If listenerName is empty is returns the statuses of all peer nodes as maintained by the ClusterProvider (gossip)

{% codetabs name="Golang", type="go" -%}
// TODO: Need example on creating a listener then using it for PeerStatus()
{%- language name="Python", type="py" -%}
TBD
{%- endcodetabs %}

## Reference