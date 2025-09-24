import { useState } from "react";
import { AuthContext } from "./context";

function AuthProvider({ children }) {
  const [user, setUser] = useState({});

  return (
    <AuthContext.Provider value={{ user, setUser }}>
      {children}
    </AuthContext.Provider>
  );
}

export default AuthProvider;
