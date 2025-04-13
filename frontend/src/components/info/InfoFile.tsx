import { FaChevronLeft } from "react-icons/fa6"
import type { InfoFileType } from "../../api"
import { useNavigate } from "react-router-dom"

type FileCategory = 'image' | 'video' | 'audio' | 'text' | 'file';

export default function InfoFile({
  data
}: {
  data: InfoFileType
}) {
  const navigate = useNavigate()

  const judgeType = (mime: string): FileCategory => {
    if (mime.startsWith('image/')) return 'image';
    if (mime.startsWith('video/')) return 'video';
    if (mime.startsWith('audio/')) return 'audio';
    if (mime.startsWith('text/')) return 'text';
    return 'file';
  }

  return (
    <div className="flex flex-col gap-6">
      <div className="flex items-center gap-2">
        <button className="cursor-pointer" onClick={() => {
          if (data.path !== "/") navigate("./../")
        }}>
          <FaChevronLeft className="text-2xl" />
        </button>
        <h2 className="text-3xl font-bold overflow-hidden text-ellipsis text-nowrap">{data.name}</h2>
      </div>
      {
        judgeType(data.mime) === "image"
          ? <img
              className="w-full object-cover rounded-xl"
              src={data.api.open}
              alt={data.name}
            />
          : <></>
      }
      {
        judgeType(data.mime) === "video"
          ? <video className="w-full aspect-video object-cover rounded-xl" controls>
              <source src={data.api.open} />
            </video>
          : <></>
      }
      {
        judgeType(data.mime) === "audio"
          ? <audio className="w-full" controls>
              <source src={data.api.open} />
            </audio>
          : <></>
      }

      <div className="flex flex-col">
        <h3 className="text-2xl font-bold mb-2">Details</h3>
        <div className="flex gap-1">
          <span className="font-bold">Name: </span>
          <span>{data.name}</span>
        </div>
        <div className="flex gap-1">
          <span className="font-bold">Size: </span>
          <span>{data.size}</span>
        </div>
        <div className="flex gap-1">
          <span className="font-bold">Mime: </span>
          <span>{data.mime}</span>
        </div>
        <div className="flex gap-1">
          <span className="font-bold">Open: </span>
          <a className="text-blue-500 underline" href={data.api.open}>Open</a>
        </div>
        <div className="flex gap-1">
          <span className="font-bold">Download: </span>
          <a className="text-blue-500 underline" href={data.api.download}>Download</a>
        </div>
      </div>
    </div>
  )
}