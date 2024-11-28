package main

import (
	"fmt"
	"log"
)

func main() {
	m1 := map[string]string{
		"a": "value1",
		"b": "value2",
		"c": "value1",
		"d": "value2",
	}

	sl3 := make([]map[string][]string, 0)

	valueToKeysMap := make(map[string][]string) // инициализируем пустую мапу
	for k, v := range m1 {                      // "a": "value1",
		if _, ok := valueToKeysMap[v]; !ok { //  <<< Эта часть кода проверяет, есть ли уже значение v в карте valueToKeysMap.
			// Если значение v там ещё не было, создается новый пустой срез строк и сохраняется в valueToKeysMap под ключом v.
			valueToKeysMap[v] = make([]string, 0) // <<< создаем пустой срез в мапе valueToKeysMap с ключем "value1"
		}
		valueToKeysMap[v] = append(valueToKeysMap[v], k)
		log.Println(valueToKeysMap)
		// ^^^ На этом этапе ключ k добавляется в конец соответствующего среза в карте valueToKeysMap.
		// То есть каждый раз, когда встречается новое значение v, оно добавляется в
		//соответствующий срез, который хранится под этим значением в valueToKeysMap.
	}

	for v, keys := range valueToKeysMap {
		m2 := make(map[string][]string)
		m2[v] = keys
		sl3 = append(sl3, m2)
	}

	fmt.Println("Result:", sl3)
	/*
		2024/11/28 18:40:55 map[value1:[a]]                    <<< первая итерация
		2024/11/28 18:40:55 map[value1:[a] value2:[b]]         <<< вторая итерация
		2024/11/28 18:40:55 map[value1:[a c] value2:[b]]       <<< ...
		2024/11/28 18:40:55 map[value1:[a c] value2:[b d]]
		Result: [map[value1:[a c]] map[value2:[b d]]]

	*/
}
