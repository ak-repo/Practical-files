import { Provider } from "react-redux";

import store from "./redux/store";
import BlogCreation from "./componets/BlogCreation";

function App() {
  return (
    <Provider store={store}>
      <BlogCreation />
    </Provider>
  );
}

export default App;
