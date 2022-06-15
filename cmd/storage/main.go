package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
)
import "github.com/sam6699/storage/internal/storage"

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	file2, err := st.GetById(file.ID)

	fileID, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	file3, err := st.GetById(fileID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hello", file)
	fmt.Println("hello", file2)
	fmt.Println("hello", file3)
}
