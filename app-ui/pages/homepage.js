import Program from "../components/program";
import Layout from "../components/layout";
import Nav from "../components/Nav";
const axios = require("axios").default;
import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import Link from "next/link";

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
      <Link href={`/promos`}>
          <button className=" h-12  font-bold mt-10  text-lg cursor-pointer bg-blue-300 rounded drop-shadow-sm px-4 ">
              <p>Offers</p>
          </button>
      </Link>
      
     
      <div className="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3 px-4">
        {prog.map((program) => (
          <Program program={program} key={program.id}></Program>
        ))}
      </div>
      <hr className="w-full my-5"></hr>
    </Layout>
  );
}
