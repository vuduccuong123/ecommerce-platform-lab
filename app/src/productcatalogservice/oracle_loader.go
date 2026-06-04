package main

import (
    "database/sql"
    "os"

    pb "github.com/GoogleCloudPlatform/microservices-demo/src/productcatalogservice/genproto"
    _ "github.com/sijms/go-ora/v2"
)

func loadCatalogFromOracle(catalog *pb.ListProductsResponse) error {

    dsn := os.Getenv("ORACLE_DSN")

    db, err := sql.Open("oracle", dsn)
    if err != nil {
        return err
    }
    defer db.Close()

    rows, err := db.Query(`
        SELECT
            id,
            name,
            description,
            picture,
            price_units,
            price_nanos
        FROM products
    `)

    if err != nil {
        return err
    }

    defer rows.Close()

    catalog.Products = catalog.Products[:0]

    for rows.Next() {

        product := &pb.Product{}
        product.PriceUsd = &pb.Money{
            CurrencyCode: "USD",
        }

        err := rows.Scan(
            &product.Id,
            &product.Name,
            &product.Description,
            &product.Picture,
            &product.PriceUsd.Units,
            &product.PriceUsd.Nanos,
        )

        if err != nil {
            return err
        }

        catalog.Products = append(
            catalog.Products,
            product,
        )
    }

    return nil
}
