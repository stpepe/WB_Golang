package main

import "time"

type APIError struct {
	Error string `json:"error"`
}

type APIResult struct {
	Result interface{} `json:"result"`
}

type APIEvent struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

type APIEventID struct {
	ID string `json:"id"`
}

type Event struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}