import logo from './logo.svg';
import './App.css';
import { Routes, Route, Link } from "react-router-dom";


function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <Routes>
        <Route path="/" element={<HelloWorld />} />
        <Route path="homework2" element={<Homework2 />} />
      </Routes>
      </header>
    </div>
  );
}

function HelloWorld() {
  return (
    <>
      <main>
        <h2>Helluuuu</h2>
        <p>Welcome to the homepage!</p>
      </main>
      <nav>
        <Link to="/homework2">Homework2</Link>
      </nav>
    </>
  );
}

function Homework2() {
  return (
    <>
      <main>
        <h2>Test</h2>
        <p>
          API here
        </p>
      </main>
      <nav>
        <Link to="/">Home</Link>
      </nav>
    </>
  );
}

export default App;
