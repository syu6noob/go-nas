import { useEffect, useState } from "react"
import { Link, Outlet, useLocation } from "react-router-dom"

export default function Layout() {
  const location = useLocation();
  const [accessToken, setAccessToken] = useState<string | null>();

  useEffect(() => {
    const accessTokenTemp = window.localStorage.getItem("accessToken")
    setAccessToken(accessTokenTemp);
  }, [location])

  return (
    <div className="flex w-screen min-h-screen flex-col bg-slate-100 text-slate-900">
      <header className="flex items-center p-6 border-b border-slate-400">
        <h1 className="flex grow text-2xl font-bold">GO-NAS</h1>
        {
          accessToken !== null
          ? <Link to="/logout" className="px-3 py-1.5 border border-slate-400 rounded-md hover:bg-slate-200 duration-150">
              Logout
            </Link>
          : <></>
        }
      </header>
      <div className="flex flex-col grow p-6">
        <Outlet />
      </div>
    </div>
  )
}