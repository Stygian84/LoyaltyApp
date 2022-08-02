import axios from "axios";
import { useState, useEffect } from "react";
import Layout from "../components/layout";
import Status from "../components/status";
import { useSession } from "next-auth/react";

export default function TransactionStatus() {
    const [statuses, setStatuses] = useState([]);
    const [user, setUser] = useState({});

    const { data: session } = useSession();
    const getUser = async () => {
        try {
            if (session) {
                const response = await axios.get(
                    `/getUserbyEmail/${session.user.email}`
                );
                const response1 = await axios.get(
                    `/transaction_status/${response.data[0].id}`
                );
                setStatuses(response1.data);
            }
        } catch (error) {
            console.log(error);
            // console.log("err", error.response.data.error);
        }
    };
    useEffect(() => {
        getUser();
    }, [session]);

    return (
        <Layout title={"Transaction Status"}>
            <header>
                <div>
                    <p className="text-lg font-bold text-black ml-10 mt-5">
                        Transaction Status
                    </p>
                </div>
            </header>
            <div className="grid-cols-1 gap-3 md:grid-cols-4 lg:grid-cols">
                {statuses.map((status) => (
                    <Status
                        status={status}
                        key={status.credit_request_id}
                    ></Status>
                ))}
            </div>
        </Layout>
    );
}
