import Home from "./Pages/Home";
import About from "./Pages/About";
import Usage from "./Pages/Usage";
import Users from "./Pages/Users";
import Details from "./Pages/Details";
import NotFound from "./Pages/NotFound";
import "./style.css";
import { Mark, Sports, Remarks } from "./Pages/Details";

import { BrowserRouter, Routes, Route, NavLink } from "react-router-dom";

export default function App() {
  return (
    <BrowserRouter>
      <MenuBar />

      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/usage" element={<Usage />} />
        <Route path="/users" element={<Users />} />
        <Route path="/details/:userID" element={<Details />}>
          <Route path="marks" element={<Mark />} />
          <Route path="sports" element={<Sports />} />
          <Route path="remarks" element={<Remarks />} />
        </Route>
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  );
}

function MenuBar() {
  return (
    <div className="menu">
      <NavLink to="/">Home</NavLink>
      <NavLink to="about">about</NavLink>
      <NavLink to="usage">Usage</NavLink>
      <NavLink to="Users">Users</NavLink>
    </div>
  );
}
