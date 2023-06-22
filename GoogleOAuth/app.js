const express = require("express");
const authRoutes = require("./routes/auth-routes");
const passport = require("passport");
const passportSetup = require("./config/passport-setup");
const keys = require("./config/keys");
const cookieSession = require("cookie-session");

const app = express();

// set up session cookies
app.use(
  cookieSession({
    maxAge: 24 * 60 * 60 * 1000,
    keys: [keys.session.cookieKey],
  })
);

// initialize passport
app.use(passport.initialize());
app.use(passport.session());
app.use("/auth", authRoutes);

// Middleware to check if the user is authenticated
function isUserAuthenticated(req, res, next) {
  if (req.user) {
    next();
  } else {
    res.send("You must login!");
  }
}

//Just a test route
//"isUserAuthenticated" is a middlware function that
//make sure that we've a valid session cookie
app.get("/testroute", isUserAuthenticated, (req, res) => {
  res.send(
    `You have reached the test route, logged in as: ${req.user.name}, id: ${req.user.id}`
  );
});

//Start listening
app.listen(3000, () => {
  console.log("Server listening for requests on port 3000.");
});
