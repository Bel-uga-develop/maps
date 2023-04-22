package maps

func FilterMap(data []map[string]interface{}, callback func(map[string]interface{}) bool) []map[string]interface{} {
	filteredData := []map[string]interface{}{}

	for _, item := range data {
		if callback(item) {
			filteredData = append(filteredData, item)
		}
	}

	return filteredData
}

func FindMap(data []map[string]interface{}, callback func(map[string]interface{}) bool) map[string]interface{} {
	foundData := map[string]interface{}{}

	for _, item := range data {
		if callback(item) {
			foundData = item
		}
	}

	return foundData
}

func GetNestedMap(data []map[string]interface{}, columnID string, columnParent string, columnChild string) []map[string]interface{} {
	nestedMap := []map[string]interface{}{}

	for _, item := range data {
		if item[columnParent] == nil {
			newItem := item
			id := item[columnID].(int)
			newItem[columnChild] = getSubRecursive(data, id, columnID, columnParent, columnChild)
			nestedMap = append(nestedMap, newItem)
		}
	}

	return nestedMap
}

func getSubRecursive(data []map[string]interface{}, parentID int, columnID string, columnParent string, columnChild string) []map[string]interface{} {
	itemsMap := []map[string]interface{}{}

	subItemsMap := FilterMap(data, func(item map[string]interface{}) bool {
		if item[columnParent] != nil && item[columnParent].(int) == parentID {
			return true
		}
		return false
	})

	if len(subItemsMap) != 0 {
		for _, value := range subItemsMap {
			item := value
			id := value[columnID].(int)
			item[columnChild] = getSubRecursive(data, id, columnID, columnParent, columnChild)
			itemsMap = append(itemsMap, item)
		}
	}

	return itemsMap
}
