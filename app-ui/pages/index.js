import React, { useEffect, useState } from 'react';
import { signIn, useSession, signOut } from 'next-auth/react';
import { useForm } from 'react-hook-form';
import Layout from '../components/layout';

import { toast } from 'react-toastify';
import { useRouter } from 'next/router';

export default function LoginScreen() {

  const { data: session } = useSession();

  const router = useRouter();
  const { redirect } = router.query;
  const [username, setUsername]=useState();
  const [password, setPassword]=useState();

 const {
    handleSubmit,
  
    } = useForm();

  useEffect(() => {
    // if (session?.user) {
    //   router.push(redirect || '/');
    //   console.log(session)
    // }
    console.log(session)
    
  }, [router, session, redirect]);
  
    const submitHandler = async () => {
      try {
        console.log(username)
        const result = await signIn('credentials', {
          redirect: false,
          username,
          password,
        });

        console.log('result: ' + result.error)

        if(result==null){
          toast.error("result.error");
        }
        else if (result.error) {
          toast.error(result.error);
        }

        else{
          router.push('/homepage/'+username);
        }
      } catch (err) {
        toast.error(err);
      }
    };

  if(session){
   

      return (
        <>
       Signed in as {session.user.name} <br />
        <button onClick={() => signOut()}>Sign out</button>
       </>
      )
      
    

  }

  else{
    return (
      <Layout title='login'>
          <form className='mx-auto max-w-screen-md' onSubmit={handleSubmit(submitHandler)}>
              <h1 className='mb-4 text-xl'>Login</h1>
              <div className='mb-4'>
                  <label htmlFor='username'>Username</label>
                  <input 
                   className='w-full' id='username' autoFocus onChange={(e) => {
                    setUsername(e.target.value);
                  }}>
  
                  </input>
                 
  
              </div>
              <div className='mb-4'>
                  <label htmlFor='password'>Password</label>
                  <input type='password' 
                 
                   className='w-full' id='password' autoFocus onChange={(e) => {
                    setPassword(e.target.value);
                  }}></input>
                   
  
              </div>
              <div className='mb-4'>
                  <button className='primary-button'>Login</button>
              </div>
              
          </form>
      </Layout>
    )
  }
  
}
