package main

/*
=== HTTP server ===
Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.
В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var events = sync.Map{}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		throwError(w, ErrorWrongMethod)
		return
	}

	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		throwError(w, ErrorCanNotParseDate)
		return
	}

	id := r.FormValue("id")
	title := r.FormValue("title")

	event := Event{
		ID:    id,
		Title: title,
		Date:  date,
	}

	events.Store(id, event)

	v, err := json.Marshal(APIResult{APIEventID{ID: id}})
	if err != nil {
		throwError(w, ErrorInternalError)
		return
	}
	fmt.Fprint(w, string(v))
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		throwError(w, ErrorWrongMethod)
		return
	}

	date, err := time.Parse("2006-01-02", r.FormValue("date"))
	if err != nil {
		throwError(w, ErrorCanNotParseDate)
		return
	}

	id := r.FormValue("id")
	title := r.FormValue("title")

	event := Event{
		ID:    id,
		Title: title,
		Date:  date,
	}

	events.Store(id, event)

	v, err := json.Marshal(APIResult{APIEventID{ID: id}})
	if err != nil {
		throwError(w, ErrorInternalError)
		return
	}
	fmt.Fprint(w, string(v))
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		throwError(w, ErrorWrongMethod)
		return
	}

	id := r.FormValue("id")

	events.Delete(id)

	v, err := json.Marshal(APIResult{APIEventID{ID: id}})
	if err != nil {
		throwError(w, ErrorInternalError)
		return
	}
	fmt.Fprint(w, string(v))
}

func GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		throwError(w, ErrorWrongMethod)
		return
	}

	result := make(map[string][]APIEvent)

	events.Range(func(k interface{}, v interface{}) bool {
		e, ok := v.(Event)
		if !ok {
			return true
		}

		date := e.Date.Format("2006-01-02")

		if _, ok := result[date]; !ok {
			result[date] = make([]APIEvent, 0)
		}

		result[date] = append(result[date], APIEvent{
			ID:    e.ID,
			Title: e.Title,
			Date:  e.Date.Format("2006-01-02"),
		})

		return true
	})

	v, err := json.Marshal(APIResult{result})
	if err != nil {
		throwError(w, ErrorInternalError)
		return
	}
	fmt.Fprint(w, string(v))
}

func GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		throwError(w, ErrorWrongMethod)
		return
	}

	result := make(map[string][]APIEvent)

	events.Range(func(k interface{}, v interface{}) bool {
		e, ok := v.(Event)
		if !ok {
			return true
		}

		date := e.Date.AddDate(0, 0, -int(e.Date.Weekday())).Format("2006-01-02")

		if _, ok := result[date]; !ok {
			result[date] = make([]APIEvent, 0)
		}

		result[date] = append(result[date], APIEvent{
			ID:    e.ID,
			Title: e.Title,
			Date:  e.Date.Format("2006-01-02"),
		})

		return true
	})

	v, err := json.Marshal(APIResult{result})
	if err != nil {
		throwError(w, ErrorInternalError)
		return
	}
	fmt.Fprint(w, string(v))
}

func GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		throwError(w, ErrorWrongMethod)
		return
	}

	result := make(map[string][]APIEvent)

	events.Range(func(k interface{}, v interface{}) bool {
		e, ok := v.(Event)
		if !ok {
			return true
		}

		date := e.Date.Format("2006-01")

		if _, ok := result[date]; !ok {
			result[date] = make([]APIEvent, 0)
		}

		result[date] = append(result[date], APIEvent{
			ID:    e.ID,
			Title: e.Title,
			Date:  e.Date.Format("2006-01-02"),
		})

		return true
	})

	v, err := json.Marshal(APIResult{result})
	if err != nil {
		throwError(w, ErrorInternalError)
		return
	}
	fmt.Fprint(w, string(v))
}

type mwfunc func(w http.ResponseWriter, r *http.Request)

func mw(f mwfunc) mwfunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		f(w, r)
	}
}

func main() {
	c, err := CreateConfig()
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/create_event", mw(CreateEvent))
	http.HandleFunc("/update_event", mw(UpdateEvent))
	http.HandleFunc("/delete_event", mw(DeleteEvent))
	http.HandleFunc("/events_for_day", mw(GetEventsForDay))
	http.HandleFunc("/events_for_week", mw(GetEventsForWeek))
	http.HandleFunc("/events_for_month", mw(GetEventsForMonth))

	err = http.ListenAndServe(c.APIAddress, nil)
	if err != nil {
		log.Fatalln(err)
	}
}