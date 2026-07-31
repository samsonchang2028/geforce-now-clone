package main

import (
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Name string `json: "name"`
}
