import React, { useContext } from "react";
import { AuthContext } from "../context/context";

function User() {
  const { user } = useContext(AuthContext);

  return (
    <div>
      User Page
      <h1>user toke: {user.accessToken}</h1>
    </div>
  );
}

export default User;
