import { NavLink } from "react-router-dom";

function Navbar() {
  return (
    <div>
      <NavLink to={"/"}>
        <button style={{ padding: "10px " }}>Products</button>
      </NavLink>
      <NavLink to={"/stopwatch"}>
        <button style={{ padding: "10px" }}> Stopwatch</button>
      </NavLink>
      <NavLink to={"/cart"}>
        <button style={{ padding: "10px" }}>Cart</button>
      </NavLink>
    </div>
  );
}

export default Navbar;
