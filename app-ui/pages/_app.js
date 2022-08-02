import { useState } from 'react'
import '../styles/globals.css'
import { SessionProvider } from "next-auth/react"


function MyApp({ Component, pageProps:{session, ...pageProps} }) {

  const [user, setUser]=useState(null)
  
  return <SessionProvider session={session}><Component {...pageProps} user={user} setUser={setUser}/></SessionProvider> 
}

export default MyApp
