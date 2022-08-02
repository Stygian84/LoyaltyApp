 import axios from "axios";
 import { useRouter } from "next/router";
 import { useState, useEffect } from "react";
 import Layout from "../components/layout";
 import Link from 'next/dist/client/link';
 import Status from "../components/status";

 export default function TransactionStatus(){
    const [statuses, setStatuses] = useState([]);
    useEffect(() => {
        getStatuses();
    }, []);
    axios.defaults.baseURL = "http://localhost:8080";
    const userID = 41;
    console.log(userID);
    const getStatuses = async () => {
        try{
            const response = await axios.get(`/transaction_status/${userID}`);
            console.log(response.data);
            setStatuses(response.data);
        }
        catch (error){
            console.log("Error", error);
        }
    };

    return(
        <Layout>
            <header>
                <div>
                    <p className="text-lg font-bold text-black ml-10 mt-5">Transaction Status</p>
                </div>
            </header>
            <div className="grid-cols-1 gap-3 md:grid-cols-4 lg:grid-cols">
                {statuses.map((status) => (
                    <Status status={status} key={status.credit_request_id}></Status>
                ))}
            </div>
        </Layout>
    )
 }