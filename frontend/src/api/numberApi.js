import axiosClient from "./axiosClient";

const numberApi = {
    get: () =>{
        const url = "/";
        return axiosClient.get(url);
    },
    getNumber: () =>{
        const url = "/getdata";
        return axiosClient.get(url);
    },
}

export default numberApi;