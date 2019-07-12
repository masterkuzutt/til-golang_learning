package presenter

import (
	"fmt"
)

type ToConsolePresenter struct {
	hoge string
}

func (p *ToConsolePresenter) Emit(data string) {
	fmt.Println(data)
}
