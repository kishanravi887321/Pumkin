import connectToMongoDB from './src/dbs/mongodb.dbs.js';
import  app from './src/app.js';
import dotenv from 'dotenv';
dotenv.config();

const PORT=process.env.PORT || 3000;

console.log("PORT",PORT);


async function startServer() {
  await connectToMongoDB();

    app.listen(PORT, () => {
      console.log(`Server is running on port ${PORT}`);
    });
  // Start your server logic here (e.g., Express app)
}


startServer().catch((error) => {
  console.error('Error starting the server:', error);
  process.exit(1);
});


