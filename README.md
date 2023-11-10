# Loyalty App

# Client
This is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app).

## Getting Started

First, run the development server:

```bash
npm run dev
# or
yarn dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `pages/index.js`. The page auto-updates as you edit the file.

[API routes](https://nextjs.org/docs/api-routes/introduction) can be accessed on [http://localhost:3000/api/hello](http://localhost:3000/api/hello). This endpoint can be edited in `pages/api/hello.js`.

The `pages/api` directory is mapped to `/api/*`. Files in this directory are treated as [API routes](https://nextjs.org/docs/api-routes/introduction) instead of React pages.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js/) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/deployment) for more details.

# Server

## Getting Started
Start the web server at [localhost:8080](http://localhost:8080)
```bash
go run main.go
```

## Framework
The server is build using the `golang gin web framework`
- [Documentation](https://pkg.go.dev/github.com/gin-gonic/gin)
- [Github](https://github.com/gin-gonic/gin)

## Database
Database of choice is [postgres](https://www.postgresql.org/). `Connection` and `data manipulataion` is done using [gorm](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL).

The `database design` diagram can be found [here](https://app.diagrams.net/#G1MqCC37AmMqTFOKrpWQ0jXXwEO-iu01I5)

Currently it has the problem to connect to remote postgres servers, but for the postgres server running on `localhost `works just fine.

## Folder Structure

This sections explains what should be contained in each folder

### **pkg**
This folder contain all the essential component of the web server

`config` - server configurations
- [app.go](/server/pkg/config/app.go) containts the connection details to the database server

`controllers` - accepts user requests, interacts with the model and selects the view for response
- [loyaltyProgController.go](/server/pkg/controllers/loyaltyProgController.go) defines all the functions that is required for CRUD(Creat, Read, Update, Delete) of the loyalty programs

`models` - manages the data, logic and rules of the application
- [loyaltyProg.go](/server/pkg/models/loyaltyProg.go) contains the interaction of the server and the database for loyalty program

`routes` - manages the routing of the web server, defines which controller should be used at which api endpoint
- [loyaltyProgRoutes.go](/server/pkg/routes/loyaltyProgRoutes.go) defines the routing rules of the loyalty program endpoints
  
`utils` - general utility functions to help with the web server processing

### test
test cases for the  functions
