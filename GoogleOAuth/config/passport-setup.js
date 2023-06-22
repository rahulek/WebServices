const passport = require("passport");
const GoogleStrategy = require("passport-google-oauth20").Strategy;
const keys = require("./keys");

//Users are temporarily stored in the memory DB
//Not ideal for the production code where such could
//be stored to a database
const users = [];

//SerializeUser is used by passport to
//create a session-cookie
passport.serializeUser((user, done) => {
  done(null, user.id);
});

//DeserializeUser is used by passport to
//create a user from the session-cookie
passport.deserializeUser((id, done) => {
  const user = users.find((user) => user.id === id);
  done(null, user);
});

//Create a Google OAuth Strategy.
//All OAuth protocol exchanges are
//handled by passport.
passport.use(
  new GoogleStrategy(
    {
      clientID: keys.google.clientID,
      clientSecret: keys.google.clientSecret,
      callbackURL: "/auth/google/redirect",
    },
    (accessToken, refreshToken, profile, done) => {
      //This is called when the signing in is successful
      console.log(
        `Google Login successful: ${profile.id}, ${profile.displayName}`
      );
      //See if the user is already logged in/exists in our system
      let user = users.find((user) => user.id === profile.id);
      if (user != null) {
        //User already exists
        console.log(`User ${profile.displayName} already signed in`);
        //Finish by asking passport for serialize the user (which creates
        //a session cookie)
        done(null, user);
      } else {
        //User signed in does not exist.
        //Create and keep it in the our user db (just a memory array)
        let newUser = { id: profile.id, name: profile.displayName };
        users.push(newUser);
        console.log(`Created new user for ${JSON.stringify(newUser)}`);
        //Finish by asking passport for serialize the user (which creates
        //a session cookie)
        done(null, newUser);
      }
    }
  )
);
