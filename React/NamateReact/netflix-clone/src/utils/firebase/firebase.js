// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
import { getAuth } from "firebase/auth";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyCfLnRqKLHIqz-HImYngs6L2Uk2yUCqxOI",
  authDomain: "netflixclone-bd0b7.firebaseapp.com",
  projectId: "netflixclone-bd0b7",
  storageBucket: "netflixclone-bd0b7.firebasestorage.app",
  messagingSenderId: "896933768093",
  appId: "1:896933768093:web:8ebabf033618cf9ed05d74",
  measurementId: "G-1KY2QCS403",
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);

//auth
export const auth = getAuth();
