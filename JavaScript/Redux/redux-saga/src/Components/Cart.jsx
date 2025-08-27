import React from "react";
const Card = ({ user }) => {
  return (
    <div className="card">
      <div className="card-body">
        <div className="card-title">{user.name}</div>
        <div className="card-subtitle mb-2 text-muted">{user.company.name}</div>
        <div className="card-text">{user.company.catchPhrase}</div>
      </div>
    </div>
  );
};
export default Card;
