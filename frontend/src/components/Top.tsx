import { useState, useEffect } from "react";
import { Link } from "react-router-dom";

export default function Top() {
  const [accessToken, setAccessToken] = useState<string | null>();
  
  useEffect(() => {
    const accessTokenTemp = window.localStorage.getItem("accessToken")
    setAccessToken(accessTokenTemp);
  }, [])

  return (
    <div className="flex flex-col grow justify-center items-center gap-6">
      <h1 className="text-3xl font-bold">GO-NAS</h1>
      {
        accessToken
        ? (
          <Link to="/info" className="px-4 py-2 text-xl border border-slate-400 rounded-md hover:bg-slate-200 duration-150">
            View files
          </Link>
        )
        : (
          <Link to="/login" className="px-4 py-2 text-xl border border-slate-400 rounded-md hover:bg-slate-200 duration-150">
            login
          </Link>
        )
      }
    </div>
  )
}