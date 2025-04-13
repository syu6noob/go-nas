import { useEffect } from "react"
import { Link } from "react-router-dom"

export default function Logout() {
  useEffect(() => {
    window.localStorage.removeItem("accessToken")
    window.localStorage.removeItem("refreshToken")
  }, [])
  return (
    <div className="flex flex-col grow justify-center items-center gap-6">
      <h2 className="text-3xl font-bold">Logout Succeed</h2>
      <Link to="/login" className="px-3 py-1.5 border border-slate-400 rounded-md hover:bg-slate-200 duration-150">
        Login
      </Link>
    </div>
  )
}