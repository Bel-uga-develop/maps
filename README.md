# Maps
Package for working with map slices in golang

## Features

- Search for the first element in a slice of maps by fields
- Filtering elements in a slice of maps
- Converting a map slice to a nested map slice with no nesting level limit

## Installation
```
go get github.com/Bel-uga-develop/maps
```
## Example
### Search for the first element in a slice of maps by fields

```	
data := []map[string]interface{}{
		{
			"id":                  1,
			"name":                "Reports",
			"parent_menu_item_id": nil,
			"route_name":          "reports",
		},
		{
			"id":                  2,
			"name":                "Orders",
			"parent_menu_item_id": nil,
			"route_name":          "orders",
		},
    }
    
foundData := maps.FindMap(
		data,
		func(item map[string]interface{}) bool {
			if item["id"].(int) == 2 {
				return true
			}
			return false
		},
	)
```

### Filtering elements in a slice of maps

 ```
 	fileteredMap := maps.FilterMap(
		data,
		func(item map[string]interface{}) bool {
			if item["id"].(int) < 5 {
				return true
			}
			return false
		},
	)
```
### Converting a map slice to a nested map slice with no nesting level limit
```
nestedMap := maps.GetNestedMap(data, "id", "parent_menu_item_id", "sub_item")
```