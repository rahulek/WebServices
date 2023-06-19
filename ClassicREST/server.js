express = require("express");
app = express();
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

const { v4: uuidv4 } = require("uuid");

let books = [
  { id: 1, name: "Roman Lives", author: "Plutarch" },
  { id: 2, name: "Greek Lives", author: "Plurach" },
  { id: 3, name: "Secret of the Vedas", author: "Sri Aurobindo" },
];

app.get("/books", (req, res) => {
  res.type("application/json").status(200).send(books);
});

app.get("/books/:id", (req, res) => {
  let id = req.params.id;
  let book = books.filter(({ id: bookId }) => {
    return id == bookId;
  });

  if (book.length != 0) {
    res.type("application/json").status(200).send(book);
  } else {
    res.status(404).end();
  }
});

app.post("/books/addnew", (req, res) => {
  let bookId = uuidv4();
  if (req.body.name.length > 0 && req.body.author.length > 0) {
    let newBook = { id: bookId, name: req.body.name, author: req.body.author };
    books.push(newBook);

    res.type("application/json").status(200).send(newBook);
  } else {
    res.status(400).end();
  }
});

app.delete("/books/:id", (req, res) => {
  let bookId = req.params.id;
  console.log(`Deleting book with id: ${bookId}`);

  let startSize = books.length;
  books = books.filter(({ id }) => {
    let startsWith = String(id).startsWith(String(bookId));
    return id != bookId && !startsWith;
  });

  if (books.length == startSize - 1) {
    //We deleted one book
    res
      .type("application/json")
      .status(200)
      .send(JSON.stringify({ deletedBookID: bookId }));
  } else {
    //Failed to delete the books
    res
      .type("text/plain")
      .status(404)
      .send(`Could not find the book with Id: ${bookId}`);
  }
});

const port = 3000;
app.listen(port, () => {
  console.log(`Server started on port ${port}`);
});
