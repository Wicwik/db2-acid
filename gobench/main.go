package main

import (
	cryptorand "crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	mathrand "math/rand"
	"testing"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres@localhost:5432/oz?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for j := 0; j < 5; j++ {
		b := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				docName := "Generated doc " + randHex(10)
				department := "Department " + randHex(10)
				contractedAmount := math.Round(mathrand.Float64() * 1_000_000)

				_, err := db.Exec(`INSERT INTO documents(name, type, created_at, department, contracted_amount) VALUES ($1, 'MyType', NOW(), $2, $3)`, docName, department, contractedAmount)
				if err != nil {
					log.Fatal(err)
				}
			}
		})

		fmt.Print(b)
		fmt.Printf(" %f \n", float64(b.N)/(float64(b.NsPerOp())/1000000000))
	}
}

func randHex(n int) string {
	bytes := make([]byte, n)
	if _, err := cryptorand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}
