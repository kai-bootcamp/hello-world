const express = require("express");
const app = express();
const server = require("http").Server(app);
const cors = require("cors");
const randomstring = require("randomstring");
const port = process.env.PORT || 5000;
app.use(cors());

app.get("/",(req, res) => {
    res.json(randomstring.generate());
});

server.listen(port, () => {
    console.log('Server start at 5000')
});