package repository

import (

)

type HelloRepository interface {
	// "Hello, World!"を返す
	Hello() string

	// "Hello, {name}!"を返す
	HelloName(name string) string
}
