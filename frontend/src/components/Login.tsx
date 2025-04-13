import axios from "axios";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

type Data = {
  access_token: string,
  refresh_token: string,
}

export default function Login() {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")

  const navigate = useNavigate();

  return (
    <div className="flex flex-col grow justify-center items-center">
      <form
        className="w-full flex flex-col justify-start gap-4"
        onSubmit={(e) => {
          e.preventDefault()
          axios.post('http://localhost/login', {
            username, password
          })
            .then(function (response) {
              const data: Data = response.data;
              if (
                Object.prototype.hasOwnProperty.call(data, "access_token")
                && Object.prototype.hasOwnProperty.call(data, "refresh_token")
              ) {
                window.localStorage.setItem("accessToken", data.access_token)
                window.localStorage.setItem("refreshToken", data.refresh_token)
              }
              navigate("/info/"); 
            })
            .catch(function (error) {
              console.log(error.response.data.error);
            });
      }}>
        <input
          name="username"
          type="text"
          placeholder="Username"
          value={username}
          className="flex px-4 py-3 bg-white border border-slate-400 rounded focus:outline-none"
          onChange={e => setUsername(e.target.value)}
        />
        <input
          name="password"
          type="password"
          placeholder="Password"
          value={password}
          className="flex px-4 py-3 bg-white border border-slate-400 rounded focus:outline-none"
          onChange={e => setPassword(e.target.value)}
        />
        <button className="flex px-4 py-3 ml-auto bg-slate-500 text-white rounded hover:bg-slate-700 duration-200" type="submit">
          Submit
        </button>
      </form>
    </div>
  )
}