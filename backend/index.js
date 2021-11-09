const express = require("express");
const app = express();
const server = require("http").Server(app);
const cors = require("cors"); 
const port = process.env.PORT || 5000;
app.use(cors());

app.get("/",(req, res) => {
    const response = [
        {
            id: 1,
            name: 'Nguyen Van A',
            age: 22,
            salary: '1000$',
            sex: '0'
        },
        {
            id: 2,
            name: 'Nguyen Van B',
            age: 23,
            salary: '1500$',
            sex: '0'
        },
        {
            id: 3,
            name: 'Nguyen Van C',
            age: 24,
            salary: '2000$',
            sex: '1'
        },
        {
            id: 4,
            name: 'Nguyen Van D',
            age: 22,
            salary: '1000$',
            sex: '0'
        },
    ]
    res.json(response)
});

server.listen(port, () => {
    console.log('Server start at 5000')
});