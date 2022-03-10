var express = require("express");
var router = express.Router();

/** Get random characters */
router.get("/", (req, res) => {
  const { length = 5 } = req.query;
  res.json(randomCharacters(length));
});

function randomCharacters(length) {
  var result = "";
  var characters =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  var charactersLength = characters.length;
  for (var i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }
  return result;
}

module.exports = router;
