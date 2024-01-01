## Setup

1. npm install
2. npm start

## Query

```
query Book($bookId: Int) {
  book(id: $bookId) {
    id
    title
    author
    price
  }
}
```

### variable

```
{
    "bookId": 1
}
```
