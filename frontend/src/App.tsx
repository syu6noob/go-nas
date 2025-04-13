import { BrowserRouter, Route, Routes } from "react-router-dom";
import Info from "./components/Info"
import Layout from "./components/Layout";
import Login from "./components/Login";
import Logout from "./components/Logout";
import Top from "./components/Top";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route path="" element={<Top />} index />
          <Route path="login" element={<Login />} />
          <Route path="logout" element={<Logout />} />
          <Route path="info/*" element={<Info />}/>
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App
