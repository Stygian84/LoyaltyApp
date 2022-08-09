import Link from "next/link";
import React, { useEffect, useState } from "react";
import Image from "next/dist/client/image";
import LabelContent from "../components/labelContent";
import axios from "axios";
import { useSession } from "next-auth/react";

export default function Promotion ({  promo }) {
    
    const [prog, setProg] = useState({});
    const [user,setUser]=useState({});
    const [validPromo, setValidPromo]=useState([]);

  const { data: session } = useSession();
  const getProg = async () => {
    try {
      if (session) {
        const response1 = await axios.get(
            `/getUserbyEmail/${session.user.email}`
          );
        setUser(response1.data[0]);
        const response = await axios.get(
          `/loyalty/${promo.program}`
        );
        setProg(response.data);
      }
    } catch (error) {
      console.log(error);
      // console.log("err", error.response.data.error);
    }
  };
  useEffect(() => {
    getProg();
  }, [session]);


    switch(promo.earn_rate_type){
        case "add":
            //getProg(promo.program)
            var earnRate = prog.initial_earn_rate+promo.constant;
            break;
        case "mul":
            //getProg(promo.program)
            var earnRate = prog.initial_earn_rate*promo.constant;
            break;
    }

    // const getProg=async ({progID})=>{
    //     const response1 = await axios.get(`/GetLoyaltyId/${progID}`);
    //     setProg(response1.data);
    // }
    // useEffect(() => {
    //     if (router.isReady) {
    //       getProg(program.id)
    //     }
    //   }, [router.isReady, session]);
    
   
    
     
    return (
        <div className="card rounded shadow flex flex-col items-center justify center p-5 border-4 border-yellow-500">
            {/* <Image 
            src= {program.image}
            alt={program.description}
            className="rounded shadow items-center "
            width="100%"
            height="100%"
            >
            </Image> */}
            <p className="text-white font-thin">{earnRate}</p>

            <div className="flex flex-col items-center justify center p-5">
                
                <LabelContent title="Program Name">
                    <h2 className="text-lg">{prog.name}</h2>
                </LabelContent>
                <LabelContent title="Promotion Type">
                    
                    <p className="">{promo.promo_type} </p>
                </LabelContent>
                <LabelContent title="Point to Rewards Ratio">
                    
                    <p className="">
                        1 point : {earnRate.toFixed(2)}{" "}
                        {prog.currency_name}
                    </p>
                </LabelContent>
                <LabelContent title="Partner Code">
                    <p className="">{prog.partner_code}</p>
                </LabelContent>
                <LabelContent title="Estimated Transfer Time">
                    <p className="">Up to {prog.processing_time}</p>
                </LabelContent>

                <LabelContent title="End Date">
                    <p className="">{promo.end_date}</p>
                </LabelContent>
                <Link href={`/transaction/${prog.id}`}>
                    <button className="mt-4 text-lg cursor-pointer bg-blue-300 rounded drop-shadow-sm px-4 ">
                        <p>Transfer credits</p>
                    </button>
                </Link>
            </div>
        </div>
    );
};

