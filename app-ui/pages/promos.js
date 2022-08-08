import Promotion from "../components/promotion";
import Layout from "../components/layout";
import Nav from "../components/Nav";
const axios = require("axios").default;
import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import { useSession } from "next-auth/react";

export default function Home() {
  const [promo, setPromo] = useState([]);
  const [program, setProg]=useState({});
  const router = useRouter();
  axios.defaults.baseURL="http://localhost:8080"
  const {data:session}=useSession();

  //runs the function once when the page is loaded

  const displayProgram = async () => {
    const response1 = await axios.get("/listPromo");
    setPromo(response1.data);
  };

  const getProg=async ({progID})=>{
    const response1 = await axios.get(`/GetLoyaltyId/${progID}`);
    setProg(response1.data);
  }

  useEffect(() => {
    if (router.isReady) {
      displayProgram();
    
      
    }

    console.log(promo)
  }, [router.isReady, session]);

  return (
    <Layout title={`Offers Page `}>
      <Nav />
      <h1 className="px-4 h-12 text-lg font-bold mt-10">Offers</h1>
      <div className="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3 px-4">
        
        {promo.map((promo) => (
          <Promotion  promo={promo} key={promo.id}></Promotion>
        ))}
      </div>
      <hr className="w-full my-5"></hr>
    </Layout>
  );
}
