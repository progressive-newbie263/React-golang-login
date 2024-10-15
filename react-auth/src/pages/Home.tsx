import React, { useEffect, useState } from "react";

const Home = (props: {name: string}) => {
  return (
    <div className="text-center">
      <h1><span className="text-slate-500">Welcome to the Home Page</span></h1>
      <h3>{props.name ? 'Hi, ' + props.name : 'Please log in for more info!'}.</h3>
    </div>
  );
}

export default Home;