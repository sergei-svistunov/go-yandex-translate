# go-yandex-translate

Go Yandex translate API client

API documentation: https://tech.yandex.com/translate/doc/dg/concepts/api-overview-docpage/

## Usage
```go
yt := translate.New("API KEY") // get the key from https://translate.yandex.com/developers/keys

langs, err := yt.GetLanguages("ru")

lang, err := yt.Detect("Test", nil)

texts, err := yt.Translate([]string{"Test", "Hello"}, "en-ru", "plain")
```