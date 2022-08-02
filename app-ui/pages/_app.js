import { useState } from "react";
import "../styles/globals.css";
import { SessionProvider } from "next-auth/react";
import axios from "axios";

function MyApp({ Component, pageProps: { session, ...pageProps } }) {
  axios.defaults.baseURL = "http://localhost:8080";

  const [user, setUser] = useState(null);

  return (
    <SessionProvider session={session}>
      <Component {...pageProps} user={user} setUser={setUser} />
    </SessionProvider>
  );
}

export default MyApp;
