import React from "react";
import Logo from "../../assets/logos/Logo";
import { signOut } from "firebase/auth";
import { auth } from "../../utils/firebase/firebase";
import { useNavigate } from "react-router-dom";

const HeaderPage = () => {
  const navigate = useNavigate();

  //signout

  const handleSignOut = () => {
    signOut(auth)
      .then(() => {
        navigate("/");
      })
      .catch((error) => {
        console.log(error.message);
      });
  };
  return (
    <div className="flex justify-between ">
      <div className="text-left p-10 m-2 ">
        {" "}
        <Logo showText={true} />
      </div>
      <div className="m-4 mt-10 ">
        <button
          onClick={handleSignOut}
          className="bg-red-600 text-white rounded p-2  "
        >
          Sign Out
        </button>
      </div>
    </div>
  );
};

export default HeaderPage;
