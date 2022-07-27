import Link from "next/link";
import React from "react";
import Image from "next/dist/client/image";
import { useRouter } from "next/router";
const axios = require('axios').default;
import { useEffect,useState } from "react";

    // const [valid,setValid] = useState(false)
    // useEffect(()=>{
    //     validateMembership()
    // },[])

    // const validateMembership=()=>{
    //     router.push('/validate')

    // }


export default function Program({program}){

    return(

        <div className="card rounded shadow flex flex-col items-center justify center p-5 ">
            
            {/* <Image 
            src= {program.image}
            alt={program.description}
            className="rounded shadow items-center "
            width="100%"
            height="100%"
            >
            </Image> */}
                
            <div className="flex flex-col items-center justify center p-5">
                
                
                <h2 className="text-lg">{program.name}</h2>
                

                
                <p className="mb-2">{program.initial_earn_rate} points</p>
                <p className="mb-2">{program.processing_time}</p>
                <Link href={`/${program.id}`}>
                <a>
                <h2 className="text-lg bg-blue-300 rounded drop-shadow-sm px-4 ">Transfer credits</h2>
                </a>
                </Link>
               
                
            </div>
        </div>
    )
}
