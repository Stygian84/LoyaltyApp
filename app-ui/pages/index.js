import Program from '../components/program'
import Layout from '../components/layout'
import Link from 'next/dist/client/link'
const axios = require('axios').default;
import { useEffect,useState } from "react";

export default function Home() {

   const [prog,setProg] = useState([])
    useEffect(()=>{
        displayProgram()
    },[]) //runs the function once when the page is loaded
    axios.defaults.baseURL="http://localhost:8080"
    const displayProgram=async()=>{
        //router.push('/validate')
        // const response = await axios.post('/loyalty',data={
        //     "programId":6,
        //     "stringToTest":"1005610"
        // })

        const response = await axios.get('/loyalty')
        setProg(response.data); 
        console.log(response.data);
       
    }


  return (
    
    <Layout>
      <header>
      <nav className="flex h-20 items-centre px-4 justify-between shadow-md bg-white">
                    <Link href="/"> 
                    
                    <a className=" text-lg font-bold text-black ml-10 mt-5">50,000 available points</a>
                    </Link>
                    
                    <div className='absolute ml-300   w-1 h-20 bg-gray-600'></div>
                    <button className='absolute ml-300  px-4  h-20'>Check Transaction status </button>
                    
                    
                    
                </nav>
      </header>
    <h1 className='px-4 h-12 text-lg font-bold'>Offers</h1>
    <div className='grid grid-cols-1 gap-3 md:grid-cols-4 lg:grid-cols'>
      {prog.map((program)=>(
        <Program program={program} key={program.PartnerCode}></Program>
      ))}</div>
      <hr className='w-full my-5'></hr>
  </Layout>
  )
}



