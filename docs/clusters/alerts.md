# Alerts
Cluster alerts are used for something. Once I figure out their purpose I will write it here.

## Interface

### EnumerateAlerts
Enumerate enumerates alerts on this cluster for the given resource within a specific time range.

#### Example

{% codetabs name="Golang", type="go" -%}
resp, err := m.EnumerateAlerts(
                timeStart,
                timeEnd, 
                api.ResourceType_RESOURCE_TYPE_VOLUME)
if err != nil {
    return nil
}
for _, alert := range resp.Alert {
    fmt.Printf("ID:=%d Severity=%v\n",
        alert.Id,
        alert.Severity)
}
{%- language name="Python", type="py" -%}
TBD
{%- endcodetabs %}

### ClearAlert
ClearAlert clears an alert for the given resource.

### EraseAlert
EraseAlert erases an alert for the given resource.

## Reference
* [Golang](https://godoc.org/github.com/libopenstorage/openstorage/cluster#ClusterAlerts)