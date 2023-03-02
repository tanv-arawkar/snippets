module main

go 1.20

replace postgres_conn => ../postgres_conn

require postgres_conn v0.0.0-00010101000000-000000000000

require github.com/lib/pq v1.10.7 // indirect
