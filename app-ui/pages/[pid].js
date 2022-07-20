import axios from "axios";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
// export async function getStaticProps(context) {
//   const programid = context.params.pid;
//   try {
//     const response = await axios.get(`/loyalty${programid}`);
//   } catch (error) {
//     console.log("Error", error);
//   }
//
//   return {
//     props: {},
//   };
// }
axios.defaults.baseURL = "http://localhost:8080";

const ProgramSection = ({ program }) => {
  return (
    <div className="flex flex-col text-center">
      <h2 className="font-bold text-2xl capitalize">{program.name}</h2>
      <h3 className="text-xl">{program.description}</h3>
      <p>{program.partner_code}</p>
    </div>
  );
};

const RenderCreditRequest = ({ referenceDoc }) => {
  return (
    <div className="text-center">
      <h1>Refernce Number: {referenceDoc.reference_number}</h1>
      <h3>Transaction Time: {referenceDoc.transaction_time.Time}</h3>
      <p>Credit Used: {referenceDoc.credit_used}</p>
      <p>Reward Expected: {referenceDoc.reward_should_receive}</p>
      <p>Transaction Status: {referenceDoc.transaction_status}</p>
    </div>
  );
};

const Transaction = ({ props }) => {
  const [program, setProgram] = useState({});
  const [creditToUse, setCreditToUse] = useState(0);
  const [rewardExpected, setRewardExpected] = useState(0);
  const [referenceDoc, setReferenceDoc] = useState(null);
  const user = 1;
  const router = useRouter();
  const getProgram = async () => {
    const { pid } = router.query;
    try {
      const response = await axios.get(`/loyalty/${pid}`);
      setProgram(response.data);
    } catch (error) {
      console.log("Error", error);
    }
  };
  useEffect(() => {
    getProgram();
  }, [router]);

  const initateTransaction = async () => {
    try {
      const data = {
        userId: user,
        creditToTransfer: parseFloat(creditToUse),
        membershipId: "1005610",
        programId: program.id,
      };
      const response = await axios.post("/initTransaction", (data = data));
      setReferenceDoc(response.data);
    } catch (error) {
      console.log("err", error);
    }
  };
  const checkReward = async () => {
    try {
      const data = {
        userId: user,
        creditToTransfer: parseFloat(creditToUse),
        membershipId: "1005610",
        programId: program.id,
      };
      const response = await axios.post("/checkReward", (data = data));
      setRewardExpected(response.data.Amount);
    } catch (error) {
      console.log("err", error);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center max-w-7xl bg-teal-300 m-auto min-h-screen">
      <h1 className="font-bold text-4xl">Transaction Page</h1>

      <ProgramSection program={program} />
      <div>
        <label className="mr-2">credit to transfer:</label>
        <input
          type="number"
          onChange={(e) => {
            setCreditToUse(e.target.value);
          }}
        />
        <button
          className="w-auto h-auto text-sm bg-red-200 p-1 ml-2 rounded-sm"
          onClick={checkReward}
        >
          Check Reward
        </button>
      </div>
      {rewardExpected != 0 && (
        <p className="text-red-500 font-2xl">
          your reward expected is: {rewardExpected}
        </p>
      )}
      <button
        className="font-bold w-auto h-auto bg-red-200 p-2 mt-4"
        onClick={initateTransaction}
      >
        Submit Request
      </button>
      {referenceDoc && <RenderCreditRequest referenceDoc={referenceDoc} />}
    </div>
  );
};
export default Transaction;
