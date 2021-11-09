import axiosClient from './axiosClient.js';

const strApi = {
    getStr() {
        const url = '/reactjs';
        return axiosClient.get(url)
    }
}

export default strApi;