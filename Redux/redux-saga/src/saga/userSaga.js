import { call, put, takeEvery } from "redux-saga/effects";

const apiUrl = "https://jsonplaceholder.typicode.com/users";

function getApiData() {
  return fetch(apiUrl)
    .then((response) => response.json())
    .catch((error) => error.message);
}

// saga

function* fetchUsers(action) {
  try {
    const users = yield call(getApiData);
    yield put({ type: "GET_USERS_SUCCESS", usersList: users });
  } catch (error) {
    yield put({ type: "GET_USERS_FAILED", message: error.message });
  }
}

function* userSaga() {
  yield takeEvery("GET_USERS_REQUESTED", fetchUsers);
}
export default userSaga;
