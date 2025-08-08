import axios from "axios";
import { useEffect, useState } from "react";
import {
  useParams,
  useNavigate,
  useLocation,
  NavLink,
  Outlet,
} from "react-router-dom";
import { useOutletContext } from "react-router-dom";

function Details() {
  const navigateFun = useNavigate();
  const { state } = useLocation();

  const { userID } = useParams();
  const [user, setUser] = useState(state);

  useEffect(() => {
    if (!user) {
      axios("http://localhost:3000/users")
        .then((res) => {
          return res.data;
          // fallback if not found
        })
        .then((data) => {
          const items = data.find((item) => item.id === userID);
          setUser(items || {});
        });
    }
  }, [userID, user]);
  if (!user) {
    return null;
  }

  return (
    <div className="page">
      <p>Name: {user.name}</p>
      <p>Address: {user.address}</p>
      <p>Phone: {user.phone}</p>

      <div>
        <NavLink className="btn" to="marks">
          mark
        </NavLink>
        <NavLink className="btn" to="sports">
          sports
        </NavLink>
        <NavLink className="btn" to="remarks">
          remarks
        </NavLink>
      </div>
      <div>
        {" "}
        <Outlet context={user} />
      </div>

      <button
        onClick={() => {
          navigateFun(-1);
        }}
      >
        Back
      </button>
    </div>
  );
}

export default Details;

export function Mark() {
  const { marks } = useOutletContext();

  return (
    <div>
      {Object.keys(marks).map((mark) => (
        <li key={mark}>
          {mark}={marks[mark]}{" "}
        </li>
      ))}
    </div>
  );
}
export function Sports() {
  return <div>Sports</div>;
}
export function Remarks() {
  return <div>Remarks</div>;
}
