## Библиотека предназначенная для получения ответов от сервера.

#### So installieren bibliothek

```
go get github.com/Aurivena/answer
```

#### Beispiele interaktion fur bibliothek
```go
	type Human struct {
    	Name string
    	Age  int
    }

    func main() {
    	c := gin.Context{}

    	output := Human{
    		Name: "Боба",
    		Age:  100,
    	}

    	answer.SendSuccess(&c, "Все круто", answer.OK)
    	answer.SendError(&c, "Все Плохо", answer.BadRequest)
    	answer.SendResponseSuccess(&c, output, answer.OK)
    }


    answer.AppendCode(2, "Test2")
    answer.ChangeCode(2,"Тест 3")
    answer.DeleteCode(2)
``
