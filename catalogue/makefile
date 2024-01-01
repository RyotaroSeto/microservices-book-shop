protoc_gen
  protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/book/catalogue.proto
grpc_list
  grpcurl -plaintext localhost:50051 list
grpc_book_catalogue
  grpcurl -plaintext localhost:50051 list book.Catalogue
get_book
  grpcurl -plaintext -d '{"id": "1"}' localhost:50051 book.Catalogue.GetBook
