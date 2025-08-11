import React from "react";

import LoginPage from "../login/LoginPage";
import Browse from "../browse/Browse";
import { createBrowserRouter, RouterProvider } from "react-router-dom";

const BodyPage = () => {
  //routes
  const appRouter = createBrowserRouter([
    {
      path: "/",
      element: <LoginPage />,
    },
    {
      path: "/browse",
      element: <Browse />,
    },
  ]);
  return (
    <div>
      <RouterProvider router={appRouter} />
    </div>
  );
};

export default BodyPage;
