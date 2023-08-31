interface IObjectKeys {
  [key: string]: string | boolean | ILanguageItems[] | null
}

interface IClient extends IObjectKeys{
  firstName: string
  lastName: string
  aboutMe: string
  regDate: string
  avatar: string
  isTeacher: boolean | null
  interfaceLang: string
  languageItems: ILanguageItems[]
}

interface ILanguageItems {
  lang: string
  lvl: string
}

interface IGetDataClient {
  dataClient: IClient
  responseStatus: IStatusResponse
}