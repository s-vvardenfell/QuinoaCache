package main

type CRUD interface {
	Create(string) error
	Read(string) (string, error)
	Update(string) error
	Delete(string) error
}
