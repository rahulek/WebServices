//We would store secret keys in the .env file
//and it won't be in the GitHub.
require("dotenv").config();

const express = require("express");
const jwt = require("jsonwebtoken");

const app = express();

//Load the JSON parsing ability
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

//JWT Refresh token is used when the Access token expires.
//Storing refresh tokens inside the plain array is ONLY
//for the demo purposes. Such important security tokens should
//be in the DB.
let refreshTokens = [];

//Login (POST) entry point. User sends in
//"username" and "password". JWT is only
//an authorization step. We don't do
//any authentication here. Its assumed that
//the authentication is successful.
//We create access and refresh JWT tokens and
//return them.
app.post("/login", (req, res) => {
  //Authorize (create a JWT token)
  const username = req.body.username;
  const user = { name: username };

  //Generate ACCESS and REFRESH tokens
  const accessToken = generateAccessToken(user);
  const refreshToken = jwt.sign(user, process.env.REFRESH_TOKEN_SECRET);

  //Store the refresh token so that access tokens for the user
  //can be generated again.
  refreshTokens.push(refreshToken);

  //Return JSON response with Access and Refresh Tokens
  res.json({ accessToken, refreshToken });
});

//Refresh Token endpoint
//Invoked when the Access Token expires
//Refresh token is passed in the POST body
//This token is validated before new Access Token
//can be generated
app.post("/token", (req, res) => {
  //Take the refresh token from the POST Body
  const refreshToken = req.body.token;

  //If no token provides, return UNAUTHORIZED error
  if (refreshToken == null) res.status(401).end();

  //See if the token specified is really the token we sent.
  if (!refreshTokens.includes(refreshToken)) {
    //Not created by us. Send FORBIDDEN
    res.status(403).end();
  }

  //Token is our own. Validate it and get the encoded "user" information
  jwt.verify(refreshToken, process.env.REFRESH_TOKEN_SECRET, (err, user) => {
    //Verification failed
    if (err) {
      //Send FORBIDDEN
      return res.status(403).end();
    }

    //All well. Generate a new access token
    const accessToken = generateAccessToken({ name: user.name });

    //Return the new access token
    res.json({ accessToken });
  });
});

//Logout entry point
//Delete the token from the refresh tokens so that
//new access token can't be minted indefinitely.
app.delete("/logout", (req, res) => {
  refreshTokens = refreshTokens.filter((token) => token !== req.body.token);
  res.status(204).end();
});

//Support function to generate Access token
//Encode the user info to send across and sign it with out Secret
//Token expires in 1m (60s)
function generateAccessToken(user) {
  return jwt.sign(user, process.env.ACCESS_TOKEN_SECRET, { expiresIn: "60s" });
}

//Start listening
app.listen(4000, () => {
  console.log("Auth Server running on port 4000");
});
