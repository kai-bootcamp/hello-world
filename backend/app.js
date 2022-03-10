const express = require("express");
const app = express();
const http = require("http");
const server = http.createServer(app);
const port = process.env.PORT || 9000;
const cors = require("cors");
app.use(cors());

const randomCharactersRouter = require("./routes/randomcharacters.route");
const randomStringRouter = require("./routes/randomstring.route");

app.use(express.json());
app.use(express.urlencoded({ extended: false }));

app.use("/api/random-characters", randomCharactersRouter);
app.use("/api/random-string", randomStringRouter);

/** Catch all request that don't match any route */
app.get("/*", (req, res) => {
  res.json({ success: false, message: "Doesnt match any routes" });
});

server.listen(port, () => {
  console.log("Server start at: 9000");
});
