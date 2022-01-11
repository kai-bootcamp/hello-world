import React from "react";
import { useEffect, useState} from "react";
import numberApi from "../api/numberApi";
const MainPage = () => {
    const [randomString , setRandomString] = useState("");

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await numberApi.getNumber();
                console.log(response);
                setRandomString(response.data);
            }
            catch (error) {
                console.log("Fail to load random string " + error);
            }
        }
        fetchData();
    }, []);

    return (
        <div>
        <h1>Kardia bootcamp</h1>
        <p>{randomString}</p>
        </div>
    );
}

export default MainPage;