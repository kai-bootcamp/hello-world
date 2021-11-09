import './App.css';
import {
  BrowserRouter,
  Routes,
  Route
} from "react-router-dom";

import HelloWorld from './pages/HelloWorld/index'


function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HelloWorld />}>
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App;
