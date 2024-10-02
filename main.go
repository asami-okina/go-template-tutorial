package main

import (
	"log"
	"os"
	"strings" 
	"text/template"
	"time"
)

type Person struct {
	Name string
	Age  int
}

type Product struct {
	Name  string
	Price float64
}

func main() {
	// テンプレート文字列の定義
	templateString := `
1. 基本的な変数の展開:
   Hello, {{.Name}}! You are {{.Age}} years old.

2. 条件分岐:
   {{if ge .Age 18}}
   You are an adult.
   {{else}}
   You are a minor.
   {{end}}

3. ループ処理:
   Your shopping list:
   {{range .ShoppingList}}
   - {{.Name}}: ${{.Price}}
   {{end}}

4. with文の使用:
   {{with .BestFriend}}
   Your best friend is {{.Name}} ({{.Age}} years old).
   {{else}}
   You don't have a best friend.
   {{end}}

5. テンプレート関数の使用:
   Your name in uppercase: {{upper .Name}}
   Today's date: {{formatDate .Today}}

6. パイプライン:
   Your name has {{.Name | len}} characters.

7. 変数の定義と使用:
   {{$discount := 0.1}}
   {{range .ShoppingList}}
   - {{.Name}}: ${{.Price}} (Discounted: ${{multiply .Price (subtract 1 $discount)}})
   {{end}}
`

	// データの準備
	data := struct {
		Person
		ShoppingList []Product
		BestFriend   *Person
		Today        time.Time
	}{
		Person: Person{
			Name: "Alice",
			Age:  25,
		},
		ShoppingList: []Product{
			{Name: "Apple", Price: 0.5},
			{Name: "Banana", Price: 0.3},
			{Name: "Orange", Price: 0.4},
		},
		BestFriend: &Person{Name: "Bob", Age: 27},
		Today:      time.Now(),
	}

	// カスタム関数の定義
	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
		"formatDate": func(t time.Time) string {
			return t.Format("2006-01-02")
		},
		"multiply": func(a, b float64) float64 {
			return a * b
		},
		"subtract": func(a, b float64) float64 {
			return a - b
		},
	}

	// テンプレートの作成と実行
	tmpl, err := template.New("test").Funcs(funcMap).Parse(templateString)
	if err != nil {
		log.Fatalf("テンプレートのパースに失敗しました: %v", err)
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("テンプレートの実行に失敗しました: %v", err)
	}
}
