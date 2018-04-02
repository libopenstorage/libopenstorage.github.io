# Inspect
Get information on a node

## Example

{% codetabs name="Golang", type="go" -%}
node, err := m.Inspect(nodeId)
if err != nil {
    return nil
}
fmt.Printf("ID:=%d MgmtIP=%s DataIP=%s\n",
    node.Id,
    node.MgmtIp,
    node.DataIp)
{%- language name="Python", type="py" -%}
TBD
{%- endcodetabs %}

## Reference