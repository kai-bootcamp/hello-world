const hw1_url = "http://localhost:8081/hw1";
const hw2_url = "http://localhost:8081/hw2";

const btn = document.querySelector('#hw2');

btn.addEventListener('click', async function () {
    const response = await fetch(hw2_url);
    var rs = await response.json();
    console.log(rs);
    show(rs);
}, false);

function show(data) {
    let result = `${data}`;
    document.querySelector('#randString').innerHTML = result;
}