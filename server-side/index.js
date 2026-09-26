import connectToMongoDB from './src/dbs/mongodb.dbs.js';


async function startServer() {
  await connectToMongoDB();
  // Start your server logic here (e.g., Express app)
}


startServer().catch((error) => {
  console.error('Error starting the server:', error);
  process.exit(1);
});


