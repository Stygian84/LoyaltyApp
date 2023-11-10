import React, { useContext,useState, useEffect } from 'react'
import Layout from '../../components/layout'
import {useRouter} from 'next/router'
import Link from 'next/link';
import Image from 'next/image'
const axios = require('axios').default;
import { useForm } from 'react-hook-form';
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

export default function Create ()  {
    
    const router=useRouter();
    const {query}=useRouter();
    const {id}=query;
  

     //runs the function once when the page is loaded
    axios.defaults.baseURL="http://localhost:8080"
    const [title, setTitle] = useState('');
    const [title2, setTitle2] = useState('');
    const{handleSubmit}=useForm();
   
    const validateMembership=async()=>{
        
        console.log("id:"+id);;
        console.log("mem_num:"+title2);
        const data={
            "programId":parseInt(id),
            "stringToTest":title2
        }

        const response = await axios.post('/loyalty/validateMembership', data)
        
        console.log(response)
        console.log('valid: '+response.data.valid)
        if(title!=title2 || response.data.valid==false){
            toast('Membership number is incorrect', {
                position: "top-center",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                });
        }

        else if(response.data.valid==true){
            toast('Membership number is correct', {
                position: "top-center",
                autoClose: 5000,
                hideProgressBar: false,
                closeOnClick: true,
                pauseOnHover: true,
                draggable: true,
                progress: undefined,
                });
        }

        
    }

    

    return (
        
     <Layout>
    
    <h2 className='ml-500'><b>Validate membership</b></h2>
    
    <div className="create items-center ml-300">
    
        
        <form className='mx-auto max-w-screen-md ' onSubmit={handleSubmit(validateMembership)}>
            <ToastContainer className='w-25' icon={false} />
            

        <div className='mb-4'>
        <label><b>Membership</b></label>
        <input 
            type="text" 
            required 
            value={title}
            
            onChange={(e) => setTitle(e.target.value)}
            className='w-full' id='text1' autoFocus
        />
        </div>

        <div>
        <p>
        <label><b>Confirm Membership</b></label>
        <input
            type="text"
            required
            value={title2}
            id="text2"
            className='w-full'autoFocus
            onChange={(e) => setTitle2(e.target.value)}
        />
        </p>
        </div>
        
        <p>

        <button className='rounded w-40 bg-blue-600 text-white mt-4'>
            Save Membership
        </button>
        

        </p>

        </form>
    </div>
    </Layout>   
    );
}