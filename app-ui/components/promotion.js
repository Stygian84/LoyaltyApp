import Link from "next/link";
import React, { useEffect, useState } from "react";
import Image from "next/dist/client/image";
import LabelContent from "../components/labelContent";
import axios from "axios";
import { useSession } from "next-auth/react";
import { dateTime } from 'Luxon'

export default function Promotion ({  promo }) {
    
    const [prog, setProg] = useState({});
    
    axios.defaults.baseURL="http://localhost:8080"

  const { data: session } = useSession();
  const getProg = async () => {
    try {
      if (session) {
        
        const response = await axios.get(
          `/loyalty/${promo.program}`
        );
        if(response.data.initial_earn_rate!=null){
            console.log(prog.initial_earn_rate==null)
            setProg(response.data);
        }
      }
    } catch (error) {
      console.log(error);
      // console.log("err", error.response.data.error);
    }
  };

  
  useEffect(() => {
    
    getProg();
  }, [session, prog.initial_earn_rate]);


    switch(promo.earn_rate_type){
        case "add":
            //getProg(promo.program)
            var earnRate=0;
            if(prog.initial_earn_rate!=null){
                

                earnRate = prog.initial_earn_rate;
                var offerInfo="extra";
            }
            break;
        case "mul":
            //getProg(promo.program)
            var earnRate=0;
            if(prog.initial_earn_rate!=null){
                earnRate = prog.initial_earn_rate*promo.constant;
                var offerInfo="X";
            }
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
            

            <div className="flex flex-col items-center justify center p-5">
             <p className=" text-pink font-bold text-lg mb-4">{promo.constant} {offerInfo} points offered!</p>
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

