import Promotion from "../components/promotion";
import Layout from "../components/layout";
import Nav from "../components/Nav";
const axios = require("axios").default;
import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import { useSession } from "next-auth/react";

export default function Home() {
  const [promo, setPromo] = useState([]);
  const [user, setUser] = useState({});

  const router = useRouter();
  axios.defaults.baseURL = "http://localhost:8080";
  const { data: session } = useSession();

  const displayProgram = async () => {
    if (session) {
      const response = await axios.get(`/getUserbyEmail/${session.user.email}`);
      const response1 = await axios.get("/listPromo");
      if (response1.data != null) {
        var plist = [];
        for (var i = 0; i < response1.data.length; i++) {
          if (response1.data[i].card_tier == response.data[0].card_tier) {
            plist.push(response1.data[i]);
          }
        }

        setPromo(plist);
      }
    }
  };

  useEffect(() => {
    if (router.isReady) {
      displayProgram();
    }
  }, [user.card_tier, router.isReady, session]);

  return (
    <Layout title={`Offers Page `}>
      <Nav />
      <h1 className="px-4 h-12 text-lg font-bold mt-10">Offers</h1>
      <div className="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3 px-4">
        {promo &&
          promo.map((promo) => (
            <Promotion promo={promo} key={promo.id}></Promotion>
          ))}
      </div>
      <hr className="w-full my-5"></hr>
    </Layout>
  );
}
