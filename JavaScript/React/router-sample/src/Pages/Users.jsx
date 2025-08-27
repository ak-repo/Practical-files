import axios from "axios";
import { useEffect, useState } from "react";
import { Link, useSearchParams, useNavigate } from "react-router-dom";

function Users() {
  const [data, setData] = useState([]);
  const [searchPa, setSearchPa] = useSearchParams();
  const navigateFun = useNavigate();
  const cls = searchPa.get("class");
  // console.log(cls);

  useEffect(() => {
    axios("http://localhost:3000/users").then((res) => {
      setData(res.data);
    });
  }, []);
  const handleSelectChange = (event) => {
    setSearchPa({ class: event.target.value });
    console.log(searchPa);
  };
  return (
    <div className="page">
      <div>
        <p>
          Select ClassS{" "}
          <select onChange={handleSelectChange}>
            <option value="1">1</option>
            <option value="5">5</option>
            <option value="6">6</option>
            <option value="7">7</option>
            <option value="8">8</option>
          </select>
        </p>
      </div>
      <ul>
        {data
          .filter((user) => {
            if (!cls) {
              return true;
            }
            const item = cls == user.class;
            console.log(item);
            return item;
          })
          .map((user) => (
            <li
              style={{ backgroundColor: "gray", margin: "30px" }}
              onClick={() => {
                navigateFun(`/details/${user.id}`, {
                  state: user,
                });
              }}
              key={user.id}
            >
              <h3>{user.name}</h3>
              <p>{user.address}</p>
              <p>{user.class}</p>
            </li>
          ))}
      </ul>
    </div>
  );
}

export default Users;
