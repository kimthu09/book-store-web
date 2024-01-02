import { endPoint } from "@/constants";
import { getApiKey } from "../auth/action";
import axios from "axios";

export default async function getReport({
    timeFrom,
    timeTo,
    type,
}: {
    timeFrom: number,
    timeTo: number,
    type: string,

}) {
    const url = `${endPoint}/v1/reports/${type}`;

    const data = {
        timeFrom,
        timeTo
    };

    const token = await getApiKey();

    const headers = {
        accept: "application/json",
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
    };

    // Make a POST request with headers
    const res = axios
        .post(url, data, { headers: headers })
        .then((response) => {
            if (response) return response.data;
        })
        .catch((error) => {
            console.error("Error:", error);
            return error.response.data;
        });
    return res;
}