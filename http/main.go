package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

func main() {
	cities := []string{
		"Владивосток",
		"Вьетнам",
	}

	for _, city := range cities {
		result, err := whatWeather(city)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Вывод до regexp: (Владивосток: Прогноз) 
		// После: (Владивосток: прогноз)
		r := regexp.MustCompile("Прогноз")

		fmt.Println(city+":", r.ReplaceAllString(result, "прогноз"))
	}
}

func whatWeather(city string) (string, error) {
	baseURL := "https://bba4h86i8icpvi5ot97n.containers.yandexcloud.net/go/weather"

	params := url.Values{}
	params.Add("0", "")
	params.Add("T", "")
	params.Add("M", "")
	params.Add("city", city)
	params.Add("lang", "ru")

	urlInstance, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse url: %w", err)
	}

	urlInstance.RawQuery = params.Encode()

	request, err := http.NewRequest(
		"GET", urlInstance.String(), nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create new request: %w", err)
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("error on server: %w", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read body: %w", err)
	}

	result := string(body)

	return result, nil
}

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// )

// func main() {
// 	baseURL := "https://bba4h86i8icpvi5ot97n.containers.yandexcloud.net/go/weather"

// 	params := url.Values{}
// 	params.Add("0", "")
// 	params.Add("T","")
// 	params.Add("M","")

// 	urlInstance, err := url.Parse(baseURL)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	urlInstance.RawQuery = params.Encode()

// 	request, err := http.NewRequest(
// 		"GET", urlInstance.String(), nil,
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	request.Header.Add("Accept-Language", "ru")

// 	client := &http.Client{}

// 	response, err := client.Do(request)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(response.Status)

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(string(body))
// }

// func main() {
// 	request, err := http.NewRequest(
// 		"GET","https://habr.com/", nil,
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	request.Header.Add("Language", "ru")

// 	client := &http.Client{}

// 	response, err := client.Do(request)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(response.Status)
// 	time.Sleep(1 * time.Second)

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(string(body))
// }

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// )

// func main() {
// 	query := "https://yandex.ru/search/"

// 	params := url.Values{}
// 	params.Add("text", "linux")
// 	params.Add("lr", "213")

// 	urlInstance, err := url.Parse(query)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	urlInstance.RawQuery = params.Encode()

// 	url := urlInstance.String()
// 	fmt.Println(url)

// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	fmt.Println(response.StatusCode)

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(string(body))
// }
