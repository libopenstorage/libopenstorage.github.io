# Remove
Remove a node or multiple nodes from the cluster permanently

> NOTE: TODO This feels like it should only take a list of IDs. The users may not know how much to fill the struct to give to `Remove()`.

## Example

{% codetabs name="Golang", type="go" -%}
nodes := []api.Node{
    {
        Id: "id",
    },
}
forceRemove := false
err := m.Remove(nodes, forceRemove)
if err != nil {
    return nil
}
{%- language name="Python", type="py" -%}
TBD
{%- endcodetabs %}

## Reference