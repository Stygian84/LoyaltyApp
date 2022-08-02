import Program from "../components/program";
import Layout from "../components/layout";
import Nav from "../components/Nav";
const axios = require("axios").default;
import { useEffect, useState } from "react";
import { useRouter } from "next/router";

export default function Home() {
  const [prog, setProg] = useState([]);
  const router = useRouter();

  //runs the function once when the page is loaded

  const displayProgram = async () => {
    const response1 = await axios.get("/loyalty");
    setProg(response1.data);
  };

  useEffect(() => {
    if (router.isReady) {
      displayProgram();
    }
  }, [router.isReady]);

  return (
    <Layout title={`Home Page `}>
      <Nav />
      <h1 className="px-4 h-12 text-lg font-bold mt-10">Offers</h1>
      <div className="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3 px-4">
        {prog.map((program) => (
          <Program program={program} key={program.id}></Program>
        ))}
      </div>
      <hr className="w-full my-5"></hr>
    </Layout>
  );
}
