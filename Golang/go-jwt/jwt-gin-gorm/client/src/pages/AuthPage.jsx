import axios from "axios";
import React, { useContext, useRef, useState } from "react";
import { ginAPI } from "../api";
import { AuthContext } from "../context/context";
import { useNavigate } from "react-router-dom";

function AuthPage() {
  const [login, setlogin] = useState(true);
  const email = useRef();
  const password = useRef();
  const navi = useNavigate();

  //context
  const { setUser } = useContext(AuthContext);

  const handleAuth = async (e) => {
    e.preventDefault();

    // data
    const input = {
      email: email.current.value,
      password: password.current.value,
    };

    // into gin
    try {
      if (login) {
        const { data } = await axios.post(`${ginAPI}/login`, input);
        setUser({
          accessToken: data.token,
          role: data.role,
        });
        if (data.role == "customer") {
          navi("/user");
        } else {
          navi("/admin");
        }
      } else {
        const { data } = await axios.post(`${ginAPI}/register`, input);
        if (data.id > 0) {
          alert("User registration successful");
          setlogin(true);
        }
      }
    } catch (error) {
      console.log(error);
    }
  };
  return (
    <>
      <div className="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
        <div className="sm:mx-auto sm:w-full sm:max-w-sm">
          <img
            alt="Your Company"
            src="https://tailwindcss.com/plus-assets/img/logos/mark.svg?color=indigo&shade=600"
            className="mx-auto h-10 w-auto"
          />
          {login ? (
            <h2 className="mt-10 text-center text-2xl/9 font-bold tracking-tight text-gray-900">
              Sign in to your account
            </h2>
          ) : (
            <h2 className="mt-10 text-center text-2xl/9 font-bold tracking-tight text-gray-900">
              Register New User{" "}
            </h2>
          )}
        </div>

        <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
          <form
            onSubmit={(e) => handleAuth(e)}
            method="POST"
            className="space-y-6"
          >
            <div>
              <label
                htmlFor="email"
                className="block text-sm/6 font-medium text-gray-900"
              >
                email
              </label>
              <div className="mt-2">
                <input
                  ref={email}
                  id="email"
                  name="email"
                  type="email"
                  required
                  //   autoComplete="email"
                  className="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6"
                />
              </div>
            </div>

            <div>
              <div className="flex items-center justify-between">
                <label
                  htmlFor="password"
                  className="block text-sm/6 font-medium text-gray-900"
                >
                  Password
                </label>
                <div className="text-sm">
                  <a
                    href="#"
                    className="font-semibold text-indigo-600 hover:text-indigo-500"
                  >
                    Forgot password?
                  </a>
                </div>
              </div>
              <div className="mt-2">
                <input
                  ref={password}
                  id="password"
                  name="password"
                  type="password"
                  required
                  autoComplete="current-password"
                  className="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6"
                />
              </div>
            </div>

            <div>
              {login ? (
                <button
                  type="submit"
                  className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                  Sign in
                </button>
              ) : (
                <button
                  type="submit"
                  className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                  Sign up
                </button>
              )}
            </div>
          </form>

          {login ? (
            <p className="mt-10 text-center text-sm/6 text-gray-500">
              Not a member?
              <button onClick={() => setlogin(!login)}>Register</button>
            </p>
          ) : (
            <p className="mt-10 text-center text-sm/6 text-gray-500">
              Already have a account?
              <button onClick={() => setlogin(!login)}>Login</button>
            </p>
          )}
        </div>
      </div>
    </>
  );
}

export default AuthPage;
