import axios from "axios";
import { useSession } from "next-auth/react";
import { useRouter } from "next/router";
import Layout from "../components/layout";
import LabelContent from "../components/labelContent";
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
    <div className="flex flex-col mb-5">
      <LabelContent title="Program Name">
        <h2 className="font-bold text-2xl capitalize">{program.name}</h2>
      </LabelContent>
      <LabelContent title="Program Description">
        <h3 className="text-xl">{program.description}</h3>
      </LabelContent>
      <LabelContent title="Partner Code">
        <p>{program.partner_code}</p>
      </LabelContent>
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

const Transaction = () => {
  const { data: session } = useSession();
  const [program, setProgram] = useState({});
  const [creditToUse, setCreditToUse] = useState(0);
  const [rewardExpected, setRewardExpected] = useState(0);
  const [membershipID, setMembershipID] = useState(0);
  const [referenceDoc, setReferenceDoc] = useState(null);
  const [errorResponse, setErrorResponse] = useState(null);
  const [user, setUser] = useState([]);

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

  const getUser = async () => {
    try {
      console.log("session", session);
      const response = await axios.get(
        `/getUserbyEmail/${session?.user.email}`
      );
      setUser(response.data[0]);
    } catch (error) {
      console.log("err", error);
    }
  };
  useEffect(() => {
    if (router.isReady) {
      getUser();
      getProgram();
    }
  }, [router.isReady, session]);

  const initateTransaction = async () => {
    try {
      const data = {
        userId: user.id,
        creditToTransfer: parseFloat(creditToUse),
        membershipId: membershipID,
        programId: program.id,
      };
      const response = await axios.post("/initTransaction", (data = data));
      setReferenceDoc(response.data);
    } catch (error) {
      console.log("err", error.response.data.error);
      setReferenceDoc(null);
      setErrorResponse(error.response.data.error);
    }
  };
  const checkReward = async () => {
    if (user.credit_balance < creditToUse) {
      setRewardExpected(0);
    } else {
      try {
        const data = {
          userId: user.id,
          creditToTransfer: parseFloat(creditToUse),
          membershipId: membershipID,
          programId: program.id,
        };
        const response = await axios.post("/checkReward", (data = data));

        setRewardExpected(response.data.Amount);
      } catch (error) {
        console.log("err", error);
      }
    }
  };

  return (
    <Layout>
      <div className="flex items-center justify-center max-w-7xl m-auto min-h-screen">
        <div className="flex flex-col items-center justify-center border-2 rounded-xl px-20 py-40 border-black">
          <h1 className="font-bold text-4xl mb-6">Transaction Page</h1>

          <ProgramSection program={program} />
          <div className="mr-30">
            <label className="mr-2">Enter Membership ID: </label>
            <input
              className="border-black border-2 rounded-sm mb-5 "
              onChange={(e) => {
                setMembershipID(e.target.value);
              }}
            />
          </div>

          <div>
            <label className="mr-2">credit to transfer:</label>
            <input
              type="number"
              className="border-black border-2 rounded-sm"
              onChange={(e) => {
                setCreditToUse(e.target.value);
              }}
            />
            <button
              className="font-bold w-auto h-auto bg-sky-300 p-2 mt-4"
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
            className="font-bold w-auto h-auto bg-sky-300 p-2 mt-4"
            onClick={initateTransaction}
          >
            Submit Request
          </button>
          {referenceDoc && <RenderCreditRequest referenceDoc={referenceDoc} />}
          <p>{errorResponse}</p>
        </div>
      </div>
    </Layout>
  );
};
export default Transaction;
