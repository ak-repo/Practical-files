import React from "react";
import Logo from "../../assets/logos/Logo";

const HeaderPage = () => {
  return (
    <div>
      <div className="text-left p-10 m-2 bg-gradient-to-b from-black">
        {" "}
        <Logo showText={true} />
      </div>
    </div>
  );
};

export default HeaderPage;
