package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {


}
func server(){
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resBody, err := io.ReadAll(r.Body)

		if err != nil {
			log.Println(err.Error())
			return
		}
		defer r.Body.Close()

		var see any
		json.Unmarshal(resBody, &see)
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusAccepted)

		res, err := json.Marshal(map[string]any{"message": "fuck u from my heart"})

		if err != nil {
			log.Println(err.Error())
			return
		}
		w.Write(res)

	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func gibberish() {
	file, err := os.Create("./palindrome/palindrom.go")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.WriteString("I am a winner")

	writer.Flush()

	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
}

func ReadWithBuffer(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	buffer := make([]byte, 8*100)

	reader := bufio.NewReader(file)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
		}

		fmt.Println(string(buffer[:n]))
	}
}

func ReadWithScanner(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func CallPost(url string) {
	body := map[string]any{"name": "bashir"}
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err.Error())
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err.Error())
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer res.Body.Close()

	r, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var jsonData any
	json.Unmarshal(r, &jsonData)
	fmt.Println(jsonData)

}
