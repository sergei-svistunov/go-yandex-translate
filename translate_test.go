package translate_test

import (
	"github.com/sergei-svistunov/go-yandex-translate"
	"os"
	"testing"
)

func TestTranslate_GetLanguages(t *testing.T) {
	yt := translate.New(os.Getenv("YA_API_KEY"))

	langs, err := yt.GetLanguages("")
	if err != nil {
		t.Fatal(err)
	}

	if len(langs) == 0 {
		t.Fatalf("No languages were received")
	}

	for id, name := range langs {
		t.Logf("%s: %s\n", id, name)
	}
}

func TestTranslate_Detect(t *testing.T) {
	yt := translate.New(os.Getenv("YA_API_KEY"))

	lang, err := yt.Detect("Test", nil)
	if err != nil {
		t.Fatal(err)
	}

	if lang != "en" {
		t.Fatalf("Lang is %s, not en")
	}
}

func TestTranslate_Translate(t *testing.T) {
	yt := translate.New(os.Getenv("YA_API_KEY"))

	texts, err := yt.Translate([]string{"Test", "Hello"}, "en-ru", "plain")
	if err != nil {
		t.Fatal(err)
	}

	if len(texts) != 2 {
		t.Fatalf("Length is %d, not 2", len(texts))
	}
}
