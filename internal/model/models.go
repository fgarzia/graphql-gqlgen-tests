// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Employee struct {
	ID           *int        `json:"id"`
	Name         *string     `json:"name"`
	Supervisor   *int        `json:"supervisor"`
	Role         *Role       `json:"role"`
	Subordinates []*Employee `json:"subordinates"`
}

type Role struct {
	ID          *int    `json:"id"`
	Description *string `json:"description"`
	Level       *int    `json:"level"`
}
