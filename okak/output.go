package okak

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(value any) {
	intValue, ok := value.(int)
	if ok {
		color.Red("Код ощибки: %d", intValue)
		return
	}
	stringValue, ok := value.(string)
	if ok {
		color.Red(stringValue)
		return // почему тут return а в switch нету зачем мы возвращаем
	}
	errorValue, ok := value.(error)
	if ok {
		color.Red(errorValue.Error())
		return
	}
	color.Red("Неизвестный код ошибки")
	// 	switch t := value.(type) {
	// 	case string:
	// 		color.Red(t)
	// 	case int:
	// 		color.Red("Код ощибки: %d", t)
	// 	case error:
	// 		color.Red(t.Error())
	// 	default:
	// 		color.Red("Неизвестный код ошибки")
	// 	}
}

func sum[T int | string](a, b T) T {
	switch d := any(a).(type) {
	case string:
		fmt.Println(d)
	}
	return a + b
}

type List[T any] struct {
	elements []T
}

func (l *List[T]) addElement() {
	
}