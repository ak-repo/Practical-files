import React, { useContext } from "react";
import { AuthContext } from "../context/context";

function Admin() {
  const { user } = useContext(AuthContext);

  return (
    <div>
      Admin page
      <h1>user token: {user.accessToken}</h1>
    </div>
  );
}

export default Admin;
