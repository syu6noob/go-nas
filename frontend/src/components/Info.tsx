import { useNavigate, useParams } from 'react-router-dom';
import InfoDir from "./info/InfoDir.tsx";
import InfoFile from "./info/InfoFile.tsx";
import { useState, useEffect } from 'react';
import axios from 'axios';

import type { InfoFileType, InfoDirType } from "../api"

export default function Info() {
  const { "*": splat } = useParams();
  const navigate = useNavigate()
  const [data, setData] = useState<{
    isSuccess: boolean,
    data: any
  }>({
    isSuccess: false,
    data: {}
  });

  const fetchTokens = () => {
    const accessToken = window.localStorage.getItem("accessToken");
    const refreshToken = window.localStorage.getItem("refreshToken");
    // console.log('aaa')
    if (accessToken == null || refreshToken == null) {
      // console.log('aaa')
      return {
        accessToken: false,
        refreshToken: false,
      }
    }
    window.localStorage.setItem("login", "true");
    return {
      accessToken: accessToken,
      refreshToken: refreshToken,
    }
  }

  const fetchInfo = (path: string) => {
    (async () => {
      const tokens = fetchTokens();
      if (!tokens.accessToken || !tokens.refreshToken) {
        setData({
          isSuccess: false,
          data: {}
        })
        return;
      }

      try {
        const res = await axios.get(`http://localhost/api/info/${path}`, {
          headers: {
            Authorization: `Bearer ${tokens.accessToken}`,
          }
        });
        setData({
          isSuccess: true,
          data: res.data,
        })
      } catch (e) {
        console.log(e);
        navigate("/");
      }
    })();
  }

  useEffect(() => {
    fetchInfo(splat ?? "");
  }, [splat])

  // console.log(data.isSuccess);
  
  return (
    <div className="flex flex-col grow">
      {
        data.isSuccess
        ? (
          data.data.isDir
            ? <InfoDir data={data.data as InfoDirType} />
            : <InfoFile data={data.data as InfoFileType} />
        )
        : <></>        
      }
    </div>
  )
}