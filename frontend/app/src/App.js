import React, { useState, useEffect } from "react";
import "./App.css";
import API from "./API";

const App = () => {
  const [text, setText] = useState("");
  const [input, setInput] = useState(0);

  useEffect(() => {
    setText("This text will be change when you click on the button");
  }, []);

  const randomString = async () => {
    const response = await API.getRandomString();
    setText(response);
  };

  const randomCharacter = async () => {
    let response;
    let random = Math.random() * 10;
    if (input === 0) {
      response = await API.getRandomCharacters(random);
    } else {
      response = await API.getRandomCharacters(input);
    }
    setText(response);
  };

  function onInputChange(event) {
    setInput(event.target.value);
  }

  return (
    <div className="app">
      <div className="app__header">
        <p>Hello World</p>
        <div className="app__container">
          <p>{text}</p>
        </div>
        <div className="button__container">
          <div className="button__container--left">
            <button onClick={() => randomString()}>Random String</button>
          </div>
          <div className="button__container--right">
            <button onClick={() => randomCharacter()}>Random Characters</button>
            <input
              value={input}
              type="number"
              onChange={(e) => onInputChange(e)}
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default App;
