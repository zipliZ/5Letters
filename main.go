package main

import (
	"5Words/connection"
	"encoding/json"
	"fmt"
	"github.com/restream/reindexer/v3"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var client http.Client

type MatchData struct {
	Word string `json:"Word"`
}

type ResponseData struct {
	Response []string `json:"Response"`
}

func GetMatches(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var data MatchData
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	word := data.Word
	letters := strings.Split(word, "")

	var hiddenLetters []string
	var formattedRespData []string

	if _, found := db.Query("words").
		Where("Word", reindexer.EQ, word).
		Get(); found {
		cookies := r.Cookies()
		for _, cookie := range cookies {
			decodedValue, _ := url.QueryUnescape(cookie.Value)
			hiddenLetters = strings.Split(decodedValue, "")

		}
		for i, letter := range letters {
			value := 0
			for i2, hiddenLetter := range hiddenLetters {
				if letter == hiddenLetter && i == i2 {
					value = 2
					break
				} else if letter == hiddenLetter {
					value = 1
				}
			}
			formattedRespData = append(formattedRespData, letter+"/"+strconv.Itoa(value))
		}

	} else {
		for _, letter := range letters {
			formattedRespData = append(formattedRespData, letter+"/"+"-1")
		}
	}
	respData := ResponseData{
		Response: formattedRespData,
	}

	jsonData, err := json.Marshal(respData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}

var db *reindexer.Reindexer

func StartGame(w http.ResponseWriter, r *http.Request) {

	iterator := db.ExecSQL("SELECT ID FROM words")
	CountWords := iterator.Count()
	var hiddenWord string
	elem, found := db.Query("words").
		Where("id", reindexer.EQ, rand.Intn(CountWords)+1).
		Get()
	if found {
		hiddenWord = elem.(*connection.Word).Word
	}

	encodedValue := url.QueryEscape(hiddenWord)
	cookie := &http.Cookie{
		Name:  "hiddenWord",
		Value: encodedValue,
		Path:  "/",
	}
	http.SetCookie(w, cookie)

	http.ServeFile(w, r, "static/index.html")

}

func main() {
	defer db.Close()
	db = connection.InitConnection()
	fmt.Println("Server started")
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/5Letters/", StartGame)
	http.HandleFunc("/api/GetMatches/", GetMatches)

	http.ListenAndServe(":8080", nil)

}
