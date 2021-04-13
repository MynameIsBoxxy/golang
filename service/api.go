package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-kit/kit/log"
)

type Service interface {
	Validate(ctx context.Context, input string) (string, error)
	Fix(ctx context.Context, input string) (string, error)
}

type service struct {
	logger log.Logger
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) Fix(ctx context.Context, input string) (string, error) {
	fmt.Println("sdfsdfs")
	br := "(){}[]"
	ob := "[({"
	cb := "])}"
	var strc string
	for i := 0; i < len(input); i++ {
		cur := input[i]

		ind := strings.Index(br, string(cur))

		if ind%2 != 0 { // закрывающая скобка
			index := strings.Index(cb, string(br[ind])) // находим индекс закрывающей скобки

			obI := string(ob[index]) // по этому индексу берем открывающую скобку

			if i-1 != -1 && string(input[i-1]) == obI {
				strc = strc + string(br[ind-1])
				strc = strc + string(br[ind])

			}

		}
	}
	return strc, nil
}

func (s service) Validate(ctx context.Context, input string) (string, error) {
	resp := "Not Balanced"
	if s.isBalance(input) {
		resp = "Balanced"
	}
	return resp, nil
}

func (s service) isBalance(str string) bool {
	br := "(){}[]"
	st := make([]string, 0, 10)

	for i := 0; i < len(str); i++ {
		cur := str[i]

		ind := strings.Index(br, string(cur))

		if ind != -1 {
			if ind%2 != 0 {
				if len(st) == 0 {
					return false
				}
				last := st[len(st)-1]
				st = st[:len(st)-1]
				if last != string(br[ind-1]) {
					return false
				}
			} else {
				st = append(st, string(str[i]))
			}
		}
	}
	if len(st) != 0 {
		return false
	} else {
		return true
	}
}
