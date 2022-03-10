export async function getDataFromServer() {
    const res = await fetch(`${process.env.SERVER_URL}/hello-world`);
    return res.json();
}