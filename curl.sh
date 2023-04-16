#Get
curl localhost:8080/books
#Create book from file
curl localhost:8080/books --header "Content-Type: application/json" --request "POST" -d @book.json
#Checkout book
curl "localhost:8080/books/checkout?id=2" --request "PATCH"
#Return book
curl "localhost:8080/books/return?id=4" --request "PATCH"
