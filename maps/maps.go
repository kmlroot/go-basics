package maps

// GetMap returns map with key string and integer values
func GetMap(name string) int {
	// mapTest := make(map[string]int)

	mapTest := map[string]int{
		"Mauricio": 22,
		"Ana":      21,
		"Alba":     52,
	}

	mapTest["key"] = 1
	mapTest["key2"] = 200

	delete(mapTest, "key")
	delete(mapTest, "key2")

	return mapTest[name]
}
