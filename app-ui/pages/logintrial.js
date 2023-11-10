import axios from "axios";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import Layout from '../components/layout'

axios.defaults.baseURL = "http://localhost:8080";





const Login = ({ user, setUser }) => {

  const [errorResponse, setErrorResponse] = useState(null);
  const [password,setPassword]=useState(null);
  const [username, setUsername]=useState(null);

  
  
  const router = useRouter();
 
 
  const verify= async ()=>{

    try {
      
      const response = await axios.get(`/getUserbyUsername/${username}`);
      console.log(response.data)
      console.log(password)
      
      if(password==response.data.password){
        console.log('here')
        setUser(response.data)
        router.push('/homepage')
      }

    } catch (error) {
      console.log("err", error.response.data.error);
      
     
      
    }

  }

 
  

  return (
    <Layout>
    <div className="flex items-center justify-center max-w-7xl m-auto min-h-screen">
      <div className="flex flex-col items-center justify-center border-2 rounded-xl px-20 py-40 border-black">
        <h1 className="font-bold text-4xl mb-6">Login Page</h1>

       
        <div className="mr-30" >
         <label className="mr-2">Enter userame: </label>
          <input
            type="string"
            className="border-black border-2 rounded-sm mb-5 "
            onChange={(e) => {
              setUsername(e.target.value);
            }}
          />
          </div>

        <div>
         
          <label className="mr-2">Enter password: </label>
          <input
            type="string"
            className="border-black border-2 rounded-sm"
            onChange={(e) => {
              setPassword(e.target.value);
            }}
          />
        
        </div>
       
        <button
          className="font-bold w-auto h-auto bg-sky-300 p-2 mt-4"
          onClick={verify}
        >
          Submit Request
        </button>
       
        <p>{errorResponse}</p>
      </div>
    </div>
    </Layout>
  );
};
export default Login;