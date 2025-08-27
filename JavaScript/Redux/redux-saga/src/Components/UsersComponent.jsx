import { useDispatch, useSelector } from "react-redux";
import { getUsers } from "../redux/Reducers/rootReducer";
import Card from "./Cart";
import { useEffect } from "react";

function UsersComponent() {
  const dispatch = useDispatch();
  const { usersList, loading, error } = useSelector((state) => state.users);

  useEffect(() => {
    dispatch(getUsers());
  }, [dispatch]);

  return (
    <>
      {usersList.length > 0 &&
        usersList.map((user) => <Card user={user} key={user.id} />)}
      {usersList.length === 0 ? <p>No usersList</p> : null}
      {usersList.length === 0 && loading === true ? <p>Loading...</p> : null}
      {error === 0 && !loading === true ? <p>{error.message}</p> : null}
    </>
  );
}

export default UsersComponent;
