require("dotenv").config();

const express = require("express");
const jwt = require("jsonwebtoken");

const app = express();
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

const { v4: uuidv4 } = require("uuid");

let books = [
  { id: 1, name: "Roman Lives", author: "Plutarch" },
  { id: 2, name: "Greek Lives", author: "Plurach" },
  { id: 3, name: "Secret of the Vedas", author: "Sri Aurobindo" },
];

app.get("/books", validateJWT, (req, res) => {
  res.type("application/json").status(200).send(books);
});

app.get("/books/:id", validateJWT, (req, res) => {
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

app.post("/books/addnew", validateJWT, (req, res) => {
  let bookId = uuidv4();
  if (req.body.name.length > 0 && req.body.author.length > 0) {
    let newBook = { id: bookId, name: req.body.name, author: req.body.author };
    books.push(newBook);

    res.type("application/json").status(200).send(newBook);
  } else {
    res.status(400).end();
  }
});

app.delete("/books/:id", validateJWT, (req, res) => {
  let bookId = req.params.id;
  console.log(`Deleting book with id: ${JSON.stringify(bookId)}`);

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

//Validate JWT middleware
function validateJWT(req, res, next) {
  //Header authorization header
  const authHeader = req.headers["authorization"];
  //Format: Bearer xxxxx. Get only the xxxx part which is the token
  const token = authHeader && authHeader.split(" ")[1];

  //Token not specifid. Return Error
  if (token == null) {
    res.status(401).end;
  }

  //Verify the token with our own Secret
  jwt.verify(token, process.env.ACCESS_TOKEN_SECRET, (err, user) => {
    //Error in validation, return FORBIDDEN
    if (err) {
      return res.status(403).end();
    }

    //Associate a user object with the request for further processing
    req.user = user;

    //Invoke the next middleware in the chain
    next();
  });
}

const port = 3000;
app.listen(port, () => {
  console.log(`Server started on port ${port}`);
});
