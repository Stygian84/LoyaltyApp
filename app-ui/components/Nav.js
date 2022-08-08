import Link from "next/link";
import axios from "axios";
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
export default function Nav() {
  const [user, setUser] = useState({});

  const { data: session } = useSession();
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
  }, [session]);
  return (
    <nav className="flex h-20 items-centre px-3 justify-around shadow-md bg-white">
      <div className=" text-lg font-bold text-black flex justify-center items-center">
        <p id="balance" className="mr-4">
          {user?.credit_balance}
        </p>
        available points
      </div>

      <div className=" w-1 h-20 bg-gray-600"> </div>
      <div className="flex justify-center items-center">
        <Link href="/status" className="  px-4  h-20">
          <p id ="transactionPageButton"className="text-lg font-bold cursor-pointer">
            Check Transaction status
          </p>
        </Link>
      </div>
    </nav>
  );
}
