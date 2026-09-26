import mongoose from 'mongoose';
import dotenv from 'dotenv';

dotenv.config();

async function connectToMongoDB() {
  const dbUrl = process.env.MONGO_URL;

  if (!dbUrl) {
    console.error('MongoDB URL is not configured. Please set the MONGO_URL environment variable.');
    return;
  }

  try {
    await mongoose.connect(dbUrl);
    console.log('Connected to MongoDB');
  } catch (error) {
    console.error('Error connecting to MongoDB:', error);
  }
}

export default connectToMongoDB;


