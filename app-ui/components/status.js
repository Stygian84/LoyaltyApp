export default function Status({status}){
    switch(status.transaction_status){
        case "created":
            var status_str = "Transaction Created";
            break;
        case "pending":
            var status_str = "Transaction Pending";
            break;
        case "approved":
            var status_str = "Transaction Approved";
            break;
        case "rejected":
            var status_str = "Transaction Rejected";
            break;
    }
    return(
        <div className="card flex place-content-between border-2 rounded-xl px-5 py-5 my-5 mx-2 border-black">
            <div className="">
                <p className="font-bold">{status.program}</p>
                <br></br>
                <p>Total Reward Points to Transfer:</p>
                <p>{status.credit_used}</p>
                <br></br>
                <p>Rewards To Be Received:</p>
                <p>{status.credit_to_receive}</p>
            </div>
            <div className="text-center">
                <p>{status_str}</p>
            </div>
        </div>
    )
}