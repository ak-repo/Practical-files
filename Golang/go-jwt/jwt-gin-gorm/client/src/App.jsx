import AuthProvider from "./context/AuthProvider";
import AuthPage from "./pages/AuthPage";

import { BrowserRouter, Routes, Route } from "react-router-dom";
import User from "./pages/User";
import Admin from "./pages/Admin";
function App() {
  return (
    <>
      <BrowserRouter>
        <AuthProvider>
          <AppRoutes />
        </AuthProvider>
      </BrowserRouter>
    </>
  );
}

export default App;

const AppRoutes = () => {
  return (
    <>
      <Routes>
        <Route path="/" element={<AuthPage />} />
        <Route path="/user" element={<User />} />
        <Route path="/admin" element={<Admin />} />
      </Routes>
    </>
  );
};
