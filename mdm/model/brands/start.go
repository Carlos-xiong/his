package main

import (
	"context"
	"log"

	"his/mdm/model/brands/ent"

	"his/mdm/model/brands/ent/mdmbrands"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=his password=6286964 sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	} else {
		log.Println("schema resources created successfully")
	}
	//查找ID
	brand, err := client.MdmBrands.Get(context.Background(), 9)
	if err != nil {
		log.Fatalf("failed getting brand: %v", err)
	}
	log.Printf("got brand: %v", brand)
	//按条件查找
	brands, err := client.MdmBrands.
		Query().
		Where(mdmbrands.BrandNameEQ("老骡子 or true")).
		All(context.Background())
	if err != nil {
		log.Fatalf("failed getting brands: %v", err)
	}
	log.Printf("got brands: %v", brands)
}
