import mongoose from 'mongoose';
import dotenv from 'dotenv';
//  check the db URl is configured or not 

let dbUrl=process.env.MONGO_URL;

// console.log("dbUrl",dbUrl);


async function connectToMongoDB() {
  if (!dbUrl) {
    console.error('MongoDB URL is not configured. Please set the MONGO_URL environment variable.');
    return;
  }
  try {
    const connection = await mongoose.connect(dbUrl, { useNewUrlParser: true, useUnifiedTopology: true });
    console.log('Connected to MongoDB');
  } catch (error) {
    console.error('Error connecting to MongoDB:', error);
  } 
  mongoose.connect(dbUrl);
}

export default connectToMongoDB;


