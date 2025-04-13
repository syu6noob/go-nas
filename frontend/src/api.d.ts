type ApiType = {
  info: string,
  open: string,
  download: string,
}

type StatType = {
  name: string,
  size: number,
  path: string,
  mime: string,
  isDir: boolean,
  link: Link
}

type InfoDirType = {
  isRoot: boolean,
  name: string,
  size: number,
  path: string,
  mime: string,
  isDir: true,
  parent: string,
  parent_link: string,
  link: string,
  api: ApiType,
  children: Stat[]
}

type InfoFileType = {
  isRoot: boolean,
  name: string,
  size: number,
  path: string,
  mime: string,
  isDir: false,
  parent: string,
  parent_link: string,
  link: string,
  api: ApiType,
  children: null
}

type InfoType = InfoDirType | InfoFileType; 

export {
  LinkType,
  StatType,
  InfoType,
  InfoFileType,
  InfoDirType
}