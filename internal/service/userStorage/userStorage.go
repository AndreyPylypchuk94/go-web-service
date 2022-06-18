package userStorage

import (
	"golang.org/x/exp/slices"
	"pylypchuk.home/pkg/utils"
)

var userStorage = make([]int64, 0)

func Contains(id int64) bool {
	return slices.Contains(userStorage, id)
}

func Add(id int64) {
	if !Contains(id) {
		userStorage = append(userStorage, id)
	}
}

func Delete(id int64) {
	if Contains(id) {
		userStorage = utils.SliceFilter(
			userStorage,
			func(el int64) bool { return el != id },
		)
	}
}
