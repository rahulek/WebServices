const router = require("express").Router();
const passport = require("passport");

//Simple routes to login/logout/test the signed in user.

router.get(
  "/login",
  passport.authenticate("google", {
    scope: ["profile"],
  })
);

router.get("/logout", (req, res) => {
  req.logout();
  res.send("Logged out.");
});

router.get("/google/redirect", passport.authenticate("google"), (req, res) => {
  res.redirect("/testroute");
});

module.exports = router;
