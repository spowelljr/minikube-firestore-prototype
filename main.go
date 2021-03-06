package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

type State struct {
	Capital    string  `firestore:"capital"`
	Population float64 `firestore:"pop"`
}

func main() {
	if err := os.Setenv("FIRESTORE_EMULATOR_HOST", "192.168.58.2:8086"); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "test")
	if err != nil {
		log.Fatalf("NewClient: %v", err)
	}
	ny := client.Doc("States/NewYork")
	getDoc(ctx, ny)
	if _, err := ny.Create(ctx, State{
		Capital:    "Albany",
		Population: 19.8,
	}); err != nil {
		log.Fatalf("Create: %v", err)
	}
	getDoc(ctx, ny)
}

func getDoc(ctx context.Context, doc *firestore.DocumentRef) {
	docsnap, err := doc.Get(ctx)
	if err != nil {
		log.Printf("no existing data found: %v\n", err)
	} else {
		log.Println(docsnap.Data())
	}
}
