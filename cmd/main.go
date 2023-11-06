package main

import (
	"fmt"
	"log"

	"github.com/oluwadamilarey/OkeleDB/okele"
)

func main() {
	user := map[string]string{
		"name": "David",
		"age":  "25",
	} //  int  , string, []byte, float, ...

	_ = user
	db, err := okele.New()

	if err != nil {
		log.Fatal(err)
	}

	id, err := db.Insert("users", user)
	// coll, err := db.CreateCollection("user")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", id)

	// userData := make(map[string]string)
	// if err := db.View(func(tx *bbolt.Tx) error {
	// 	bucket := tx.Bucket([]byte("users"))

	// 	if bucket == nil {
	// 		return fmt.Errorf("bucket (%s) not found", "users")
	// 	}

	// 	bucket.ForEach(func(k, v []byte) error {
	// 		userData[string(k)] = string(v)
	// 		return nil
	// 	})
	// 	return nil
	// }); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(user)
}
