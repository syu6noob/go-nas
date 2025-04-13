import { clsx } from "clsx";
import { MouseEventHandler } from "react";
import {
  FaFolder,
  FaFolderOpen,
  FaHardDrive,
  FaImage,
  FaHeadphones,
  FaVideo,
  FaFileLines,
  FaFile
} from "react-icons/fa6"; 

export default function FileButton({
  child = false,
  name,
  type,
  onClick
}: {
  child?: boolean,
  name: string,
  type: 'root' | 'parent' | 'dir' | 'file' | 'image' | 'audio' | 'video' | 'text'
  onClick?: MouseEventHandler,
}) {
  let icon;
  switch (type) {
    case 'root':
      icon = <FaHardDrive className="w-6 text-xl fill-slate-500" />
      break;
    case 'parent':
      icon = <FaFolderOpen className="w-6 text-xl fill-amber-500" />
      break;
    case 'dir':
      icon = <FaFolder className="w-6 text-xl fill-amber-500" />
      break;
    case 'image':
      icon = <FaImage className="w-6 text-xl fill-red-400" />
      break;
    case 'audio':
      icon = <FaHeadphones className="w-6 text-xl fill-lime-500" />
      break;
    case 'video':
      icon = <FaVideo className="w-6 text-xl fill-purple-400" />
      break;
    case 'text':
      icon = <FaFileLines className="w-6 text-xl fill-blue-400" />
      break;
    default:
      icon = <FaFile className="w-6 text-xl fill-stone-400" />
      break;
  }
  return (
    <div className={clsx(
      "flex flex-col justify-start",
      child ? "ml-4" : "ml-0"
    )}>
      <button
        onClick={onClick}
        className={clsx(
          "flex px-3 py-2 rounded-md items-center gap-2 bg-transparent hover:bg-slate-200 duration-100",
        )}
      >
        {icon}
        <span className="text-left">{name}</span>
      </button>
    </div>
  )
} 