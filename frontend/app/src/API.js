const hostAPI = "http://localhost:9000";

const API = {
  getRandomString: async () => {
    const endpoint = `${hostAPI}/api/random-string`;
    return await (await fetch(endpoint)).json();
  },
  getRandomCharacters: async (length) => {
    const endpoint = `${hostAPI}/api/random-characters?length=${length}`;
    return await (await fetch(endpoint)).json();
  },
};

export default API;
