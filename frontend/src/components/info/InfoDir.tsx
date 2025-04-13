import { useNavigate } from "react-router-dom";
import type { InfoDirType, StatType } from "../../api"
import LinkButton from "./LinkButton"

type FileCategory = 'image' | 'video' | 'audio' | 'text' | 'file';

export default function InfoDir({
  data
}: {
  data: InfoDirType
}) {
  // console.log(`../${data.parent}`);
  const navigate = useNavigate()

  const judgeType = (mime: string): FileCategory => {
    console.log(mime);
    if (mime.startsWith('video/')) return 'video';
    if (mime.startsWith('audio/')) return 'audio';
    if (mime.startsWith('text/')) return 'text';
    return 'file';
  }

  return (
    <div className="flex flex-col">
      <LinkButton
        type={data.isRoot ? "root" : "parent"}
        name={data.isRoot ? "Root" : data.name}
        onClick={() => data.path === "/" ? false : navigate("./../")}
      />
      {
        Object.assign(data.children).map((item: StatType, i: number) => {
          return (
            <LinkButton
              child
              key={i}
              type={item.isDir ? 'dir' : judgeType(item.mime)}
              name={item.name}
              onClick={() => navigate(encodeURI(`${item.name}`))}
            />
          )
        })
      }
    </div>
  )
}