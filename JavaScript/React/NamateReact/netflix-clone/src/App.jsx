import { Provider } from "react-redux";
import BodyPage from "./components/body/BodyPage";
import store from "./utils/redux/store";
function App() {
  return (
    <Provider store={store}>
      <BodyPage />
    </Provider>
  );
}

export default App;
