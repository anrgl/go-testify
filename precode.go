package main

import (
	"net/http"
	"strconv"
	"strings"
)

var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

const CountMissing = "count missing"
const WrongCountValue = "wrong count value"
const WrongCityValue = "wrong city value"

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(CountMissing))
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil || count < 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(WrongCountValue))
		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(WrongCityValue))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}
