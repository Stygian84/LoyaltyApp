import Link from "next/link";
import React from "react";
import Image from "next/dist/client/image";
import LabelContent from "../components/labelContent";

import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
const axios = require("axios").default;

const Program = ({ program }) => {

  const [offer, setOffer] = useState(false);
  const [promo, setPromo] = useState({});
  const [user,setUser]=useState({});

  const { data: session } = useSession();
  axios.defaults.baseURL="http://localhost:8080"
  const getPromo = async () => {
    try {
      
      if (session) {  
        getUser();       
        const response = await axios.get(
            `/getPromoByDate/${program.id}`
            );
            //console.log(response.data)
        
       

        if(response.data!=null && user.card_tier!=null){
            for(var i=0;i<response.data.length;i++){
                
                console.log(user.card_tier==null)
                if( response.data[i].card_tier==user.card_tier){
                    
                    setPromo(response.data[i]);
                    setOffer(true)
                    
                    break;

                }
            }
        }
        
      }
    } catch (error) {
      console.log(error);
      // console.log("err", error.response.data.error);
    }
  };
  const getUser = async () => {
    try {
      if (session) {
        const response = await axios.get(
          `/getUserbyEmail/${session.user.email}`
        );
        setUser(response.data[0]);
      }
    } catch (error) {
      console.log(error);
      // console.log("err", error.response.data.error);
    }

  };

  
  
  useEffect(() => {
    
        getUser();
        getPromo();
    
    
    
  }, [user.card_tier, session]);

  switch(promo.earn_rate_type){
    case "add":
        //getProg(promo.program)
        var earnRate = "additional"
        break;
    case "mul":
        //getProg(promo.program)
        var earnRate = "X"
        break;
}

  if(offer){
    return (

        <div className="card rounded shadow flex flex-col items-center justify center p-5 border-4  border-yellow-500">
            {/* <Image 
            src= {program.image}
            alt={program.description}
            className="rounded shadow items-center "
            width="100%"
            height="100%"
            >
            </Image> */}
            
            <div className="flex flex-col items-center justify center p-5">  
            
            <div className=" flex flex-col items-center justify center mb-4">
                <p className="font-bold center text-lg ">OFFER! </p>
                <p className="text-lg text-center"><span className="text-xl">{promo.constant.toFixed(2)}</span> <span className='font-bold text-red-500'>{earnRate}</span> points offered</p>
            </div>

                <LabelContent title="Program Name">
                    <h2 className="text-lg">{program.name}</h2>
                </LabelContent>
                <LabelContent title="Point to Rewards Ratio">
                    <p className="">
                        1 point : {program.initial_earn_rate.toFixed(2)}{" "}
                        {program.currency_name}
                    </p>
                </LabelContent>
                <LabelContent title="Partner Code">
                    <p className="">{program.partner_code}</p>
                </LabelContent>
                <LabelContent title="Estimated Transfer Time">
                    <p className="">Up to {program.processing_time}</p>
                </LabelContent>
                <Link href={`/transaction/${program.id}`}>
                    <button className="mt-4 text-lg cursor-pointer bg-blue-300 rounded drop-shadow-sm px-4 ">
                        <p>Transfer credits</p>
                    </button>
                </Link>
            </div>
        </div>
    );
} else{
    return(
        <div className="card rounded shadow flex flex-col items-center justify center p-5  border">
            {/* <Image 
            src= {program.image}
            alt={program.description}
            className="rounded shadow items-center "
            width="100%"
            height="100%"
            >
            </Image> */}

            <div className="flex flex-col items-center justify center p-5">  

                <LabelContent title="Program Name">
                    <h2 className="text-lg">{program.name}</h2>
                </LabelContent>
                <LabelContent title="Point to Rewards Ratio">
                    <p className="">
                        1 point : {program.initial_earn_rate.toFixed(2)}{" "}
                        {program.currency_name}
                    </p>
                </LabelContent>
                <LabelContent title="Partner Code">
                    <p className="">{program.partner_code}</p>
                </LabelContent>
                <LabelContent title="Estimated Transfer Time">
                    <p className="">Up to {program.processing_time}</p>
                </LabelContent>
                <Link href={`/transaction/${program.id}`}>
                    <button className="mt-4 text-lg cursor-pointer bg-blue-300 rounded drop-shadow-sm px-4 ">
                        <p>Transfer credits</p>
                    </button>
                </Link>
            </div>
        </div>
    );

    
}
}


export default Program;
