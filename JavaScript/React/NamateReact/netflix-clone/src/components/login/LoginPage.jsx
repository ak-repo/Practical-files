import React, { useRef, useState } from "react";
import { useDispatch } from "react-redux";
import {
  createUserWithEmailAndPassword,
  signInWithEmailAndPassword,
} from "firebase/auth";

import HeaderPage from "../header/HeaderPage";
import { checkValidation } from "../../utils/form/validate";
import { auth } from "../../utils/firebase/firebase";
import { addUser } from "../../utils/redux/Slices/userSlice";
import { useNavigate } from "react-router-dom";
const LoginPage = () => {
  const [isSignInForm, setSignInForm] = useState(true);
  const [error, setError] = useState(null);
  const name = useRef(null);
  const email = useRef(null);
  const password = useRef(null);
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const toggleSignUp = () => {
    setSignInForm(!isSignInForm);
  };

  //click
  const handleButtonClick = (e) => {
    e.preventDefault();
    //form validation
    const message = checkValidation(
      email.current.value,
      password.current.value
    );
    setError(message);
    if (message) return;
    //signIn and signUp

    if (isSignInForm) {
      //signIn
      signInWithEmailAndPassword(
        auth,
        email.current.value,
        password?.current?.value
      )
        .then((userCredential) => {
          const user = userCredential.user;
          navigate("/browse");
        })
        .catch((error) => {
          console.log(error.code);
          console.log(error.message);
        });
    } else {
      //signUp
      createUserWithEmailAndPassword(
        auth,
        email?.current?.value,
        password?.current?.value
      )
        .then((userCredential) => {
          const user = userCredential.user;
          navigate("/browse");
        })
        .catch((error) => {
          console.log(error.code);
          console.log(error.message);
        });
    }
  };
  return (
    <div>
      <HeaderPage />

      {/* login form */}
      <div className="mx-auto w-full md:w-3/6 flex justify-center opacity-90">
        <form
          onClick={(e) => handleButtonClick(e)}
          action=""
          className="flex flex-col items-center w-full max-w-md p-8 md:p-12 bg-black bg-opacity-80 text-white rounded-lg"
        >
          <h1 className="text-center text-3xl md:text-4xl mb-8">
            {isSignInForm ? "Sign In" : "Sign Up"}
          </h1>

          {!isSignInForm && (
            <input
              type="text"
              ref={name}
              placeholder="User Name"
              className="w-full p-3 my-3 bg-gray-700 rounded border border-gray-600 focus:border-red-700 focus:outline-none"
            />
          )}

          <input
            type="text"
            ref={email}
            placeholder="Email Address"
            className="w-full p-3 my-3 bg-gray-700 rounded border border-gray-600 focus:border-red-700 focus:outline-none"
          />
          <input
            type="password"
            ref={password}
            placeholder="Password"
            className="w-full p-3 my-3 bg-gray-700 rounded border border-gray-600 focus:border-red-700 focus:outline-none"
          />
          <button
            type="submit"
            className="w-full p-3 my-6 bg-red-600 hover:bg-red-700 rounded font-bold cursor-pointer"
          >
            {isSignInForm ? "Sign In" : "Sign Up"}
          </button>
          {isSignInForm ? (
            <p className="text-center mt-4">
              <button
                className="cursor-pointer text-gray-400 hover:underline"
                onClick={toggleSignUp}
              >
                Forgot Password?
              </button>
            </p>
          ) : (
            <p className="text-center mt-4 text-gray-400">
              Already registered?{" "}
              <button
                className="cursor-pointer text-white hover:underline"
                onClick={toggleSignUp}
              >
                Sign In Now
              </button>
            </p>
          )}
          {error && <p className="text-red-500 mt-4 text-center">{error}</p>}
        </form>
      </div>
    </div>
  );
};

export default LoginPage;
