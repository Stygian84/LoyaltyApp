import Program from '../../components/program'
import Layout from '../../components/layout'

import Link from 'next/dist/client/link'
const axios = require('axios').default;
import { useEffect,useState } from "react";
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/router';

export default function Home() {
  
   const { data: session } = useSession();
   
   const [prog,setProg] = useState([])
   const [user,setUser]=useState({})
   const router=useRouter();
   const {query}=useRouter();
   const {currUser}=query;
   
   axios.defaults.baseURL="http://localhost:8080"
   
    //runs the function once when the page is loaded
    const getUser=async ()=> {
        
        
        try{
            const response = await axios.get(`/getUserbyUsername/${currUser}`);
        setUser(response.data);
        } catch (error) {
            console.log("err", error.response.data.error);
          
          }
        
        

    }
       
    
    const displayProgram=async()=>{
        //router.push('/validate')
        // const response = await axios.post('/loyalty',data={
        //     "programId":6,
        //     "stringToTest":"1005610"
        // })
       
        
        const response1 = await axios.get('/loyalty')
        setProg(response1.data); 
        console.log(response1.data);
       
    }

    useEffect(()=>{
        if(router.isReady){
        getUser()
        
        displayProgram()
        }
    },[router.isReady]) 


  return (
    
    <Layout>
      <header>
      <nav className="flex h-20 items-centre px-3 justify-between shadow-md bg-white">
                    <Link href="/"> 
                    
                    <a className=" text-lg font-bold text-black ml-10 mt-5">{user.credit_balance} available points</a>
                    </Link>
                    
                    <div className='absolute ml-300   w-1 h-20 bg-gray-600'></div>
                    <button className='absolute ml-300  px-4  h-20'>Check Transaction status </button>
                    
                    
                    
                    
                </nav>
      </header>
    <h1 className='px-4 h-12 text-lg font-bold'>Offers</h1>
    <div className='grid grid-cols-1 gap-3 md:grid-cols-4 lg:grid-cols px-4'>
      {prog.map((program)=>(
        <Program program={program} key={program.PartnerCode}></Program>
      ))}</div>
      <hr className='w-full my-5'></hr>
  </Layout>
  )
}
