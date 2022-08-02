import NextAuth from 'next-auth';
import CredentialsProvider from 'next-auth/providers/credentials';
import axios from 'axios';


export default NextAuth({
    
  session: {
    strategy: 'jwt',
  },
  callbacks: {
    async jwt({ token, user }) {
      if (user?._id) token._id = user._id;
      if (user?.isAdmin) token.isAdmin = user.isAdmin;
      
      return token;
    },
    async session({ session, token }) {
      if (token?._id) session.user._id = token._id;
      if (token?.isAdmin) session.user.isAdmin = token.isAdmin;
      
      return session;
    },
  },
  
  providers: [
    CredentialsProvider({
      async authorize(credentials) {
        try {
            
            axios.defaults.baseURL = "http://localhost:8080";
      
            const response = await axios.get('/getUserbyUsername/'+credentials.username);
           

            console.log(response.data)
            console.log("tell me why")
            
            if(credentials.password==response.data.password){
              console.log('here')
              return response.data
              
            }
      
          } catch (error) {
            console.log("err", error.response.data.error);
            
            
           
            
          }

           throw new Error('invalid')
      },
    }),
  ],
});