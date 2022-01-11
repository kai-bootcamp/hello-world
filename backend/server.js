const express = require('express');
const path = require('path');
const cors = require('cors')
const app = express();

app.use(cors())

function getRandomInt(max) {
  return Math.floor(Math.random() * max);
}

app.get('/api/getdata', (req, res) => {
  const number = getRandomInt(1000);
  const stringRes = "Random number from 0 to 999: " + number.toString();
  res.setHeader('Content-Type', 'application/json');
  res.end(JSON.stringify({ data: stringRes }));

})

// app.get('/', (req, res) => {
//   res.sendFile(path.join(__dirname, '/views/index.html'));
// })
app.get('/api', (req, res) => {
  res.setHeader('Content-Type', 'application/json');
  res.end(JSON.stringify({ data: "Hello world" }));
})


const PORT = process.env.PORT || 8000;
app.listen(PORT, _ =>{
    console.log(`App deployed on port http://localhost:${PORT}/api`);
})