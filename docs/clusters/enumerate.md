# Enumerate
Enumerate is used to list all the nodes in the cluster.

## Example

{% codetabs name="Golang", type="go" -%}
clusterInfo, err := m.Enumerate()
if err != nil {
    return nil
}
fmt.Printf("Cluster ID:%s\n", clusterInfo.Id)
for _, node := range clusterInfo.Nodes {
    fmt.Printf("ID:=%d MgmtIP=%s DataIP=%s\n",
        node.Id,
        node.MgmtIp,
        node.DataIp)
}
{%- language name="Python", type="py" -%}
TBD
{%- endcodetabs %}

## Reference