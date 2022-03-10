var express = require("express");
var router = express.Router();
var crypto = require("crypto");

/* GET random string. */
router.get("/", (req, res) => {
  res.json(crypto.randomBytes(20).toString("hex"));
});

module.exports = router;
